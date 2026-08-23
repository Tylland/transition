 # Story
 <!-- @api @integration @receiver -->
 > As an integration operator
 > I want the receiver service to accept, validate, and acknowledge incoming messages
 > So that downstream processing receives well-formed, authorized payloads

 The receiver listens on configured HTTP/S endpoints and accepts messages (SOAP and JSON) from external partners. It validates payloads against the appropriate schema, enforces authentication and idempotency, persists accepted messages to the incoming store, and emits acknowledgements. Transient errors should be retried; unrecoverable failures are routed to a dead-letter queue for operator review.

 ### Scenario: Background
 ```gherkin
 Given the receiver service is running at "http://localhost:8080/receiver"
   And the schema repository is populated with the current XSD/JSON schemas
   And the incoming message store and message broker are reachable
   And a sender credential "test-sender-key" exists and is valid
 ```

 ## Requirement: Accept and persist valid messages

 > **RCV-100.01-T**
 >
 > The receiver SHALL accept well-formed messages and persist them to the incoming message store for downstream processing.
 >
 > _Comment_: Applies to SOAP and HTTP JSON payloads; validation performed against configured XSD/JSON schemas.

 ### Scenario: Receiver accepts and persists a valid message
 ```gherkin
 Given a well-formed "planningdata" message payload
   And the sender provides credential "test-sender-key"
 When  the sender POSTs the message to "/receiver"
 Then  the receiver responds with HTTP 200
   And the message is persisted to the incoming store
   And an acknowledgement with status "accepted" is returned
 ```

 ### Scenario: Receiver rejects message with invalid schema
 ```gherkin
 Given a message payload that does not conform to the "planningdata" schema
   And the sender provides credential "test-sender-key"
 When  the sender POSTs the message to "/receiver"
 Then  the receiver responds with HTTP 400
   And the response contains an error indicating "schema validation failed"
   And the message is not persisted
 ```

 ## Requirement: Authenticate and authorize senders

 > **RCV-100.02-T**
 >
 > The receiver SHALL authenticate and authorize senders before accepting messages.
 >
 > _Comment_: Authentication may use API keys, mTLS or token-based auth depending on deployment.

 ### Scenario: Valid credential allows processing
 ```gherkin
 Given a well-formed message payload
   And the sender provides credential "test-sender-key"
 When  the sender POSTs the message to "/receiver"
 Then  the receiver responds with HTTP 200
   And the message is processed and persisted
 ```

 ### Scenario: Missing or invalid credential is rejected
 ```gherkin
 Given a well-formed message payload
   And the sender provides an invalid credential "bad-key"
 When  the sender POSTs the message to "/receiver"
 Then  the receiver responds with HTTP 401
   And the message is not persisted
 ```

 ## Requirement: Idempotency and duplicate detection

 > **RCV-100.03-T**
 >
 > The receiver SHALL detect duplicate messages (by `Message-Id` header or a configured deduplication key) and ensure each logical message is processed at most once.
 >
 > _Comment_: The deduplication window and storage backend are deployment-configurable.

 ### Scenario: Duplicate message is ignored
 ```gherkin
 Given the message with Message-Id "msg-123" has already been processed
   And the sender provides credential "test-sender-key"
 When  the sender POSTs the same message to "/receiver"
 Then  the receiver responds with HTTP 200
   And the receiver does not persist a second copy
   And the acknowledgement indicates "duplicate"
 ```

 ### Scenario: New message with unique id is processed
 ```gherkin
 Given a message with Message-Id "msg-456" that has not been seen before
   And the sender provides credential "test-sender-key"
 When  the sender POSTs the message to "/receiver"
 Then  the receiver persists the message
   And an acknowledgement with status "accepted" is returned
 ```

 ## Requirement: Retry and dead-letter handling

 > **RCV-100.04-T**
 >
 > When transient processing failures occur, the receiver SHALL retry delivery according to the configured retry policy and move messages exceeding retry limits to a dead-letter queue.
 >
 > _Comment_: Retry policy (attempts, backoff) is configurable per deployment.

 ### Scenario: Transient storage failure triggers retry
 ```gherkin
 Given the incoming store returns a transient error on the first attempt
   And the sender provides credential "test-sender-key"
 When  the sender POSTs a valid message to "/receiver"
 Then  the receiver responds with HTTP 202
   And the message is scheduled for retry
   And the message is eventually persisted when the transient error resolves
 ```

 ### Scenario: Permanent failure moves message to dead-letter queue
 ```gherkin
 Given the incoming store returns a permanent error for the message
   And the configured retry limit is 3 attempts
 When  the sender POSTs a valid message to "/receiver"
 Then  the receiver responds with HTTP 500
   And the message is moved to the dead-letter queue after 3 failed attempts
   And an error notification is emitted for operator review
 ```

 ## Requirement: Acknowledgement format

 > **RCV-100.05-T**
 >
 > The receiver SHALL return an acknowledgement containing the processing result and a reference identifier that downstream systems can use to correlate processing.
 >
 > _Comment_: Acknowledgement structure is documented in the API spec; include `message_id`, `status`, and optional `error` fields.

 ### Scenario: Acknowledgement contains message id and status
 ```gherkin
 Given a well-formed message payload with Message-Id "msg-789"
   And the sender provides credential "test-sender-key"
 When  the sender POSTs the message to "/receiver"
 Then  the receiver responds with HTTP 200
   And the acknowledgement contains "message_id: msg-789" and "status: accepted"
 ```

 ## Scenario: Accepts multiple supported payload types
 ```gherkin
 Given a well-formed "<payload_type>" message payload
   And the sender provides credential "test-sender-key"
 When  the sender POSTs the message to "/receiver"
 Then  the receiver responds with HTTP 200
   And the message is persisted to the incoming store

 Examples:
 <payload_type>
 |planningdata|
 |facilitydata|
 ```
