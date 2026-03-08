Feature: Planning Data to Track Route Matching
  A DagligFärdplan (daily timetable) contains an ordered list of Färdplats (stops).
  The matcher converts this sequence into journey segments and looks up the
  corresponding TrackRoute for each segment using Type, Place, and Track as key.

  Background:
    Given a set of TrackRoutes indexed by Type, StartPlace, EndPlace, StartTrack, and EndTrack
    And a DagligFärdplan with an ordered sequence of Färdplats stops

  # ------------------------------------------------------------------
  # Journey segment generation from Färdplats sequence
  # ------------------------------------------------------------------

  Scenario: Journey through stops A, B, C produces Departure-Line-Arrival pattern
    Given a timetable with stops "A", "B", and "C" in order
    When the journey segments are generated
    Then the segments are in this order:
      | Type      | StartPlace | EndPlace |
      | Departure | A          | A        |
      | Line      | A          | B        |
      | Arrival   | B          | B        |
      | Departure | B          | B        |
      | Line      | B          | C        |
      | Arrival   | C          | C        |

  Scenario: Single stop produces no segments
    Given a timetable with only one stop "A"
    When the journey segments are generated
    Then no segments are produced

  Scenario: Two stops produce Departure-Line-Arrival
    Given a timetable with stops "A" and "B"
    When the journey segments are generated
    Then the segments are in this order:
      | Type      | StartPlace | EndPlace |
      | Departure | A          | A        |
      | Line      | A          | B        |
      | Arrival   | B          | B        |

  Scenario: First stop has no Arrival segment
    Given a timetable with stops "A", "B", and "C"
    When the journey segments are generated
    Then the first segment is a Departure from "A"
    And there is no Arrival at "A"

  Scenario: Last stop has no Departure segment
    Given a timetable with stops "A", "B", and "C"
    When the journey segments are generated
    Then the last segment is an Arrival at "C"
    And there is no Departure from "C"

  # ------------------------------------------------------------------
  # Track assignment per segment type
  # ------------------------------------------------------------------

  Scenario: Departure segment uses Platsspår as StartTrack and LinjespårAvgång as EndTrack
    Given a Färdplats with Signatur "Kra", Platsspår "1", and LinjespårAvgång "E"
    And it is not the last stop
    When a Departure segment is created
    Then the segment has StartPlace "Kra" and EndPlace "Kra"
    And the segment has StartTrack "1" and EndTrack "E"

  Scenario: Arrival segment uses LinjespårAnkomst as StartTrack and Platsspår as EndTrack
    Given a Färdplats with Signatur "Kia", Platsspår "3", and LinjespårAnkomst "E"
    And it is not the first stop
    When an Arrival segment is created
    Then the segment has StartPlace "Kia" and EndPlace "Kia"
    And the segment has StartTrack "E" and EndTrack "3"

  Scenario: Line segment uses previous LinjespårAvgång as StartTrack and current LinjespårAnkomst as EndTrack
    Given the previous Färdplats has Signatur "Kra" and LinjespårAvgång "E"
    And the current Färdplats has Signatur "Kia" and LinjespårAnkomst "E"
    When a Line segment is created
    Then the segment has StartPlace "Kra" and EndPlace "Kia"
    And the segment has StartTrack "E" and EndTrack "E"

  # ------------------------------------------------------------------
  # TrackRoute lookup
  # ------------------------------------------------------------------

  Scenario: TrackRoute is looked up by exact key match
    Given a TrackRoute with Type "Departure", StartPlace "Kra", EndPlace "Kra", StartTrack "1", EndTrack "E"
    When a Departure segment is matched with StartPlace "Kra", EndPlace "Kra", StartTrack "1", EndTrack "E"
    Then the segment is matched to that TrackRoute

  Scenario: No matching TrackRoute leaves segment unmatched
    Given no TrackRoute with Type "Departure", StartPlace "X", EndPlace "X", StartTrack "1", EndTrack "E"
    When a Departure segment is matched with StartPlace "X", EndPlace "X", StartTrack "1", EndTrack "E"
    Then the segment's TrackRoute is nil

  Scenario: All five key fields must match
    Given a TrackRoute with Type "Arrival", StartPlace "Kia", EndPlace "Kia", StartTrack "E", EndTrack "3"
    When a segment is matched with Type "Arrival", StartPlace "Kia", EndPlace "Kia", StartTrack "E", EndTrack "2"
    Then the segment's TrackRoute is nil
    Because the EndTrack "2" does not match "3"

  # ------------------------------------------------------------------
  # Full journey matching example
  # ------------------------------------------------------------------

  Scenario: Complete journey Kra → Kia → Kra matches all segments
    Given the following timetable for train 41721:
      | Ordningsnummer | Signatur | Platsspår | LinjespårAnkomst | LinjespårAvgång |
      | 1              | Kra      | 1         |                  | E               |
      | 2              | Kia      | 3         | E                | E               |
      | 3              | Kra      | 1         | E                |                 |
    And the following TrackRoutes exist:
      | Type      | StartPlace | EndPlace | StartTrack | EndTrack |
      | Departure | Kra        | Kra      | 1          | E        |
      | Line      | Kra        | Kia      | E          | E        |
      | Arrival   | Kia        | Kia      | E          | 3        |
      | Departure | Kia        | Kia      | 3          | E        |
      | Line      | Kia        | Kra      | E          | E        |
      | Arrival   | Kra        | Kra      | E          | 1        |
    When the timetable is matched to TrackRoutes
    Then all 6 segments are matched
    And no segments have a missing TrackRoute

  Scenario: Journey with missing Departure route
    Given the following timetable for train 12345:
      | Ordningsnummer | Signatur | Platsspår | LinjespårAnkomst | LinjespårAvgång |
      | 1              | Kv       | 1         |                  | E               |
      | 2              | Rut      | 1         | E                |                 |
    And no Departure TrackRoute exists for Kv with StartTrack "1" and EndTrack "E"
    When the timetable is matched to TrackRoutes
    Then the Departure segment for "Kv" has a missing TrackRoute
    And the Line and Arrival segments may still be matched

  # ------------------------------------------------------------------
  # Match result summary
  # ------------------------------------------------------------------

  Scenario: Match result contains train identification
    Given a DagligFärdplan with Tågläge 41721 and Utgångsdag "2022-06-30"
    When the timetable is matched
    Then the MatchResult has Tågläge 41721
    And the MatchResult has Utgångsdag "2022-06-30"

  Scenario: Match result counts matched and unmatched segments
    Given a journey producing 6 segments
    And 5 segments are matched to TrackRoutes
    And 1 segment has no matching TrackRoute
    Then the MatchResult reports Matched 5 and Missing 1
