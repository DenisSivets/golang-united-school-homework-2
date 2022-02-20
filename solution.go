package square

import "math"

type sides int

const (
	SidesTriangle sides=3
	SidesSquar  sides=4
	SidesCircl  sides=0
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	var square float64
	if sidesNum==SidesTriangle{
		square=math.Sqrt(3) * sideLen * sideLen / 4
	}
	if sidesNum==SidesSquar{
		square=sideLen*sideLen
	}
	if sidesNum==SidesCircl{
		square=math.Pi * sideLen * sideLen
	}
	return square
}
