package common

import (
	"bufio"
	"io"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/term"
)

const (
	defaultTermWidth  int = 80 // Ifstdout is being used
	defaultTermHeight int = 80 // If stdout is being used
)

func GetRandomFrom[T any](collection []T) T {
	rand.Seed(time.Now().UnixNano())
	return collection[rand.Intn(len(collection))]
}

func getPipedInput(reader io.Reader) (string, error) {
	var message []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		message = append(message, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(message, "\n"), nil
}

func GetTermSize() (int, int) {
	width, height, err := term.GetSize(1)
	if err != nil {
		return defaultTermWidth, defaultTermHeight
	}
	if width <= 0 {
		width = defaultTermWidth
	}
	if height <= 0 {
		height = defaultTermHeight
	}
	return width, height
}

func toInterface[T any](collection []T) []interface{} {
	var output []interface{}
	for _, item := range collection {
		output = append(output, item)
	}
	return output
}
