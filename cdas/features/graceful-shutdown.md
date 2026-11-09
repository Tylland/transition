# Shutting Down the Segment Service

> **As a** segment service operator  
> **I want** the service to stop on an interrupt or termination signal  
> **So that** background processing stops and the NATS connection is closed  

## Scenario: Stop after an interrupt or termination signal
```gherkin
Given the segment service is monitoring facility-data files
When it receives an interrupt or termination signal
Then it cancels the monitoring context
  And the file-monitoring loop stops
  And the service closes its NATS connection as it exits
```

## Scenario: Finish the current file scan before exiting
```gherkin
Given the service is processing a file during a scan
When it receives an interrupt or termination signal
Then it exits the monitoring loop after the current scan and processing call return
```
