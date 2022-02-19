package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type intCustomType int

func calcEquilateralFigurineSquare(sideLen float64, sidesNum intCustomType) float64 {
	alpha := 2 * math.Pi / float64(sidesNum)
	return math.Pow(sideLen, 2) / (4 * math.Tan(alpha/2))
}

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case 0:
		return math.Pow(sideLen, 2) * math.Pi
	case 3:
		fallthrough
	case 4:
		return calcEquilateralFigurineSquare(sideLen, sidesNum)
	default:
		return 0
	}
}
