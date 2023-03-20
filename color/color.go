package color

import (
	"fmt"
	"strings"

	"github.com/rogeecn/draw"
)

type Color struct {
	color string
	delta string
	// nolint // placeholder for future use
	gray string
}

func NewColor(color string) *Color {
	return &Color{color: color, delta: "000000"}
}

func (c *Color) WithDelta(d string) *Color {
	c.delta = d
	return c
}

func (c *Color) String() string {
	return fmt.Sprintf("%s-%s", c.color, c.delta)
}

type Colors struct {
	colors []*Color
}

func NewColors() *Colors {
	return &Colors{}
}

func (c *Colors) Add(col ...*Color) *Colors {
	c.colors = append(c.colors, col...)
	return c
}

func (c *Colors) String() string {
	var colorStrs []string
	for _, color := range c.colors {
		colorStrs = append(colorStrs, color.String())
	}

	return strings.Join(colorStrs, "|")
}

type ColorPosition struct {
	color string
	Point *draw.Point
}

func NewColorPosition(color string, ptX, ptY int) *ColorPosition {
	return &ColorPosition{
		color: color,
		Point: draw.NewPoint(ptX, ptY),
	}
}
