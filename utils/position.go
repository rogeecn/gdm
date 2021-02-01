package utils

type Point struct {
	X, Y int
}

func (p Point) Equal(pt Point) bool {
	return pt.X == p.X && pt.Y == p.Y
}

type Size struct {
	Width, Height int
}

type Rect struct {
	Left, Top, Right, Bottom int
}

func (r Rect) CenterPoint() Point {
	return Point{r.Size().Width / 2, r.Size().Height / 2}
}

func (r Rect) Size() Size {
	return Size{
		Width:  r.Right - r.Left,
		Height: r.Bottom - r.Top,
	}
}

func (r Rect) LeftRegion() Rect {
	size := r.Size()
	return Rect{
		0, 0, size.Width / 2, size.Height,
	}
}
func (r Rect) TopRegion() Rect {
	size := r.Size()
	return Rect{
		0, 0, size.Width, size.Height / 2,
	}
}

func (r Rect) RightRegion() Rect {
	size := r.Size()
	return Rect{
		size.Width / 2, 0, size.Width, size.Height,
	}
}

func (r Rect) BottomRegion() Rect {
	size := r.Size()
	return Rect{
		0, size.Height / 2, size.Width, size.Height,
	}
}

func (r Rect) LeftTopRegion() Rect {
	size := r.Size()
	return Rect{
		0, 0, size.Width / 2, size.Height / 2,
	}
}
func (r Rect) LeftBottomRegion() Rect {
	size := r.Size()

	return Rect{
		0, size.Height / 2, size.Width / 2, size.Height,
	}
}
func (r Rect) RightTopRegion() Rect {
	size := r.Size()

	return Rect{
		size.Width / 2, 0, size.Width, size.Height / 2,
	}
}

func (r Rect) RightBottomRegion() Rect {
	size := r.Size()

	return Rect{
		size.Width / 2, size.Height / 2, size.Width, size.Height,
	}
}
