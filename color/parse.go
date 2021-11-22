package color

import (
	"strings"
)

func Parse(c string) []*Color {
	colors := strings.Split(strings.TrimSpace(c), "|")

	colorLst := make([]*Color, len(colors))
	for i, c := range colors {
		cs := strings.Split(strings.TrimSpace(c), "-")

		if len(cs) == 1 {
			cs = append(cs, "010101")
		}
		colorLst[i] = NewColor(cs[0]).WithDelta(cs[1])
	}

	return colorLst
}
