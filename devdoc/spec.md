# Functional Requirements — Segment Service (transition)

## 1. Input Data

- **REQ-IN-01**: The system shall read Swedish railway infrastructure data from a FacilityData XML export file (schema version 1.82) produced by BIS.
- **REQ-IN-02**: The FacilityData XML shall be deserialized into a typed schema model comprising Infrastructure (nodes, links, locations), Facility (signaling facilities, geometries, other facilities), and Divisions (e.g., track numbers).
- **REQ-IN-03**: The system shall import **nodes** (noder) identified by a location number (trafikplatsId) and node number.
- **REQ-IN-04**: The system shall import **links** (länkar) identified by from-node, to-node, and a sequence number, with a length in meters.
- **REQ-IN-05**: The system shall import **locations** (platser) with signature (signatur), name, type, and PLC code, positioned on links.
- **REQ-IN-06**: The system shall import **signals** from InformationPoints, with type (INF=entry, UBL=exit, M=intermediate, MBL=intermediate block) and direction.
- **REQ-IN-07**: The system shall import **kilometer boards** (kilometertavlor) with their kilometer post values and positions on links.
- **REQ-IN-08**: The system shall import **speed profiles** with speeds for four train categories: A, B, C, and S, each with a direction of validity.
- **REQ-IN-09**: The system shall import **vertical geometries** (gradients/lutningar) with gradient values positioned on links.
- **REQ-IN-10**: The system shall import **track numbers** (spårnummer) from Divisions of type "Tracknumber", with division code (number) and description (name), spread across links.

## 2. Domain Model / Network (Nät)

- **REQ-DM-01**: The system shall build a topological network (nät) consisting of nodes (noder), links (länkar), and places (platser).
- **REQ-DM-02**: A **node** (nod) shall be identified by `NodId` = `{TrafikplatsId}:{Nummer}` and shall contain node paths (nodvägar) connecting pairs of links through the node.
- **REQ-DM-03**: A **link** (länk) shall be identified by `LankId` = `{FrånNod}:{TillNod}:{Sekvensnummer}` and have a length. A link connects exactly two nodes (från/till).
- **REQ-DM-04**: Each link shall carry collections of: signals, boards (tavlor), place centers (platsmittar), road protection (vägskydd), platforms (plattformar), track numbers (spårnummer), speed profiles (hastighetsprofiler), and gradients (lutningar).
- **REQ-DM-05**: Entities on a link shall have either a **point position** (`Lankposition` — normalized 0..1 along the link) or a **spread/extent** (`Lankutbredning` — start and end positions normalized 0..1).
- **REQ-DM-06**: All positioned and spread entities shall have a **validity direction** (gilltighetsriktning): Normal (med), Reverse (omvänd/mot), Both (båda), or None (ingen).
- **REQ-DM-07**: Link direction (länkriktning) shall be either **Med** (with, from→to) or **Mot** (against, to→from). Position and distance calculations shall be inverted when direction is Mot.
- **REQ-DM-08**: FacilityData directions (In, Against, Any) shall be mapped to domain validity directions (Normal, Omvänd, Båda).
- **REQ-DM-09**: FacilityData link positions (StartDistanceFromNode / EndDistanceFromNode) shall be normalized to fractions of link length (0..1).
- **REQ-DM-10**: Speed profiles shall be grouped by train category (A, B, C, S) per link. Each speed section (hastighetssträcka) records category, speed, and extent on the link.
- **REQ-DM-11**: The system shall support recursive network traversal via a **Visitor pattern** (`NatVisitor`), starting from a given link position and direction, proceeding through nodes and connected links.
- **REQ-DM-12**: The traversal shall respect visited-link tracking to prevent cycles, and shall terminate when the visitor signals completion.
- **REQ-DM-13**: A **signal route** (signalsträcka) is a path between two stop signals, consisting of an ordered list of link usages (länkanvändningar). Stop signals are of type INF, UBL, M, or MBL.
- **REQ-DM-14**: The `SignalstrackaVisitor` shall traverse from a start stop signal in a given direction and find the next stop signal, collecting the link usages in between to form a `Signalstracka`.
- **REQ-DM-15**: Filtering of entities by validity direction and containment within a link usage shall be supported via `GilltigaInom`, `GilltigaUtbredningar`, and `GilltigaPositioner` functions.

## 3. Segment Profiles (SFERA)

- **REQ-SP-01**: The system shall generate one SFERA segment profile per link per direction (Med and Mot), yielding two segment profiles per link.
- **REQ-SP-02**: The segment profile ID (`SP_ID`) shall be derived from the link ID and direction: `{FrånNod}:{TillNod}:{Sekvensnummer}:{Riktning}` where Riktning is 0 (Med) or 1 (Mot).
- **REQ-SP-03**: The segment profile length (`SP_Length`) shall equal the link length in meters.
- **REQ-SP-04**: The `SP_Zone` shall set `IM_ID` to 74 (Trafikverket, the Swedish infrastructure manager).
- **REQ-SP-05**: Each segment profile shall have `SP_VersionMajor` = 1, `SP_VersionMinor` = 0, and `SP_Status` = "Valid".
- **REQ-SP-06**: **Timing Points** shall be generated from place centers (platsmittar) combined with the track number valid for the current direction. The `TP_ID` shall be `{PLC}:{TrackNumber}`. The name shall be `"{TrackName}, {PlaceName}"` with language "se".
- **REQ-SP-07**: **Signals** shall be included in the segment profile for signals valid in the Med direction, with physical and object IDs set to the signal's string ID, and location adjusted for profile direction.
- **REQ-SP-08**: **Balises** shall be included for board entities (tavlor) of type "Bal" valid for the current direction, with identifier and location.
- **REQ-SP-09**: **Unprotected level crossings** shall be included for road protections (vägskydd) where `HarSkydd` = false, with name and start location.
- **REQ-SP-10**: **Platforms** shall be included as areas with start and end locations and qualifier "StartsEnds".
- **REQ-SP-11**: **Static speed profiles** shall be generated per train category. Category A shall be the default profile. Non-default profiles shall use `SpecificSSP` with `Category` = "Other_SSP_Replaces_CDSSP" and `TrainCategoryCode`.
- **REQ-SP-12**: Speed profile sections valid for the profile direction shall be sorted by start distance. The first section becomes the start speed; subsequent sections become speed changes at their respective locations.
- **REQ-SP-13**: **Gradient average** shall be computed from gradient sections (lutningar) valid for the profile direction. Gradient values shall be absolute, with direction (Uphill/Downhill) determined by sign and profile direction (inverted for Mot).
- **REQ-SP-14**: **Kilometre reference points** shall be included from kilometer boards (kilometertavlor), with identifier, location, and km reference value.
- **REQ-SP-15**: A `LineAndTrackIdentifier` element shall be included in context information.

## 4. Signal Routes (Signalsträckor / SignalRoutes)

- **REQ-SR-01**: The system shall generate signal routes by traversing the network from every stop signal in both directions (Med and Mot).
- **REQ-SR-02**: A signal route shall only be generated for directions in which the start signal is valid.
- **REQ-SR-03**: Each signal route shall record the start signal ID, end signal ID, start place (trafikplatssignatur), and end place.
- **REQ-SR-04**: Each signal route shall contain an ordered list of segment profile references (`SegmentProfileRef`), each with `SP_ID`, start position, end position, and direction.
- **REQ-SR-05**: The segment profile references shall represent the partial or full link usages that make up the path from start signal to end signal.

## 5. Track Routes (Spårsträckor / TrackRoutes)

- **REQ-TR-01**: Each track route shall have a **type**: Arrival, Departure, or Line.
- **REQ-TR-02**: An **Arrival** track route starts at an INF (entry) signal and ends at an M (intermediate) signal inside a station.
- **REQ-TR-03**: A **Departure** track route starts at an M (intermediate) signal and ends at a UBL (exit) signal inside a station.
- **REQ-TR-04**: A **Line** track route starts at a UBL (exit) signal at one place and ends at an INF (entry) signal at the next place.
- **REQ-TR-05**: Track routes shall be grouped by type, start/end place, and start/end track numbers.
- **REQ-TR-06**: Track numbers at start and end shall be determined by finding the track number (spårnummer) covering the signal's position on its link.
- **REQ-TR-07**: Each track route shall record type, start place, end place, start track, end track, start signal, end signal, and a list of signal route references.
- **REQ-TR-08**: Signal route references within a track route shall identify signal routes by their start and end signal IDs.

## 6. Output

- **REQ-OUT-01**: Segment profiles shall be serialized as individual XML files conforming to SFERA v2.00 schema, with XML header.
- **REQ-OUT-02**: Segment profile filenames shall be derived from `SP_ID` with colons replaced by hyphens, suffixed with `.xml`.
- **REQ-OUT-03**: Segment profile files shall be written to a configurable output directory (default: `testdata/sfera/`). The directory shall be created if it does not exist.
- **REQ-OUT-04**: Signal routes and track routes shall be serialized together into a single XML file (`routes.xml`) under a `<Routes>` root element with `<SignalRoutes>` and `<TrackRoutes>` sub-elements.
- **REQ-OUT-05**: The segment profile writer shall be abstracted behind a `SegmentProfileWriter` interface, allowing alternative output targets.
- **REQ-OUT-06**: All XML output shall use indented formatting (2-space indent) with a standard XML declaration header.

## 7. Algorithms & Utilities

- **REQ-ALG-01**: A Douglas-Peucker line simplification algorithm shall be available to reduce point density in track profiles while preserving shape, using an epsilon tolerance of 1.0.