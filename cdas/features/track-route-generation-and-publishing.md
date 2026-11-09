# Generating and Publishing Track Routes

> As a segment service
> I want to group signal routes into track routes
> So that consumers can look up routes by their station and track context

## Scenario: Classify a signal route
```gherkin
Given a signal route with start and end signals
When the service classifies the route
Then it identifies the track route as an arrival, departure, or line route based on the signal types
```

## Scenario: Group signal routes into track routes
```gherkin
Given generated signal routes with resolved places and tracks
When the service builds track routes
Then it groups routes by type, places, direction place, and start and end tracks
  And each group becomes one track route referencing its signal routes
```

## Scenario: Publish track routes to the key-value bucket
```gherkin
Given track routes have been generated
When the service writes the routes
Then it stores each route as XML in the "TRACK_ROUTES" bucket
  And uses the route ID as the key
```
