package utils

import (
	"fmt"
	"testing"
)

func TestNewFindItemsResult(t *testing.T) {
	words := []string{"a", "b"}
	str := "0,10,20|1,10,20"
	str = ""
	r := NewFindItemsResult(words, str)
	for _, item := range r.Items {
		fmt.Println(item.Index, item.Item, item.Point)
	}
}
