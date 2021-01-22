package utils

import (
	"fmt"
	"strings"
)

type Color struct {
	color string
	delta string
	gray  string
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
	for _, color := range col {
		c.colors = append(c.colors, color)
	}
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
	Point Point
}

func NewColorPosition(color string, ptX, ptY int) *ColorPosition {
	return &ColorPosition{
		color: color,
		Point: Point{
			X: ptX,
			Y: ptY,
		},
	}
}
