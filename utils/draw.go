package utils

import (
	"strconv"
	"strings"

	"github.com/rogeecn/gdm/point"
	"github.com/rogeecn/gdm/rect"
	"github.com/rogeecn/gdm/size"
)

func ParseToInt(raw string) ([]int, error) {
	strLst1 := strings.Split(strings.Trim(raw, ", "), ",")
	if len(strLst1) == 0 {
		return nil, ErrorEmptyData
	}
	var strLst []string
	for _, s := range strLst1 {
		ss := strings.Split(strings.Trim(s, ", "), "|")
		strLst = append(strLst, ss...)
	}

	intLst := make([]int, len(strLst))
	for i, s := range strLst {
		d, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		intLst[i] = d
	}

	return intLst, nil
}

func ParsePointSize(x, y, w, h int) (*point.Point, *size.Size) {
	return point.New(x, y), size.New(w, h)
}

func ParsePoint(x, y, x1, y1 int) (*point.Point, *point.Point) {
	return point.New(x, y), point.New(x1, y1)
}

func PointToRect(x, y, x1, y1 int) *rect.Rect {
	return rect.FromPoint(ParsePoint(x, y, x1, y1))
}

func SizeToRect(x, y, w, h int) *rect.Rect {
	return rect.FromSize(ParsePointSize(x, y, w, h))
}
