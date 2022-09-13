package common

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGetRandomFrom(t *testing.T) {
	var (
		stringCollection = []string{"able", "baker", "charlie", "dog"}
		intCollection    = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	)

	for i := 0; i < len(stringCollection)*2; i++ {
		stringSample := GetRandomFrom(stringCollection)
		if !slices.Contains(stringCollection, stringSample) {
			t.Errorf("Got element that was not in original collection")
		}
	}
	for i := 0; i < len(intCollection)*2; i++ {
		intSample := GetRandomFrom(intCollection)
		if !slices.Contains(intCollection, intSample) {
			t.Errorf("Got element that was not in original collection")
		}
	}
}

func TestGetPipedInput(t *testing.T) {
	var tests = []struct {
		Input string
	}{
		{"foobar"},
		{"test\ntwo\n\three"},
	}
	for _, test := range tests {

		var reader = strings.NewReader(test.Input)
		input, err := getPipedInput(reader)
		if err != nil {
			t.Error(err)
		}
		if input != test.Input {
			t.Errorf("Expected input [%s], got [%s]", test.Input, input)
		}
	}
}

func TestGetTermSize(t *testing.T) {
	expectedSize := defaultTermSize
	size := GetTermSize()
	if size != expectedSize {
		t.Errorf("Expected default terminal size %d, got %d", expectedSize, size)
	}
}
