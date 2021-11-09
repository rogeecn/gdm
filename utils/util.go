package utils

import (
	"github.com/rogeecn/draw"
	"strconv"
	"strings"
)

func ToPoint(pts string) *draw.Point {
	pt := strings.Split(pts, ",")
	ptX, _ := strconv.Atoi(pt[0])
	ptY, _ := strconv.Atoi(pt[1])
	return draw.NewPoint(ptX, ptY)
}
