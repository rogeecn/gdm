package rect

import (
	"fmt"

	"github.com/rogeecn/gdm/point"
	"github.com/rogeecn/gdm/size"
)

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

func New(s *size.Size) *Rect {
	return &Rect{0, 0, s.Width, s.Height}
}

func FromSize(lt *point.Point, s *size.Size) *Rect {
	return &Rect{lt.X, lt.Y, s.Width, s.Height}
}

func FromPoint(lt *point.Point, rb *point.Point) *Rect {
	return &Rect{lt.X, lt.Y, rb.X - lt.X, rb.Y - lt.Y}
}
func (r *Rect) Clone() *Rect {
	return &Rect{r.X, r.Y, r.Width, r.Height}
}

func (r *Rect) Left() int {
	return r.X
}

func (r *Rect) Right() int {
	return r.X + r.Width
}

func (r *Rect) Top() int {
	return r.Y
}

func (r *Rect) Bottom() int {
	return r.Y + r.Height
}

func (r *Rect) TopLeft() *point.Point {
	return point.New(r.X, r.Y)
}

func (r *Rect) TopRight() *point.Point {
	return point.New(r.X+r.Width, r.Y)
}

func (r *Rect) BottomLeft() *point.Point {
	return point.New(r.X, r.Y+r.Height)
}

func (r *Rect) BottomRight() *point.Point {
	return point.New(r.X+r.Width, r.Y+r.Height)
}
func (r *Rect) Center() *point.Point {
	return point.New(r.X+r.Width/2, r.Y+r.Height/2)
}

func (r *Rect) Empty() bool {
	return r.Width <= 0 || r.Height <= 0
}

func (r *Rect) Location() *point.Point {
	return point.New(r.X, r.Y)
}

func (r *Rect) SetLocation(pt *point.Point) {
	r.X = pt.X
	r.Y = pt.Y
}

func (r *Rect) Offset(ox, oy int) *Rect {
	r.X = r.X + ox
	r.Y = r.Y + oy

	return r
}

func (r *Rect) Contains(pt *point.Point) bool {
	return r.X < pt.X && r.Y < pt.Y && pt.X < r.Right() && pt.Y < r.Bottom()
}

func (r *Rect) RegionLeft() *Rect {
	return FromSize(r.TopLeft(), size.New(r.Width/2, r.Height))
}

func (r *Rect) RegionRight() *Rect {
	rect := r.RegionLeft()
	rect.SetLocation(point.New(r.Left()+r.Width/2, r.Top()))
	return rect
}

func (r *Rect) RegionTop() *Rect {
	return FromSize(r.TopLeft(), size.New(r.Width, r.Height/2))
}

func (r *Rect) RegionBottom() *Rect {
	rect := r.RegionTop()
	rect.SetLocation(point.New(r.Left(), r.Top()+r.Width/2))
	return rect
}

func (r *Rect) RegionTopLeft() *Rect {
	return FromSize(r.TopLeft(), size.New(r.Width/2, r.Height/2))
}

func (r *Rect) RegionTopCenter() *Rect {
	return FromSize(point.New(r.Width/3, 0), size.New(r.Width/3, r.Height/2))
}

func (r *Rect) RegionTopRight() *Rect {
	rt := r.RegionTopLeft()
	rt.SetLocation(point.New(r.X+r.Width/2, r.Y))
	return rt
}

func (r *Rect) RegionBottomLeft() *Rect {
	rt := r.RegionTopLeft()
	rt.SetLocation(point.New(r.X, r.Y+r.Height/2))
	return rt
}

func (r *Rect) RegionBottomCenter() *Rect {
	return FromSize(point.New(r.Width/3, r.Height/2), size.New(r.Width/3, r.Height/2))
}

func (r *Rect) RegionBottomRight() *Rect {
	rt := r.RegionTopLeft()
	rt.SetLocation(point.New(r.X+r.Width/2, r.Y+r.Height/2))
	return rt
}

func (r *Rect) RegionCenter() *Rect {
	rt := (size.New(r.Width/2, r.Height))
	rt.SetLocation(point.New(r.X+r.Width/3, r.Y))
	return rt
}

func (r *Rect) RegionMiddle() *Rect {
	rt := (size.New(r.Width, r.Height/2))
	rt.SetLocation(point.New(r.X, r.Y+r.Height/3))
	return rt
}

func (r *Rect) RegionMiddleCenter() *Rect {
	rt := r.RegionTopLeft()
	rt.SetLocation(point.New(r.X+r.Width/3, r.Y+r.Height/3))
	return rt
}

func (r *Rect) String() string {
	return fmt.Sprintf("Rect(x:%d, y:%d, width:%d, height:%d)", r.X, r.Y, r.Width, r.Height)
}
