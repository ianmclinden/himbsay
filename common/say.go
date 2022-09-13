package common

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func splitLines(message string, length int) ([]string, int) {
	var lines []string
	maxLength := 0
	addFormatString := ""
	messageLines := strings.Split(message, "\n")
	for _, line := range messageLines {

		// While there's line left
		for len(line) > 0 {
			subline := line
			formattedLength := len(format(subline, ClearFormatRules))
			lastString := len(subline)

			// Work from the back of the line, until the formatted string is just less than the allowed length
			for formattedLength > length {
				lastString = strings.LastIndex(subline, " ")
				if lastString < 0 {
					// Can't be split...
					lastString = len(subline)
					break
				}
				subline = subline[:lastString]
				formattedLength = len(format(subline, ClearFormatRules))
			}

			lines = append(lines, format(addFormatString+subline, FormatRules))

			lastColor := []int{0, 0}
			lastReset := []int{0, 0}
			if indices := RegexAnsiFormat.FindAllStringIndex(subline, -1); indices != nil {
				lastColor = indices[len(indices)-1]
			}
			if indices := RegexAnsiReset.FindAllStringIndex(subline, -1); indices != nil {
				lastReset = indices[len(indices)-1]
			}
			if lastColor[1] > 0 && lastColor[1] > lastReset[1] {
				addFormatString = subline[lastColor[0]:lastColor[1]]
			} else if lastReset[1] > 0 {
				addFormatString = ""
			}

			if formattedLength > maxLength {
				maxLength = formattedLength
			}
			line = strings.TrimPrefix(line[lastString:], " ")
		}
	}

	return lines, maxLength
}

func boxMessage(message string, maxWidth int, minHeight int) []string {
	lines, maxLen := splitLines(message, maxWidth-4) // len("| x |")==4

	if len(lines) == 0 || maxLen == 0 {
		lines = append(lines, "")
	}

	// Pad line endings
	for i, line := range lines {
		line = format(line, ClearFormatRules)
		if len(line) < maxLen {
			lines[i] += strings.Repeat(" ", maxLen-len(line))
		}
	}

	var box []string

	// Pad to ensure formatting (+ top & bottom line)
	if len(lines)+2 < minHeight {
		for i := len(lines) + 2; i < minHeight; i++ {
			box = append(box, "")
		}
	}

	// build box
	bar := " " + strings.Repeat("-", maxLen+2)

	box = append(box, bar)
	for i, line := range lines {
		if i == 0 && len(lines) > 1 {
			box = append(box, "/ "+line+" \x1b[0m\\")
		} else if i == len(lines)-1 && len(lines) > 1 {
			box = append(box, "\\ "+line+" \x1b[0m/")
		} else if len(lines) > 1 {
			box = append(box, "| "+line+" \x1b[0m|")
		} else {
			box = append(box, "( "+line+" \x1b[0m)")
		}
	}
	box = append(box, bar)
	return box
}

func Say(writer io.Writer, message string, defaultMessage string, template string, padding int, formatLines int, termWidth int, escape bool) error {

	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if info.Mode()&os.ModeCharDevice == 0 || info.Size() > 0 {
		if pipedMessage, err := getPipedInput(os.Stdin); err == nil {
			message = pipedMessage
		}
	}

	if len(message) == 0 {
		message = defaultMessage
	}

	if escape {
		message = format(message, ExtendedFormatRules)
	}
	lines := boxMessage(message, (termWidth - padding), formatLines)

	var extraLines string
	for i := len(lines); i > formatLines; i-- {
		extraLines += strings.Repeat(" ", padding) + "%v\n"
	}

	_, err = fmt.Fprintf(writer, (extraLines + template), toInterface(lines)...)
	return err
}
