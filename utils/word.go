package utils

import (
	"fmt"
	"github.com/rogeecn/draw"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type FindItemsResult struct {
	Items []*FindItemResult
}

func NewFindItemsResult(words []string, ret string) *FindItemsResult {
	f := &FindItemsResult{}
	if len(ret) == 0 {
		return f
	}

	rets := strings.Split(ret, "|")
	for _, item := range rets {
		item = strings.Replace(item, ",", "|", -1)
		f.Items = append(f.Items, NewFindItemResult(words, item))
	}
	return f
}

func (f *FindItemsResult) IsOK() bool {
	return f.Count() > 0
}

func (f *FindItemsResult) Count() int {
	return len(f.Items)
}

type FindItemResult struct {
	Index int
	Item  string
	Point *draw.Point
}

func NewFindItemResult(words []string, ret string) *FindItemResult {
	log.Debugf("NewFindItemResult Parse Result words(%+v), ret: %s", words, ret)

	f := &FindItemResult{}
	if len(ret) == 0 {
		return f
	}

	rets := strings.Split(ret, "|")
	f.Index, _ = strconv.Atoi(rets[0])
	if f.Index == -1 {
		log.Debugf("NewFindItemResult got index: %d", f.Index)

		return f
	}
	if f.Index > len(words) {
		f.Index = -1
		return f
	}
	log.Debugf("NewFindItemResult got index: %d", f.Index)

	f.Item = words[f.Index]
	f.Point = draw.NewPoint(-1, -1)

	f.Point.X, _ = strconv.Atoi(rets[1])
	f.Point.Y, _ = strconv.Atoi(rets[2])

	return f
}

func (f *FindItemResult) IsOK() bool {
	return IsFindStrOK(int64(f.Index))
}

type OcrResult struct {
	Char  string
	Point *draw.Point
}

func NewOcrResult(ret string) *OcrResult {
	rets := strings.Split(ret, "$")

	f := &OcrResult{}
	f.Char = rets[0]

	f.Point, _ = draw.NewPointFromString(fmt.Sprintf("%s,%s", rets[1], rets[2]))

	return f
}

type OcrResuts struct {
	results []*OcrResult
}

func NewOcrResuts(ret string) *OcrResuts {
	rets := strings.Split(ret, "|")

	results := &OcrResuts{}
	for _, item := range rets {
		results.results = append(results.results, NewOcrResult(item))
	}
	return results
}
