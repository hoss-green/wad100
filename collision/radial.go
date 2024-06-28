package collision

import "math"

func RadialCollision(point Point, x float64, y float64, inner float64, width float64) bool {
	distance := math.Sqrt(math.Pow(point.X-x, 2) + math.Pow(point.Y-y, 2))
  return !(distance < inner || distance > inner + width)
}
