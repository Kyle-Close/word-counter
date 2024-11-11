package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	data := "This! is a tesT\nthis, Is a\nthis, is\nThis?"
	result := CountWords(data)

	expect := map[string]int{
		"this": 4,
		"is":   3,
		"a":    2,
		"test": 1,
	}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Result does not match.\n got: %v\nwant: %v", result, expect)
	}
}
