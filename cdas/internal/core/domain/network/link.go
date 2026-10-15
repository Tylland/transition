package network

import "fmt"

// Link represents a link in the network.
type Link struct {
	Id     LinkId
	From   *Node
	To     *Node
	Length float64

	Signals        Entities[*Signal]
	Boards         Entities[*Board]
	PlaceCenters   Entities[*PlaceCenter]
	Levelcrossings Entities[*Levelcrossing]
	Platforms      Entities[*Platform]
	StationTracks  Entities[*StationTrack]
	LineTracks     Entities[*LineTrack]
	SpeedProfiles  Entities[*SpeedProfile]
	Slopes         Entities[*Slope]
}

// NewLink creates a new Link.
func NewLink(id LinkId, from, to *Node, length float64) *Link {
	return &Link{
		Id:     id,
		From:   from,
		To:     to,
		Length: length,
	}
}

// AddSpeedSegment adds a speed section to the appropriate speed profile,
// creating the profile if it doesn't already exist.
func (l *Link) AddSpeedSegment(hs *SpeedSegment) {
	var profile *SpeedProfile
	for _, sp := range l.SpeedProfiles.Items() {
		if sp.Category == hs.Category {
			profile = sp
			break
		}
	}

	if profile == nil {
		profile = NewSpeedProfile(hs.Category)
		l.SpeedProfiles.Add(profile)
	}

	profile.Segments.Add(hs)
}

// GetNode returns the node at the end of the link in the given direction.
func (l *Link) GetNode(direction LinkDirection) *Node {
	switch direction {
	case LinkDirectionWith:
		return l.To
	case LinkDirectionAgainst:
		return l.From
	default:
		panic("LinkDirection unknown")
	}
}

// GetOtherNode returns the other node of the link (not the given one).
func (l *Link) GetOtherNode(node *Node) *Node {
	if node == l.From {
		return l.To
	}
	if node == l.To {
		return l.From
	}
	panic("Link do not connect to node!")
}

// GetLinkDirection returns the direction of the link relative to the given node.
func (l *Link) GetLinkDirection(node *Node) LinkDirection {
	if node == l.From {
		return LinkDirectionWith
	}
	if node == l.To {
		return LinkDirectionAgainst
	}
	panic("Link do not connect to node!")
}

// String returns a string representation.
func (l *Link) String() string {
	return fmt.Sprintf("%v->%v %v m", l.From, l.To, l.Length)
}
