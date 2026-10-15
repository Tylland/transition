package port

import (
	"github.com/tylland/transition/internal/core/domain"
	"github.com/tylland/transition/internal/core/domain/network"
	sfera "github.com/tylland/transition/internal/schema/sfera/v3"
)

type SignalRouteWriter interface {
	Write(sr *domain.SignalRoutes) error
}

type TrackRouteWriter interface {
	Write(tr *domain.TrackRoutes) error
}

type SegmentProfileWriter interface {
	Write(sp *sfera.SegmentProfile) (string, error)
}

// type NetworkFunc func(*network.Network) error

// type NetworkProvider interface {
// 	Start(ctx context.Context, callback NetworkFunc)
// }

type FacilityService interface {
	ProcessNetwork(net *network.Network) error
}
