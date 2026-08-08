---
name: bug:expose
description: >
  Use this skill to expose a bug by writing a failing Go test.
  Triggers include: "expose a bug", "write a failing test", "reproduce a bug",
  "prove this is broken", "add a regression test for a bug", or any request to
  document and pin down incorrect behaviour before fixing it.
  Output is always a failing Go test function in the relevant _test.go file.
---

# Bug-Expose Skill

## Overview

Capture a bug as a **red test**: a Go test that fails because the system
behaves incorrectly. The test asserts the _correct_ expected behaviour so it
flips from red to green when the fix lands. Do **not** fix the bug here —
only expose it.

---

## Workflow

```
1. Understand the bug
   └─ What is the current (wrong) output?
   └─ What is the correct expected output?
   └─ What is the minimal topology / input that triggers it?

2. Write the failing Go test
   └─ Descriptive name: Test<Subject><Condition><ExpectedOutcome>
   └─ Minimal Given chain — strip everything not needed to reproduce
   └─ Then step asserts the CORRECT output → RED while bug exists

3. Run the test to confirm it fails
   └─ go test -run Test<Name> ./path/to/package/...
   └─ Record the actual failure message

4. Stop — leave the test red. Do not fix the bug here.
```

---

## Test Skeleton

```go
// Test<Subject><Condition><ExpectedOutcome> reproduces <one-line bug description>.
// RED until the fix is applied.
func Test<Subject><Condition><ExpectedOutcome>(t *testing.T) {
    b := c2build.NewTopologyBuilder()

    tst.NewScenario(t).
        Given(b.Topology()).
        Given(b.AdminZone("<zone>")).
        Given(b.Device("<device>", "<schema>", "<protocol>")).
        And( /* minimal additional setup */ ).
        When(ConfigBuiltFor("<device>")).
        Then(<AssertCorrectBehaviour>()).
        Execute()
}

func <AssertCorrectBehaviour>() *tst.StepResult {
    sr := tst.NewStepResult("<plain English description of correct outcome>")
    got := /* read from ctx.config */
    if got != <expected> {
        return sr.WithErrorf("want %q, got %q", <expected>, got)
    }
    return sr.Success()
}
```

### Key rules

- **One scenario per bug.** Exactly one `When` and one primary failure mode.
  Add `And` assertions only if directly related.
- **Minimal topology.** Strip the `Given` chain to the smallest setup that
  still reproduces the bug. Extra devices or port pairs obscure the root cause.
- **Assert the correct value.** The `Then` step compares against what the
  system _should_ produce, not what it currently produces.
- **Descriptive name.** `Test<Subject><Condition><ExpectedOutcome>` — name
  what is tested and what the correct result is, e.g.
  `TestDeviceTwoDevicesSameMgmtAddressReportsError`,
  `TestNetworkNoCIDRBuildsWithoutPanic`.
- **One-line comment.** Minimal text about what the test proves.

---

## Step Function Patterns

Reuse existing step functions from `builder_test.go` where possible. Write a
new step only when no existing one matches the assertion.

### Wrong value

```go
func <FieldIs>(expected string) *tst.StepResult {
    sr := tst.NewStepResult("<field> is %q", expected)
    got := string(ctx.config.<Section>.<Field>)
    if got != expected {
        return sr.WithErrorf("want %q, got %q", expected, got)
    }
    return sr.Success()
}
```

### Absent section (nil guard)

```go
func <SectionIsAbsent>() *tst.StepResult {
    sr := tst.NewStepResult("<section> is absent")
    if ctx.config.<Section> != nil {
        return sr.WithErrorf("expected nil, got %+v", ctx.config.<Section>)
    }
    return sr.Success()
}
```

### Wrong count

```go
func <SectionCountIs>(n int) *tst.StepResult {
    sr := tst.NewStepResult("<section> count is %d", n)
    got := len(ctx.config.<Section>.<Map>)
    if got != n {
        return sr.WithErrorf("want %d, got %d", n, got)
    }
    return sr.Success()
}
```

---

## Full Example

```go
// TestNtpPoolServerAddressIsPresent reproduces a bug where enabling the pool flag
// causes the NTP server address to be cleared in the generated config.
// RED until the fix is applied.
func TestNtpPoolServerAddressIsPresent(t *testing.T) {
    b := c2build.NewTopologyBuilder()

    tst.NewScenario(t).
        Given(b.Topology()).
        Given(b.AdminZone("Global")).
        Given(b.Device("App001", "6.3", "IPsec")).
        And(b.WithNtpPool("pool.ntp.org")).
        When(ConfigBuiltFor("App001")).
        Then(NtpServerAddressIsNotEmpty("srv-1")).
        Execute()
}

func NtpServerAddressIsNotEmpty(key string) *tst.StepResult {
    sr := tst.NewStepResult("NTP server %s address is not empty", key)
    srv, ok := ctx.config.System.Ntp.Server[key]
    if !ok {
        return sr.WithErrorf("NTP server %s not found", key)
    }
    if srv.Server == "" {
        return sr.WithErrorMessage("NTP server address is empty — bug is present")
    }
    return sr.Success()
}
```

Confirm it is red:

```
go test -run TestNtpPoolServerAddressIsPresent ./internal/device/config/generic/...
--- FAIL: TestNtpPoolServerAddressIsPresent (0.00s)
    scenario: NTP server srv-1 address is not empty
    FAIL: NTP server address is empty — bug is present
```

---

## Checklist

- [ ] Test name is `Test<Subject><Condition><ExpectedOutcome>` — describes what is tested and the correct result
- [ ] One-line comment above the function states the bug and what the test proves
- [ ] Minimal topology — only what is needed to reproduce
- [ ] `Then` step asserts the CORRECT expected value (not the current buggy value)
- [ ] Failure message is informative: states what was expected and what was received
- [ ] Running the test produces a `FAIL`
- [ ] No fix has been applied — the bug is only exposed
