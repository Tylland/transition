# Generating and Publishing Signal Routes

> As a segment service
> I want to generate signal routes from the rail network
> So that consumers can retrieve routes between stop signals

## Scenario: Generate routes between reachable stop signals
```gherkin
Given a rail network with stop signals and traversable links
When the service generates signal routes
Then it traverses from each stop signal in every valid direction
  And creates a route when it reaches the next stop signal
```

## Scenario: Include segment profile references in a signal route
```gherkin
Given a generated signal route traverses links with generated segment profiles
When the service constructs the signal route
Then it includes a reference for each matching segment profile along the route
  And each reference includes the profile ID, version, position range, and direction
```

## Scenario: Publish signal routes to the key-value bucket
```gherkin
Given signal routes have been generated
When the service writes the routes
Then it stores each route as XML in the "SIGNAL_ROUTES" bucket
  And uses the route ID as the key
```
