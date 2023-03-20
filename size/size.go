package size

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Size struct {
	Width  int
	Height int
}

func New(w, h int) *Size {
	return &Size{w, h}
}

func FromString(sStr string) (*Size, error) {
	strItems := strings.Split(strings.Trim(sStr, " ,"), ",")
	if len(strItems) != 2 {
		return nil, errors.New("坐标字符串参数格式错误")
	}
	wStr, hStr := strItems[0], strItems[1]

	w, err := strconv.Atoi(wStr)
	if err != nil {
		return nil, err
	}

	h, err := strconv.Atoi(hStr)
	if err != nil {
		return nil, err
	}

	return &Size{w, h}, nil
}

func (s *Size) Empty() bool {
	return s.Width <= 0 || s.Height <= 0
}

func (s *Size) String() string {
	return fmt.Sprintf("Size(width:%d, height:%d)", s.Width, s.Height)
}

func (s *Size) Equals(s1 *Size) bool {
	return s.Width == s1.Width && s.Height == s1.Height
}
