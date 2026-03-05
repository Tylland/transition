package convertion

import (
	dm "github.com/tylland/transition/domainmodel"
	rml "github.com/tylland/transition/schema/railml"
)

// RailML extension functions for converting RailML schema types to domain model types.

// defaultName is used when no names are available.
var defaultName = rml.Name{NameProperty: "", Description: "", Language: "se"}

// RMLSpotLocationGetRiktning returns the Gilltighetsriktning for a spot location.
func RMLSpotLocationGetRiktning(location rml.RTMSpotLocation) dm.Gilltighetsriktning {
	return RMLApplicationDirectionToGilltighetsriktning(location.ApplicationDirection)
}

// RMLApplicationDirectionToGilltighetsriktning converts a RailML TApplicationDirection
// to a domain Gilltighetsriktning.
func RMLApplicationDirectionToGilltighetsriktning(direction rml.TApplicationDirection) dm.Gilltighetsriktning {
	switch direction {
	case rml.TApplicationDirectionNormal:
		return dm.GilltighetsriktningNormal
	case rml.TApplicationDirectionReverse:
		return dm.GilltighetsriktningOmvand
	case rml.TApplicationDirectionBoth:
		return dm.GilltighetsriktningBada
	default:
		return dm.GilltighetsriktningIngen
	}
}

// RMLGetFirstName returns the first name from a slice of Name, or "" if empty.
func RMLGetFirstName(names []rml.Name) string {
	if len(names) > 0 {
		return names[0].NameProperty
	}
	return defaultName.NameProperty
}

// RMLGetFirstDescription returns the first description from a slice of Name, or "" if empty.
func RMLGetFirstDescription(names []rml.Name) string {
	if len(names) > 0 {
		return names[0].Description
	}
	return defaultName.Description
}
