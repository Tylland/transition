# Partner repository for outbound CI messages

> As a CI operator
> I want to maintain a registry of partner endpoints keyed by company code
> So that outbound TAF/TAP TSI messages are routed to the correct partner CI instance

The sender resolves the destination endpoint at send time by looking up the receiver's
ERA/UIC company code in the `Partners` map of `SenderConfig`. Each entry holds the
company code and the HTTPS URL of that partner's inbound web service. The map is
populated at configuration time; there is no runtime add/remove API.

## Requirement: Partner lookup by company code

> **CI-001.01-T**
>
> When a message is sent to a receiver code that exists in the partner repository,
> the sender shall route the message to the endpoint URL registered for that code.

### Scenario: Sending to a registered partner routes to the correct endpoint
```gherkin
Given the partner repository contains:
  | CompanyCode | EndpointURL                        |
  | 0000074     | https://partner-a.example.com/ci/  |
  | 0000099     | https://partner-b.example.com/ci/  |
When  a message is sent to receiver code "0000074"
Then  the HTTP POST is made to "https://partner-a.example.com/ci/"
```

### Scenario: Correct partner is selected when multiple partners are registered
```gherkin
Given the partner repository contains:
  | CompanyCode | EndpointURL                        |
  | 0000074     | https://partner-a.example.com/ci/  |
  | 0000099     | https://partner-b.example.com/ci/  |
When  a message is sent to receiver code "0000099"
Then  the HTTP POST is made to "https://partner-b.example.com/ci/"
  And no request is made to "https://partner-a.example.com/ci/"
```

## Requirement: Unknown partner is rejected before any network call

> **CI-001.02-T**
>
> When a message is sent to a receiver code that is not present in the partner
> repository, the sender shall return an error without making any HTTP request.
>
> _Comment_: The error is returned synchronously from `Send`; no partial send or
> retry is attempted. This prevents silent misrouting when a partner is misconfigured.

### Scenario: Sending to an unknown company code returns an error
```gherkin
Given the partner repository contains:
  | CompanyCode | EndpointURL                       |
  | 0000074     | https://partner-a.example.com/ci/ |
When  a message is sent to receiver code "0000099"
Then  an error is returned indicating that "0000099" is an unknown partner
  And no HTTP request is made
```

### Scenario: Sending with an empty partner repository returns an error
```gherkin
Given the partner repository is empty
When  a message is sent to receiver code "0000074"
Then  an error is returned indicating that "0000074" is an unknown partner
```
