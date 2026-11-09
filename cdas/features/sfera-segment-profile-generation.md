# Generating SFERA Segment Profiles

> As a segment service
> I want to generate SFERA segment profiles from the rail network
> So that consumers can use infrastructure data in each travel direction

## Scenario: Generate profiles in both directions
```gherkin
Given a rail network containing links
When the service processes the network
Then it generates a SFERA segment profile for every link in each direction
```

## Scenario: Include supported network data in profiles
```gherkin
Given a link has applicable place centers, signals, level crossings, platforms, speed data, gradients, or kilometre boards
When the service generates its segment profiles
Then each profile includes the applicable timing points, signals, level-crossing stops, platforms, speed profiles, gradients, and kilometre references
  And directional locations reflect the profile's travel direction
```

## Scenario: Persist profiles before generating routes
```gherkin
Given the service has generated segment profiles for a network
When it processes the network
Then it writes the segment profiles before writing signal routes and track routes
```
