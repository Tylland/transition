package convertion

import "math"

// TrackPoint is the interface for points used in the Douglas-Peucker algorithm.
type TrackPoint interface {
	GetX() float64
	GetY() float64
}

// DouglasPeuckerTransform applies the Douglas-Peucker line simplification algorithm
// to reduce the number of points while preserving the shape.
func DouglasPeuckerTransform[T TrackPoint](positions []T) []T {
	return douglasPeucker(positions, 1.0)
}

func calcPerpendicularDistance[T TrackPoint](location, start, end T) float64 {
	a := location.GetY() - start.GetY()
	b := location.GetX() - start.GetX()
	c := end.GetY() - start.GetY()
	d := end.GetX() - start.GetX()

	return math.Abs(a*d-c*b) / math.Sqrt(c*c+d*d)
}

func douglasPeucker[T TrackPoint](positions []T, epsilon float64) []T {
	if len(positions) < 3 {
		return positions
	}

	maxDistance := 0.0
	maxAltitudeIndex := 0

	for i := 1; i < len(positions)-2; i++ {
		distance := calcPerpendicularDistance(positions[i], positions[0], positions[len(positions)-1])
		if distance > maxDistance {
			maxDistance = distance
			maxAltitudeIndex = i
		}
	}

	var result []T

	if maxDistance > epsilon {
		result1 := douglasPeucker(positions[:maxAltitudeIndex+1], epsilon)
		result2 := douglasPeucker(positions[maxAltitudeIndex:], epsilon)

		// Add all of result1 except the last point (to avoid duplication)
		result = append(result, result1[:len(result1)-1]...)
		result = append(result, result2...)
	} else {
		result = append(result, positions[0], positions[len(positions)-1])
	}

	return result
}
