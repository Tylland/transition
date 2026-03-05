package convertion

import (
	"fmt"

	rml "github.com/tylland/transition/schema/railml"
)

// RailMLTopologyReader provides helper methods for navigating RailML topology.
type RailMLTopologyReader struct {
	infra     rml.Infrastructure
	elements  map[string]rml.NetElement
	relations map[string]rml.NetRelation
}

// NewRailMLTopologyReader creates a new RailMLTopologyReader.
func NewRailMLTopologyReader(infra rml.Infrastructure) *RailMLTopologyReader {
	elements := make(map[string]rml.NetElement, len(infra.Topology.NetElements))
	for _, ne := range infra.Topology.NetElements {
		elements[ne.Id] = ne
	}

	relations := make(map[string]rml.NetRelation, len(infra.Topology.NetRelations))
	for _, nr := range infra.Topology.NetRelations {
		relations[nr.Id] = nr
	}

	return &RailMLTopologyReader{
		infra:     infra,
		elements:  elements,
		relations: relations,
	}
}

// GetElement returns a net element by ID.
func (r *RailMLTopologyReader) GetElement(id string) (rml.NetElement, error) {
	el, ok := r.elements[id]
	if !ok {
		return rml.NetElement{}, fmt.Errorf("net element not found: %s", id)
	}
	return el, nil
}

// GetRelations returns the net relations referenced by the given element.
func (r *RailMLTopologyReader) GetRelations(element rml.NetElement) []rml.NetRelation {
	var result []rml.NetRelation
	for _, ref := range element.Relation {
		if rel, ok := r.relations[ref.Ref]; ok {
			result = append(result, rel)
		}
	}
	return result
}

// GetRelatedReference returns the ID of the element on the other side of a relation.
func (r *RailMLTopologyReader) GetRelatedReference(element rml.NetElement, relation rml.NetRelation) (string, error) {
	if element.Id == relation.ElementA.Ref {
		return relation.ElementB.Ref, nil
	}
	if element.Id == relation.ElementB.Ref {
		return relation.ElementA.Ref, nil
	}
	return "", fmt.Errorf("relationen innehåller inte elementet")
}

// GetRelatedElement returns the element on the other side of a relation.
func (r *RailMLTopologyReader) GetRelatedElement(element rml.NetElement, relation rml.NetRelation) (rml.NetElement, error) {
	ref, err := r.GetRelatedReference(element, relation)
	if err != nil {
		return rml.NetElement{}, err
	}
	return r.GetElement(ref)
}
