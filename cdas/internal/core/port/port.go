package port

import (
	"context"

	"github.com/tylland/transition/cdas/internal/core/domain"
	"github.com/tylland/transition/cdas/internal/core/domain/network"
	sfera "github.com/tylland/transition/cdas/internal/schema/sfera/v3"
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

type FacilityService interface {
	ProcessNetwork(net *network.Network) error
}

type SFERAMessagePublisher interface {
	Publish(msg any) error
}

type SegmentProfileReader interface {
	Read(ctx context.Context, key string) (*sfera.SegmentProfile, error)
}

type KeyProvider[T any] interface {
	Create(id string) string
}

type SPKeyProvider interface {
	Info(id string) string
	Data(id string) string
}
