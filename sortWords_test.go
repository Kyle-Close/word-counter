package main

import (
	"reflect"
	"testing"
)

func TestSortWordsBase(t *testing.T) {
	data := map[string]int{
		"testA": 3,
		"testB": 8,
		"testC": 2,
		"testD": 5,
	}

	result := SortWords(data)
	expect := []WordCount{
		{word: "testB", count: 8},
		{word: "testD", count: 5},
		{word: "testA", count: 3},
		{word: "testC", count: 2},
	}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Result incorrect. got: %v, want: %v", result, expect)
	}
}

func TestSortWordsWithEqualCount(t *testing.T) {
	data := map[string]int{
		"testA":  3,
		"testB":  8,
		"testC":  2,
		"testD":  5,
		"testAE": 5,
		"testBF": 8,
	}
	result := SortWords(data)
	expect := []WordCount{
		{word: "testB", count: 8},
		{word: "testBF", count: 8},
		{word: "testAE", count: 5},
		{word: "testD", count: 5},
		{word: "testA", count: 3},
		{word: "testC", count: 2},
	}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Result incorrect.\ngot: %v\n want: %v\n", result, expect)
	}
}
