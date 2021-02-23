package utils

type Point struct {
	X, Y int
}

func (p Point) Equal(pt Point) bool {
	return pt.X == p.X && p.Y == pt.Y
}

type Size struct {
	Width, Height int
}

type Rect struct {
	Left, Top, Right, Bottom int
}

func (r Rect) Size() Size {
	return Size{
		Width:  r.Right - r.Left,
		Height: r.Bottom - r.Top,
	}
}

func (r Rect) PointIn(pt Point) bool {
	if pt.X >= r.Left && pt.X <= r.Right && pt.Y >= r.Top && pt.Y <= r.Bottom {
		return true
	}
	return false
}

func (r Rect) CenterPoint() Point {
	return Point{r.Left + r.Size().Width/2, r.Top + r.Size().Height/2}
}

func (r Rect) CenterRegion() Rect {
	size := r.Size()

	return Rect{
		r.Left + size.Width/3,
		r.Top + size.Height/3,
		r.Left + 2*(size.Width/3),
		r.Top + 2*(size.Height/3),
	}
}

func (r Rect) CenterTopRegion() Rect {
	size := r.Size()

	return Rect{
		r.Left + size.Width/3,
		r.Top,
		r.Left + 2*(size.Width/3),
		r.Top + size.Height/3,
	}
}

func (r Rect) CenterBottomRegion() Rect {
	size := r.Size()

	return Rect{
		r.Left + size.Width/3,
		r.Top + 2*(size.Height/3),
		r.Left + 2*(size.Width/3),
		r.Bottom,
	}
}

func (r Rect) LeftRegion() Rect {
	size := r.Size()
	return Rect{
		r.Left, r.Top, r.Left + size.Width/2, r.Top + size.Height,
	}
}

func (r Rect) TopRegion() Rect {
	size := r.Size()
	return Rect{
		r.Left, r.Top, r.Left + size.Width, r.Top + size.Height/2,
	}
}

func (r Rect) RightRegion() Rect {
	size := r.Size()
	return Rect{
		r.Left + size.Width/2, r.Top, r.Left + size.Width, r.Top + size.Height,
	}
}

func (r Rect) BottomRegion() Rect {
	size := r.Size()
	return Rect{
		r.Left, r.Top + size.Height/2, r.Left + size.Width, r.Top + size.Height,
	}
}

func (r Rect) LeftTopRegion() Rect {
	size := r.Size()
	return Rect{
		r.Left, r.Top, r.Left + size.Width/2, r.Top + size.Height/2,
	}
}
func (r Rect) LeftBottomRegion() Rect {
	size := r.Size()

	return Rect{
		r.Left, r.Top + size.Height/2, r.Left + size.Width/2, r.Top + size.Height,
	}
}
func (r Rect) RightTopRegion() Rect {
	size := r.Size()

	return Rect{
		r.Left + size.Width/2, r.Top, r.Left + size.Width, r.Top + size.Height/2,
	}
}

func (r Rect) RightBottomRegion() Rect {
	size := r.Size()

	return Rect{
		r.Left + size.Width/2, r.Top + size.Height/2, r.Right, r.Bottom,
	}
}
