---
name: feature
description: >
  Use this skill whenever the user wants to create, scaffold, or improve a feature file
  (.md). Triggers include: "create a feature file",
  "write a scenario", "add a scenario outline", "BDD test", "Gherkin", "Given/When/Then",
  "acceptance test", or any request to document behaviour in
  plain English. Also use when converting existing test cases, user stories, or acceptance
  criteria into the markdown+Gherkin format. Use this skill proactively whenever feature
  files, step definitions are mentioned — even if the user just
  says "write me a test for X". Output is always a .md file, NOT a .feature file.
---

# Feature File Skill

## Overview

This skill guides the creation of well-structured feature files in **Markdown format** (`.md`).
Each file documents a user story with its scenarios using fenced Gherkin blocks.
Follow the **Decision Tree** to choose the right construct, then use the **Templates** to fill it in.

---

## Decision Tree — What to write

```
Does the feature have multiple distinct requirements?
  └─ YES → group scenarios under ## Requirement: headings
           (scenarios become ### Scenario: instead of ## Scenario:)

Does the scenario have fixed inputs?
  ├─ YES → use Scenario
  └─ NO / multiple input rows → use Scenario with Examples table

Does every scenario in this feature share the same setup steps?
  └─ YES → add a Background scenario at the top

Does the feature belong to a sprint, story, or test category?
  └─ YES → add tag comments below the Story heading

Does a step pass a multi-field payload?
  └─ YES → use an inline table inside the Gherkin block
```

---

## File Structure

The output is a `.md` file. Use the flat layout when all scenarios belong to one coherent
story. Use the grouped layout when the story has multiple named requirements.

**Flat layout** (no requirements):
```
# Story          ← H1, one per file
> blockquote     ← user story (As a / I want / So that)

plain text       ← optional: background context, domain explanation, key concepts
                    (one or more paragraphs; no heading needed)

## Scenario:     ← H2, one per scenario
  ```gherkin     ← fenced code block with Gherkin steps
  step
  Examples:      ← only for parameterised scenarios
  <col>  ...
  |val|  ...
  ```
```

**Grouped layout** (with requirements):
```
# Story          ← H1, one per file
> blockquote     ← user story (As a / I want / So that)

plain text       ← optional: background context, domain explanation, key concepts
                    (one or more paragraphs; no heading needed)

## Requirement:  ← H2, one per requirement — names the rule being tested

### Scenario:    ← H3, one per scenario (demoted from H2)
  ```gherkin
  step
  ```
```

---

## Templates

### 1 — Basic Scenario

```markdown
## Scenario: <Descriptive title in plain English>
```gherkin
Given <the system or user is in a known state>
  And <additional context if needed>          
When  <the user or system performs an action> 
Then  <the expected observable outcome>       
  And <any additional expected outcome>
```
```

**Rules:**
- One behaviour per scenario — never combine two "When" blocks.
- "Given" sets context; "When" triggers action; "Then" asserts outcome.
- "And" continues the preceding keyword — do not change the semantic role.
- Indent "And" with two spaces to visually subordinate it to its parent keyword.

---

### 2 — Scenario with Examples (data-driven)

```markdown
## Scenario: <Title describing the varying behaviour>        
```gherkin
Given <A> is <param_a>
  And <B> is <param_b>
When  <action>
Then  result equals <param_result>

Examples:
<param_a>   <param_b>   <param_result>
|value1     |value2     |value3     |  
|value4     |value5     |value6     |  
```
```

**Rules:**
- Column headers must exactly match the `<placeholder>` names used in the steps.
- Provide at least two data rows; use a plain Scenario if only one row exists.
- Pad column values so all pipe characters align vertically.

---

### 3 — Background (shared setup)

Represent shared setup as a dedicated scenario at the top of the file using
the keyword `Background` in the scenario title.

```markdown
## Scenario: Background
```gherkin
Given the application is running on "https://staging.example.com"
  And the database has been seeded with fixture data
  And I am authenticated as a standard user
```
```

**Rules:**
- Background steps run before every scenario — keep them fast.
- Never include "When" or "Then" steps in Background.
- If only one scenario needs the setup, duplicate the steps there instead.

---

### 4 — Negative / failure scenario

```markdown
## Scenario: <Title describing the failure case>             
```gherkin
Given <the system is in a valid initial state>
When  <an action that should fail is performed>
Then  <the expected error outcome>
  And <the system remains in a safe state>                   
```
```


---

### 5 — Data Table (structured step input)

```markdown
## Scenario: <Title>                                       
```gherkin
Given I have the following user details:                   
  | Field     | Value                  |
  | firstName | Jane                   |
  | lastName  | Smith                  |
  | email     | jane.smith@example.com |
  | role      | Editor                 |
When  I submit the registration form                       
Then  the account should be created successfully           
  And a confirmation email should be sent to "jane.smith@example.com"
```
```

**When to use:** multi-field input that would make a step unreadably long as inline args.

---

### 6 — Requirement grouping

Use when a feature has multiple distinct requirements, each with several scenarios.
The `## Requirement:` heading names the rule; `### Scenario:` headings name the cases.
Each requirement heading is followed by a blockquote block that states the requirement
number, the normative text, and an optional comment.

Requirement numbers follow the format `<prefix>-<section>.<seq>-<type>` where:
- `<prefix>` — system/component abbreviation, e.g. `TS`
- `<section>` — three-digit code: chapter + section digits, e.g. `302` = chapter 3, section 2
- `<seq>` — two-digit sequential number within the section, e.g. `03`
- `<type>` — `A` (administrative) or `T` (technical)

Example: `TS-302.03-T`

```markdown
## Requirement: <Plain-English statement of the rule>

> **<prefix>-<section>.<seq>-<type>**
>
> <Requirement text>
>
> _Comment_: <Comment>

### Scenario: <Happy-path title>
```gherkin
Given <context>
When  <action>
Then  <outcome>
```

### Scenario: <Failure / edge-case title>
```gherkin
Given <context>
When  <action that should fail>
Then  <expected error outcome>
```
```

**Rules:**
- The `## Requirement:` title should match the corresponding requirement in the spec.
- Scenarios are demoted to `### Scenario:` (H3) — never mix H2 and H3 Scenario headings
  in the same file.
- Background still uses `### Scenario: Background` (H3) at the top of the file when
  requirement grouping is in use.
- The `_Comment_:` line is optional — omit it when no clarification is needed.

#### When to use `_Comment_:`

A comment is a short note attached to the normative requirement text. It records
information that the requirement text alone does not make obvious — typically a design
decision, a scope boundary, or a constraint on how the requirement must be implemented
or tested. It does **not** restate the requirement; it adds context that would otherwise
be lost.

Use a comment when any of the following is true:

**Design decision** — the implementation makes a non-obvious algorithmic choice that
testers need to understand:

> _Comment_: `HasCommonNetworks` is symmetric by design. Even when only a one-directional
> gateway A→B exists, device A's reachable set contains B and device B's reachable set
> contains B, so both devices share B — enabling tunnel formation without a reverse gateway.

**Scope boundary** — the requirement deliberately excludes something a reader might expect
to be included:

> _Comment_: This check applies to the Virtual Network as a whole; a mixed-zone VN is
> skipped entirely rather than filtering individual port pairs within it.

**Dependency or pre-condition** — the requirement only makes sense in the context of
another requirement or a domain constraint:

> _Comment_: Requires TG-301.01-T. The AdminTunnel flag is set at tunnel-creation time;
> it cannot be changed after the tunnel object is constructed.

**Non-obvious test implication** — the way a requirement must be tested differs from
what the wording suggests:

> _Comment_: A disabled port pair has `enabled=false` and a nil crypto port. The eligibility
> check must guard against nil before dereferencing the crypto network — this scenario
> exercises that guard.

Omit `_Comment_:` when the requirement text is self-contained and the scenarios make
the intent clear without additional explanation.

---

## Tag Conventions

Place tags as a comment line directly below the `# Story` heading.

```markdown
# Story
<!-- @smoke @regression @epic-auth @PROJ-1234 -->
> As a ...
```

| Category       | Example tags                                        |
|----------------|-----------------------------------------------------|
| Test layer     | `@ui` `@api` `@unit` `@integration`                 |
| Execution tier | `@smoke` `@regression` `@full-suite`                |
| Test type      | `@happy-path` `@negative` `@edge-case` `@boundary`  |
| Status         | `@wip` `@manual` `@ignore` `@flaky`                 |
| Traceability   | `@PROJ-1234` `@sprint-42` `@epic-auth`              |

---

## Naming Conventions

| Element          | Convention                                            | Example                                   |
|------------------|-------------------------------------------------------|-------------------------------------------|
| Story heading    | Short imperative phrase — what the system does        | `# Story`                                 |
| User story       | As a / I want / So that in blockquote lines           | `> As a Network Administrator`            |
| Scenario title   | Full sentence — readable without the story context    | `Successful login with valid credentials` |
| Outline title    | Present-tense description of the varying behaviour    | `Login redirects based on user role`      |
| Step text        | Plain English; avoid technical jargon in Given/When   | `Given I am on the login page`            |
| Parameters       | snake_case inside angle brackets                      | `<user_role>`                             |
| File name        | kebab-case matching the story subject                 | `user-authentication.md`                  |

---

## Step Definition Skeleton (Go)

Each Go test function maps to one Gherkin scenario. The framework is in `scenario.go`
(package `testing`). The chain is: `Given → And → When → Then → And → Execute`.
**Always call `.Execute()` at the end** — it prints the scenario report to stdout.

### Framework API at a glance

| Method | Receiver | Returns | Purpose |
|---|---|---|---|
| `NewScenario(t)` | — | `*Scenario` | Create scenario; title derived from `t.Name()` |
| `NewStepResult(format, args...)` | — | `*StepResult` | Create a named step |
| `.Success()` | `*StepResult` | `*StepResult` | Mark step passed |
| `.WithError(err)` | `*StepResult` | `*StepResult` | Mark step failed with error |
| `.WithErrorf(fmt, args...)` | `*StepResult` | `*StepResult` | Mark step failed with formatted message |
| `.WithErrorMessage(msg)` | `*StepResult` | `*StepResult` | Mark step failed with string message |
| `.With(other)` | `*StepResult` | `*StepResult` | Propagate error from another StepResult |
| `.Background(sr)` | `*Scenario` | `*GivenStepBlock` | Shared setup step (shown as "Background") |
| `.Given(sr)` | `*Scenario` | `*GivenStepBlock` | First Given step |
| `.Given(sr)` / `.And(sr)` | `*GivenStepBlock` | `*GivenStepBlock` | Additional Given/And steps |
| `.When(sr)` | `*GivenStepBlock` | `*WhenStepBlock` | Trigger action |
| `.Then(sr)` | `*WhenStepBlock` | `*ThenStepBlock` | First assertion |
| `.And(sr)` | `*ThenStepBlock` | `*ThenStepBlock` | Additional assertions |
| `.Execute()` | `*ThenStepBlock` | — | Print scenario report (required) |

---

### Basic scenario

```go
func TestUserLogin(t *testing.T) {
    s := testing.NewScenario(t)

    s.Given(navigateToLoginPage()).
     And(enterCredentials("jane@example.com", "Secret123!")).
     When(clickLoginButton()).
     Then(redirectedToDashboard()).
     And(welcomeMessageContains("Jane")).
     Execute()
}

func navigateToLoginPage() *testing.StepResult {
    sr := testing.NewStepResult("navigate to login page")
    // perform navigation
    return sr.Success()
}

func enterCredentials(username, password string) *testing.StepResult {
    sr := testing.NewStepResult("enter credentials for %s", username)
    // fill form fields
    return sr.Success()
}

func clickLoginButton() *testing.StepResult {
    sr := testing.NewStepResult("click Login button")
    return sr.Success()
}

func redirectedToDashboard() *testing.StepResult {
    sr := testing.NewStepResult("redirected to dashboard")
    return sr.Success()
}

func welcomeMessageContains(name string) *testing.StepResult {
    sr := testing.NewStepResult("welcome message contains %q", name)
    return sr.Success()
}
```

---

### Negative / failure scenario

```go
func TestUserLoginInvalidCredentials(t *testing.T) {
    s := testing.NewScenario(t)

    err := performLogin("bad@user.com", "wrong")

    s.Given(navigateToLoginPage()).
     When(testing.NewStepResult("enter invalid credentials").Success()).
     Then(func() *testing.StepResult {
         sr := testing.NewStepResult("error message is shown")
         if err == nil {
             return sr.WithErrorMessage("expected login to fail but it succeeded")
         }
         return sr.Success()
     }()).
     Execute()
}
```

---

### Background + shared setup

```go
func TestCreateUserAccount(t *testing.T) {
    s := testing.NewScenario(t)

    s.Background(seedDatabase()).
     Given(openRegistrationForm()).
     And(fillUserDetails("Jane", "Smith", "jane@example.com")).
     When(submitRegistrationForm()).
     Then(accountCreated()).
     And(confirmationEmailSentTo("jane@example.com")).
     Execute()
}

func seedDatabase() *testing.StepResult {
    sr := testing.NewStepResult("database seeded with fixture data")
    return sr.Success()
}
```

---

### Data-driven scenarios (Examples equivalent)

```go
func TestLoginRedirectsByRole(t *testing.T) {
    cases := []struct {
        role        string
        username    string
        password    string
        landingPage string
    }{
        {"Admin",  "admin@example.com",  "Admin@123",  "/admin"},
        {"Editor", "editor@example.com", "Editor@123", "/editor"},
        {"Viewer", "viewer@example.com", "Viewer@123", "/home"},
    }

    for _, tc := range cases {
        tc := tc
        t.Run("LoginAs"+tc.role, func(t *testing.T) {
            s := testing.NewScenario(t)

            s.Given(navigateToLoginPage()).
             And(enterCredentials(tc.username, tc.password)).
             When(clickLoginButton()).
             Then(landingPageIs(tc.landingPage)).
             Execute()
        })
    }
}

func landingPageIs(expected string) *testing.StepResult {
    sr := testing.NewStepResult("landed on %s", expected)
    return sr.Success()
}
```

---

## Setup and Teardown (Go)

There are no hooks in this framework. Use standard Go mechanisms instead.

```go
// TestMain — one-time setup and teardown for the entire test binary
func TestMain(m *testing.M) {
    startTestServer()
    code := m.Run()
    stopTestServer()
    os.Exit(code)
}

// Per-test teardown via t.Cleanup — AfterScenario equivalent
func TestSomething(t *testing.T) {
    db := openTestDatabase(t)
    t.Cleanup(func() {
        db.Rollback()
        db.Close()
    })

    s := testing.NewScenario(t)
    s.Given(seedDatabase()).
     When(performAction()).
     Then(outcomeIsCorrect()).
     Execute()
}

// Shared fixture helper — call at top of each test that needs it
func setupFeatureFixtures(t *testing.T) {
    t.Helper()
    // shared fixture setup
    t.Cleanup(func() {
        // shared fixture teardown
    })
}
```

---

## Common Mistakes to Avoid

| Mistake                                       | Fix                                                             |
|-----------------------------------------------|-----------------------------------------------------------------|
| Multiple "When" steps in one scenario         | Split into separate scenarios                                   |
| Assertions inside Background                  | Move to individual scenario "Then" steps                        |
| Technical jargon in Gherkin steps             | Write steps as business language; hide detail in step functions |
| Examples scenario with a single data row      | Use a plain Scenario instead                                    |
| Reusing the same scenario title               | Titles must be unique within a feature file                     |
| Placeholder in step doesn't match column name | `<param_name>` in step and Examples header must match exactly   |

---

## Quick Reference Checklist

Before saving a feature file, verify:

- [ ] File is saved as `kebab-case-title.md`
- [ ] `# Story` heading is present with a blockquote user story
- [ ] User story follows As a / I want / So that format
- [ ] If the file has `## Requirement:` headings, all scenarios use `### Scenario:` (H3)
- [ ] If there are no `## Requirement:` headings, all scenarios use `## Scenario:` (H2)
- [ ] `## Requirement:` titles match the corresponding spec requirement names
- [ ] Every scenario has a unique, descriptive title
- [ ] Each scenario tests exactly one behaviour
- [ ] Examples scenarios have at least two data rows
- [ ] Background scenario contains only Given/And steps (no When or Then)
