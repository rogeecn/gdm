package color

import (
	"strings"
)

func ParseColor(c string) []*Color {
	colors := strings.Split(strings.TrimSpace(c), "|")

	colorLst := make([]*Color, len(colors))
	for i, c := range colors {
		cs := strings.Split(strings.TrimSpace(c), "-")

		colorLst[i] = NewColor(cs[0]).WithDelta(cs[1])
	}

	return colorLst
}
