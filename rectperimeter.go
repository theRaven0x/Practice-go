package piscine

func RectPerimeter(w, h int) int {

	if w < 0 || h < 0 {
		return -1
	}
	perimeter := 2*h + 2*w
	return perimeter
}
