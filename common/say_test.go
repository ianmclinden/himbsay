package common

import (
	"testing"
)

func TestSplitLines(t *testing.T) {
	var tests = []struct {
		Input    string
		MaxWidth int
		Output   []string
		MaxLine  int
	}{
		{"no split", 40, []string{"no split"}, 8},
		{"this line should be split", 14, []string{"this line", "should be", "split"}, 9},
		{"this\nshould\nbe\nsplit\ntoo", 40, []string{"this", "should", "be", "split", "too"}, 6},
		{"extra  spaces  should  be  preserved", 40, []string{"extra  spaces  should  be  preserved"}, 36},
		{"\x1b[1;31mUnclosed format string\nneeds to be preserved", 40, []string{"\x1b[1;31mUnclosed format string", "\x1b[1;31mneeds to be preserved"}, 22},
	}
	for _, test := range tests {
		output, maxLineSize := splitLines(test.Input, test.MaxWidth)
		if len(test.Output) != len(output) {
			t.Errorf("Expected %d lines, got %d", len(test.Output), len(output))
		}
		if test.MaxLine != maxLineSize {
			t.Errorf("Expected maximum split line of size %d, got %d", test.MaxLine, maxLineSize)
		}
		for i, line := range output {
			if test.Output[i] != line {
				t.Errorf("Wanted line [%s], got [%s]", test.Output[i], line)
			}
		}
	}
}

func TestBoxMessage(t *testing.T) {
	var tests = []struct {
		Input     string
		MaxWidth  int
		MaxHeight int
		Output    []string
	}{
		{"", 40, 1, []string{
			" --",
			"(  \x1b[0m)",
			" --",
		}},
		{"a nice message", 40, 1, []string{
			" ----------------",
			"( a nice message \x1b[0m)",
			" ----------------",
		}},
		{"a nice message", 14, 1, []string{
			" ----------------",
			"( a nice message \x1b[0m)",
			" ----------------",
		}},
		{"\033[0;31ma\033[m \033[0;32mnice\033[m \033[0;34mmessage\033[m", 40, 1, []string{
			" ----------------",
			"( \033[0;31ma\033[m \033[0;32mnice\033[m \033[0;34mmessage\033[m \x1b[0m)",
			" ----------------",
		}},
		{"a nice truncated message", 21, 2, []string{
			" ------------------",
			"/ a nice truncated \x1b[0m\\",
			"\\ message          \x1b[0m/",
			" ------------------",
		}},
		{"a nice long message that will surely be wrapped", 21, 3, []string{
			" ---------------------",
			"/ a nice long message \x1b[0m\\",
			"| that will surely be \x1b[0m|",
			"\\ wrapped             \x1b[0m/",
			" ---------------------",
		}},
		{"a nice long message that will surely be truncated", 21, 2, []string{
			" ---------------------",
			"/ that will surely be \x1b[0m\\",
			"\\ truncated           \x1b[0m/",
			" ---------------------",
		}},
	}
	for _, test := range tests {
		output := boxMessage(test.Input, test.MaxWidth, test.MaxHeight)
		for i, line := range output {
			if test.Output[i] != line {
				t.Errorf("Expected line [%s], got [%s]", test.Output[i], line)
			}
		}
	}
}
