package utils

import (
	"strconv"
	"strings"
)

type FindItemsResult struct {
	Items []*FindItemResult
}

func NewFindItemsResult(words []string, ret string) *FindItemsResult {
	rets := strings.Split(ret, "|")

	f := &FindItemsResult{}
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
	Point Point
}

func NewFindItemResult(words []string, ret string) *FindItemResult {
	rets := strings.Split(ret, "|")

	f := &FindItemResult{}
	f.Index, _ = strconv.Atoi(rets[0])
	if f.Index == -1 {
		return f
	}

	f.Item = words[f.Index]
	f.Point.X, _ = strconv.Atoi(rets[1])
	f.Point.Y, _ = strconv.Atoi(rets[2])

	return f
}

func (f *FindItemResult) IsOK() bool {
	return IsFindStrOK(int64(f.Index))
}

type OcrResult struct {
	Char  string
	Point Point
}

func NewOcrResult(ret string) *OcrResult {
	rets := strings.Split(ret, "$")

	f := &OcrResult{}
	f.Char = rets[0]
	f.Point.X, _ = strconv.Atoi(rets[1])
	f.Point.Y, _ = strconv.Atoi(rets[2])

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
