package checkpoint

func RectPerimeter(h int, w int) int {
	if h <= 0 || w <= 0 {
		return -1
	}
	perimeter := h*2 + w*2
	return perimeter
}
