package domainmodel

// SignalPath represents a signal section between two signals.
type SignalPath struct {
	Start     *Signal
	End       *Signal
	LinkPaths []*LinkPath
}

func NewSignalPath(start, end *Signal, paths []*LinkPath) *SignalPath {
	return &SignalPath{
		Start:     start,
		End:       end,
		LinkPaths: paths,
	}
}
