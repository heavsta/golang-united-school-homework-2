package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type side int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesCircle side = 0
	SidesTriangle side = 3
	SidesSquare side = 4
)

func CalcSquare(sideLen float64, sidesNum side) float64 {
	switch sidesNum {
	case SidesCircle:
		// area = radius² * π
		return math.Pow(sideLen, 2) * math.Pi
	case SidesTriangle:
		// area = side² * √3 / 4
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case SidesSquare:
		// area = side²
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
