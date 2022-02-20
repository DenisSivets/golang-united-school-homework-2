package square

import "math"

type sides int

const (
	SidesTriangle sides = 3
	SidesSquare    sides = 4
	SidesCircle    sides = 0
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	var square float64
	if sidesNum == SidesTriangle {
		square = math.Sqrt(3) * sideLen * sideLen / 4
	}
	if sidesNum == SidesSquare {
		square = sideLen * sideLen
	}
	if sidesNum == SidesCircle {
		square = math.Pi * sideLen * sideLen
	}
	return square
}
