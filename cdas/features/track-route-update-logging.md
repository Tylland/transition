# Logging Track Route Bucket Entries

> As a segment service operator
> I want the service to observe entries in the track-route bucket
> So that I can see the keys reported by the bucket watcher

## Scenario: Log keys received by the bucket watcher
```gherkin
Given the service is watching all entries in the "TRACK_ROUTES" bucket
When the watcher reports an entry
Then the service prints that entry's key
```

## Scenario: Create the bucket when it does not exist
```gherkin
Given the "TRACK_ROUTES" bucket does not exist
When the service creates its track-route logger
Then it creates the bucket and starts watching its entries
```
