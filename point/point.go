package point

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/rogeecn/gdm/size"
)

type Point struct {
	X int
	Y int
}

func New(x, y int) *Point {
	return &Point{x, y}
}

func ParseToInt(raw string) ([]int, error) {
	strLst1 := strings.Split(strings.Trim(raw, ", "), ",")
	if len(strLst1) == 0 {
		return nil, errors.New("empty")
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

func FromString(ptStr string) (*Point, error) {
	intLst, err := ParseToInt(ptStr)
	if err != nil {
		return nil, err
	}

	return &Point{intLst[0], intLst[1]}, nil
}

func (p *Point) Equal(pt *Point) bool {
	return p.X == pt.X && p.Y == pt.Y
}

func (p *Point) Offset(x, y int) *Point {
	p.X = p.X + x
	p.Y = p.Y + y

	return p
}

func (p *Point) Clone() *Point {
	return New(p.X, p.Y)
}

func (p *Point) OffsetSize(s *size.Size) *Point {
	p.X = p.X + rand.Intn(s.Width)
	p.Y = p.Y + rand.Intn(s.Height)

	return p
}

func (p *Point) Valid() bool {
	return p.X >= 0 && p.Y >= 0
}

func (p *Point) String() string {
	return fmt.Sprintf("Point(x:%d, y:%d)", p.X, p.Y)
}
