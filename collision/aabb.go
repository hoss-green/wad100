package collision

func AabbCollision(point Point, x float64, y float64, width float64, height float64) bool {
	left := x - width/2
	right := x + width/2
	top := y - height/2
	bottom := y + height/2

	return !(point.X < left || point.X > right || point.Y > bottom || point.Y < top)
}
