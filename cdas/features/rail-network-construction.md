# Building a Rail Network from Facility Data

> As a segment service
> I want to convert facility data into a rail network
> So that segment profiles and routes can be generated from the network

## Scenario: Build a network from facility data
```gherkin
Given a facility-data document with infrastructure, facility, and division data
When the service converts the document into a network
Then it creates the network's nodes and links
  And builds node routes
  And adds supported places, track divisions, signals, kilometre boards, speed profiles, and gradients
```

## Scenario: Ignore facility data that references an unknown link
```gherkin
Given facility data contains a positioned object whose link is not in the network
When the service builds the network
Then it does not attach that object to a link
```
