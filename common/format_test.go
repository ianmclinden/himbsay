package common

import "testing"

func TestNoFormatRules(t *testing.T) {
	var tests = []struct {
		Input  string
		Output string
	}{
		{"A normal string should not be modified", "A normal string should not be modified"},
		{"\033[1;31mA format string should also not be stripped\033[m", "\033[1;31mA format string should also not be stripped\033[m"},
		{"\t\t\ttabs should not be converted to spaces", "\t\t\ttabs should not be converted to spaces"},
	}
	for _, test := range tests {
		output := format(test.Input, []FormatRule{})
		if test.Output != test.Input {
			t.Errorf("Wanted fomatted string [%s], got [%s]", test.Input, output)
		}
	}
}

func TestClearFormatRules(t *testing.T) {
	var tests = []struct {
		Input  string
		Output string
	}{
		{"A normal string should not be modified", "A normal string should not be modified"},
		{"\033[1;31mA format string needs to be stripped\033[m", "A format string needs to be stripped"},
		{"\x1b[1;31mA format string needs to be stripped\x1b[m", "A format string needs to be stripped"},
		{"\t\t\ttabs should be converted to spaces", "            tabs should be converted to spaces"},
	}
	for _, test := range tests {
		output := format(test.Input, ClearFormatRules)
		if test.Output != output {
			t.Errorf("Wanted fomatted string [%s], got [%s]", test.Output, output)
		}
	}
}

func TestFormatRules(t *testing.T) {
	var tests = []struct {
		Input  string
		Output string
	}{
		{"A normal string should not be modified", "A normal string should not be modified"},
		{"\033[1;31mA format string needs to be converted\033[m", "\x1b[1;31mA format string needs to be converted\x1b[m"},
		{"\x1b[1;31mA properly escaped string should not be modified\x1b[m", "\x1b[1;31mA properly escaped string should not be modified\x1b[m"},
		{"\t\t\ttabs should be converted to spaces", "            tabs should be converted to spaces"},
	}
	for _, test := range tests {
		output := format(test.Input, FormatRules)
		if test.Output != output {
			t.Errorf("Wanted fomatted string [%s], got [%s]", test.Output, output)
		}
	}
}

func TestExtendedFormatRules(t *testing.T) {
	var tests = []struct {
		Input  string
		Output string
	}{
		{"A normal string should not be modified", "A normal string should not be modified"},
		{"\033[1;31mA format string needs to be converted\033[m", "\x1b[1;31mA format string needs to be converted\x1b[m"},
		{"\x1b[1;31mA properly escaped string should not be modified\x1b[m", "\x1b[1;31mA properly escaped string should not be modified\x1b[m"},
		{"\t\t\ttabs should not be converted to spaces", "\t\t\ttabs should not be converted to spaces"},
		{"\\tescaped\\ncharacters should be unescaped", "\tescaped\ncharacters should be unescaped"},
	}
	for _, test := range tests {
		output := format(test.Input, ExtendedFormatRules)
		if test.Output != output {
			t.Errorf("Wanted fomatted string [%s], got [%s]", test.Output, output)
		}
	}
}
