package common

import (
	"regexp"
	"strconv"
)

type FormatRule func(string) string

var (
	RegexAnsiFormat = regexp.MustCompile(`\x1b\[[0-9;]+m`)
	RegexAnsiReset  = regexp.MustCompile(`\x1b\[[0]?m`)
)

var (
	ClearFormatRules = []FormatRule{
		func(m string) string { return regexp.MustCompile(`\x1b\[[0-9;]*m`).ReplaceAllString(m, "") }, // ANSI excape chars
		func(m string) string { return regexp.MustCompile(`\t`).ReplaceAllString(m, "    ") },         // Tabs to spaces
	}
	FormatRules = []FormatRule{
		func(m string) string { return regexp.MustCompile(`\t`).ReplaceAllString(m, "    ") }, // Tabs to spaces
	}
	ExtendedFormatRules = []FormatRule{
		func(m string) string { return regexp.MustCompile(`\\033`).ReplaceAllString(m, "\x1b") }, // ANSI excape chars
		func(m string) string { m, _ = strconv.Unquote("\"" + m + "\""); return m },              // Any other escape characters
	}
)

func format(message string, rules []FormatRule) string {
	for _, rule := range rules {
		if rule != nil {
			message = rule(message)
		}
	}
	return message
}
