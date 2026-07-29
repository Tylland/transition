package domainmodel

// Tagkategori contains standard train category constants.
type TrainCateory string

const (
	TrainCateoryA TrainCateory = "A"
	TrainCateoryB TrainCateory = "B"
	TrainCateoryC TrainCateory = "C"
	TrainCateoryS TrainCateory = "S"
)

// SpeedProfile groups speed sections by category.
type SpeedProfile struct {
	Category TrainCateory
	Segments Entities[*SpeedSegment]
}

func NewSpeedProfile(category TrainCateory) *SpeedProfile {
	return &SpeedProfile{
		Category: category,
	}
}

// SpeedSegment represents a speed section with a spread across a link.
type SpeedSegment struct {
	*EntityWithExtent
	Category TrainCateory
	Speed    int
}

func NewSpeedSegment(category TrainCateory, speed int, extent *LinkExtent) *SpeedSegment {
	return &SpeedSegment{
		EntityWithExtent: NewEntityWithExtent(extent),
		Category:         category,
		Speed:            speed,
	}
}
