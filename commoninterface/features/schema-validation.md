# XML schema validation for CI messages

> As a CI operator
> I want all TAF/TAP TSI messages to be validated against their XSD schema before sending and after receiving
> So that malformed messages are rejected at the boundary and never forwarded to partner systems or internal consumers

Each TAF/TAP TSI message carries a `MessageType` element in its header whose value is the
root element name of the payload (e.g. `TrainRunningInformationMessage`). The schema
repository maps each `MessageType` to an XSD file loaded from a configurable directory on
startup. Validation is performed in-process using the loaded schema set; no network access
is required at runtime.

The `MessageTypeVersion` element in the header is recorded alongside `MessageType` but
does not affect schema selection — the repository holds one XSD per `MessageType`.

## Requirement: Schema repository resolves XSD files from a configurable directory

> **CI-002.01-T**
>
> The schema repository shall load all `.xsd` files found in the configured directory at
> initialisation time and index them by the `MessageType` name derived from the file name.
>
> _Comment_: The directory path must be provided through configuration so that operators can
> point at different schema bundles (e.g. a local copy of `TD_105_xsd/`) without recompiling.
> File name to `MessageType` mapping follows the TAF/TAP TSI convention: the file name stem
> equals the root element name, e.g. `TrainRunningInformationMessage.xsd`.

### Scenario: Repository loads schemas from the configured directory
```gherkin
Given the schema directory contains:
  | File                                  |
  | TrainRunningInformationMessage.xsd    |
  | PathRequestMessage.xsd                |
When  the schema repository is initialised with that directory
Then  the repository contains schemas for:
  | MessageType                        |
  | TrainRunningInformationMessage     |
  | PathRequestMessage                 |
```

### Scenario: Repository is empty when the directory contains no XSD files
```gherkin
Given the schema directory contains no XSD files
When  the schema repository is initialised with that directory
Then  the repository contains no schemas
```

### Scenario: Initialisation fails when the configured directory does not exist
```gherkin
Given the schema directory path "/nonexistent/schemas" does not exist
When  the schema repository is initialised with that path
Then  an error is returned indicating the directory was not found
```

## Requirement: Outbound messages are validated before sending

> **CI-002.02-T**
>
> Before posting a message to a partner CI, the sender shall validate the XML payload
> against the XSD schema registered for the `MessageType` found in the message header.
> If validation fails the message shall not be sent and an error shall be returned to the caller.

### Scenario: A valid outbound message is sent without error
```gherkin
Given the schema repository contains a schema for "TrainRunningInformationMessage"
  And the partner "0000099" is registered
When  a "TrainRunningInformationMessage" payload that conforms to the schema is sent to "0000099"
Then  the message is posted to the partner endpoint
  And no validation error is returned
```

### Scenario: An outbound message that fails schema validation is rejected before sending
```gherkin
Given the schema repository contains a schema for "TrainRunningInformationMessage"
  And the partner "0000099" is registered
When  a "TrainRunningInformationMessage" payload that violates the schema is sent to "0000099"
Then  an error is returned describing the schema violation
  And no HTTP request is made to the partner endpoint
```

### Scenario: Sending a message whose type has no registered schema returns an error
```gherkin
Given the schema repository does not contain a schema for "PathRequestMessage"
  And the partner "0000099" is registered
When  a "PathRequestMessage" payload is sent to "0000099"
Then  an error is returned indicating no schema is registered for "PathRequestMessage"
  And no HTTP request is made to the partner endpoint
```

## Requirement: Inbound messages are validated after receiving

> **CI-002.03-T**
>
> When a message is received from a partner CI, the receiver shall validate the XML payload
> against the XSD schema registered for the `MessageType` found in the message header.
> If validation fails the message shall be rejected and a NACK acknowledgement returned to the sender.
>
> _Comment_: Validation happens before the payload is forwarded to any internal consumer,
> ensuring that downstream components only ever process schema-conformant messages.

### Scenario: A valid inbound message is accepted and forwarded
```gherkin
Given the schema repository contains a schema for "TrainRunningInformationMessage"
When  a "TrainRunningInformationMessage" payload that conforms to the schema is received
Then  the message is forwarded to the internal consumer
  And an ACK acknowledgement is returned to the sender
```

### Scenario: An inbound message that fails schema validation is rejected
```gherkin
Given the schema repository contains a schema for "TrainRunningInformationMessage"
When  a "TrainRunningInformationMessage" payload that violates the schema is received
Then  the message is not forwarded to the internal consumer
  And a NACK acknowledgement is returned to the sender
```

### Scenario: An inbound message whose type has no registered schema is rejected
```gherkin
Given the schema repository does not contain a schema for "PathRequestMessage"
When  a "PathRequestMessage" payload is received
Then  the message is not forwarded to the internal consumer
  And a NACK acknowledgement is returned to the sender
```
