package utils

import (
	"strconv"
	"strings"
)

func ToPoint(pts string) Point {
	pt := strings.Split(pts, ",")
	ptX, _ := strconv.Atoi(pt[0])
	ptY, _ := strconv.Atoi(pt[1])
	return Point{ptX, ptY}
}
