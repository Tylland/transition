// Package convertion provides the application logic for the CDAS Segment Factory.
// It contains network sources, SFERA segment profile generation, and file I/O.
package convertion

import dm "github.com/tylland/transition/internal/domainmodel"

// NatSupplier is the interface for creating a network (Nat) from a data source.
type NatSupplier interface {
	CreateNat() *dm.Nat
}
