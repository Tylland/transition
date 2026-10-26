package visitor

import (
	"testing"

	dm "github.com/tylland/transition/cdas/internal/core/domain/network"
	"github.com/tylland/transition/specs"
)

// ── Story declarations ────────────────────────────────────────────────────────

var findPlaceStory = specs.NewStory(specs.UserStory{
	AsA:     "builder",
	IWantTo: "find adjecent places",
	SoThat:  "",
})

var b = &NetworkBuilder{}
var foundPlaceSignature = ""

func FindingNextPlaceFrom(id string) *specs.StepResult {
	res := specs.NewStepResult("finding next place from signal %q", id)

	signalId, err := dm.ParseSignalId(id)

	if err != nil {
		return res.WithError(err)
	}

	findingPlace(signalId, false)

	return res.Success()
}

func FindingPreviousPlaceFrom(id string) *specs.StepResult {
	res := specs.NewStepResult("finding previous place from signal %q", id)

	signalId, err := dm.ParseSignalId(id)

	if err != nil {
		return res.WithError(err)
	}

	findingPlace(signalId, true)

	return res.Success()
}

func findingPlace(signalId *dm.SignalId, reverse bool) {
	b.network.BuildNodeRoutes()

	signal := b.network.FindSignal(signalId)

	v := NewPlaceFinder(signal, reverse)
	dm.Traverse(v)

	foundPlaceSignature = v.PlaceSignature()
}

func FoundPlaceIs(signature string) *specs.StepResult {
	sr := specs.NewStepResult("found place is %q", signature)

	if foundPlaceSignature != signature {
		return sr.WithErrorf("expected %q, got %q", signature, foundPlaceSignature)
	}

	return sr.Success()
}

func TestNextPlaceFromDeparture(t *testing.T) {
	findPlaceStory.Scenario(t).
		Given(b.Network()).
		Given(b.Link("1:1:1:2:1")). //place A
		Given(b.Link("1:2:2:1:1")). //line A-B
		Given(b.Link("2:1:2:2:1")). //place B
		Given(b.Signal("A.102.UBL", "1:1:1:2:1", 1, dm.ValidityDirectionWith)).
		Given(b.Signal("B.101.INF", "2:1:2:2:1", 0, dm.ValidityDirectionWith)).
		When(FindingNextPlaceFrom("A.102.UTB")).
		Then(FoundPlaceIs("B")).
		Execute()
}
func TestPreviousPlaceFromArrival(t *testing.T) {
	findPlaceStory.Scenario(t).
		Given(b.Network()).
		Given(b.Link("1:1:1:2:1")). //place A
		Given(b.Link("1:2:2:1:1")). //line A-B
		Given(b.Link("2:1:2:2:1")). //place B
		Given(b.Signal("A.102.UBL", "1:1:1:2:1", 1, dm.ValidityDirectionWith)).
		Given(b.Signal("B.101.INF", "2:1:2:2:1", 0, dm.ValidityDirectionWith)).
		When(FindingPreviousPlaceFrom("B.101.INF")).
		Then(FoundPlaceIs("A")).
		Execute()
}

type NetworkBuilder struct {
	network *dm.Network
}

func (b *NetworkBuilder) Network() *specs.StepResult {
	sr := specs.NewStepResult("network")

	b.network = dm.NewNetwork()

	return sr.Success()
}

func (b *NetworkBuilder) Node(id string) *specs.StepResult {
	res := specs.NewStepResult("network")

	nodeId, err := dm.ParseNodeId(id)

	if err != nil {
		return res.WithError(err)
	}

	b.network.SaveNode(nodeId, "")

	return res.Success()
}

func (b *NetworkBuilder) Link(id string) *specs.StepResult {
	res := specs.NewStepResult("link %s", id)

	linkID, err := dm.ParseLinkId(id)

	if err != nil {
		return res.WithError(err)
	}

	from := b.network.GetOrCreateNode(linkID.From)
	to := b.network.GetOrCreateNode(linkID.To)

	b.network.AddLink(dm.NewLink(linkID, from, to, 1000.0))

	return res.Success()
}

func (b *NetworkBuilder) Signal(sID, lID string, pos float64, dir dm.ValidityDirection) *specs.StepResult {
	res := specs.NewStepResult("signal %q", sID)

	signalID, err := dm.ParseSignalId(sID)
	if err != nil {
		return res.WithError(err)
	}

	linkID, err := dm.ParseLinkId(lID)
	if err != nil {
		return res.WithError(err)
	}

	link := b.network.GetLink(linkID)

	position, err := dm.NewLinkPosition(link, pos, dir)
	if err != nil {
		return res.WithError(err)
	}

	link.Signals.Add(dm.NewSignal(signalID, signalID.Type, position))

	return res.Success()
}
