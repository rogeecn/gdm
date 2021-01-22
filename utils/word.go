package utils

import (
	"strconv"
	"strings"
)

type FindWordsResult struct {
	Words []*FindWordResult
}

func NewFindWordsResult(words []string, ret string) *FindWordsResult {
	rets := strings.Split(ret, "|")

	f := &FindWordsResult{}
	for _, item := range rets {
		item = strings.Replace(item, ",", "|", -1)
		f.Words = append(f.Words, NewFindWordResult(words, item))
	}
	return f
}

func (f *FindWordsResult) IsOK() bool {
	return f.Count() > 0
}

func (f *FindWordsResult) Count() int {
	return len(f.Words)
}

type FindWordResult struct {
	Index int
	Word  string
	Point Point
}

func NewFindWordResult(words []string, ret string) *FindWordResult {
	rets := strings.Split(ret, "|")

	f := &FindWordResult{}
	f.Index, _ = strconv.Atoi(rets[0])
	if f.Index == -1 {
		return f
	}

	f.Word = words[f.Index]
	f.Point.X, _ = strconv.Atoi(rets[1])
	f.Point.Y, _ = strconv.Atoi(rets[2])

	return f
}

func (f *FindWordResult) IsOK() bool {
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
