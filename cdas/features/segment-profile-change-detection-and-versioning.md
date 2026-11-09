# Detecting and Versioning Segment Profile Changes

> As a segment service operator
> I want unchanged segment profiles to be skipped and changed profiles versioned
> So that consumers receive profile updates only when profile content changes

The writer calculates a checksum from the profile's serialized content fields and stores profile metadata in the segment-profile key-value bucket.

## Scenario: Skip an unchanged segment profile
```gherkin
Given a segment profile with the same checksummed content has already been published
When the writer processes the profile again
Then it returns the profile key without publishing another profile version
```

## Scenario: Publish a changed segment profile with an incremented major version
```gherkin
Given a segment profile has been published before
When the checksum of its content changes
Then the writer publishes the profile with a major version incremented from the stored version
  And updates the stored profile metadata
```

## Scenario: Publish a profile with its initial major version
```gherkin
Given no profile metadata is stored for a segment profile
When the writer processes the profile
Then it publishes the profile with major version "1"
  And stores profile metadata for subsequent change detection
```
