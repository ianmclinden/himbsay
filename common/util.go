package common

import (
	"bufio"
	"io"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/term"
)

const defaultTermSize int = 80 // If stdin and stdout are both being used

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

func GetTermSize() int {
	termSize, _, _ := term.GetSize(1)
	if termSize <= 0 {
		termSize = defaultTermSize
	}
	return termSize
}

func toInterface[T any](collection []T) []interface{} {
	var output []interface{}
	for _, item := range collection {
		output = append(output, item)
	}
	return output
}
