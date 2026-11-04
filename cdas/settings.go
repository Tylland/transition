package cdas

const (
	TimestampFormat = "2006-01-02T15:04:05.000Z07:00"
	YearFormat      = "2006"
	DateFormat      = "2006-01-02"
)

const (
	TrackRoutesBucket    = "TRACK_ROUTES"
	SegmentProfileBucket = "SEGMENT_PROFILES"
	JourneyProfileBucket = "JOURNEY_PROFILES"
)

const (
	SFERASubjectB2GRequest = "sfera.b2g.request"
	SFERASubjectB2GEvent   = "sfera.b2g.event"
	SFERASubjectB2GReply   = "sfera.b2g.reply"
	SFERASubjectG2BRequest = "sfera.g2b.request"
	SFERASubjectG2BEvent   = "sfera.g2b.event"
	SFERASubjectG2BReply   = "sfera.g2b.reply"
	SFERASubjectInvalid    = "sfera.invalid"
)

const (
	B2GStreamName = "B2G_MESSAGES"
	G2BStreamName = "G2B_MESSAGES"
)

const (
	B2GDefaultConsumerName = "B2G_MESSAGES_READER"
)

const (
	B2GSubjectFilter = "sfera.b2g.*"
	G2BSubjectFilter = "sfera.g2b.*"
)

var B2GSubjects []string = []string{SFERASubjectB2GRequest, SFERASubjectB2GEvent, SFERASubjectB2GReply}
var G2BSubjects []string = []string{SFERASubjectG2BRequest, SFERASubjectG2BEvent, SFERASubjectG2BReply}
