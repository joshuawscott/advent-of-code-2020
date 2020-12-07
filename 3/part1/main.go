package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Slope defines the map of the area that we are toboggoning down
type Slope struct {
	Width  int
	Height int
	// god, I hope tree or not is the only thing we care about here...
	visibleArea []string
}

// NewSlope takes a input []byte representing the slope in ASCII, with # = tree and . = empty.
// Returns a new Slope object that can find if a tree exists at a coordinate. Coordinates are
// measured screen-style, with 0,0 top left, and 3, 1 being three right, 1 down.
func NewSlope(data []byte) Slope {
	strValues := strings.Split(string(data), "\n")
	height := len(strValues)
	width := len(strValues[0])
	return Slope{
		Width:       width,
		Height:      height,
		visibleArea: strValues,
	}
}

// At returns a string saying what is at coordinates x, y.
// Error codes:
// EOM: Beyond the End of the Map.
// INVALID: invalid thing at requested point.
func (s Slope) At(x, y int) (string, error) {
	if y >= s.Height {
		return "", errors.New("EOM")
	}
	adjustedX := x % s.Width
	char := s.visibleArea[y][adjustedX]
	switch char {
	case '#':
		return "tree", nil
	case '.':
		return "empty", nil
	default:
		return "", fmt.Errorf("Unexpected Character %#U", char)
	}

}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("can't find file")
	}

	data := make([]byte, 20000)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	slope := NewSlope(data)
	x := 0
	y := 0
	xIncr := 3
	yIncr := 1

	trees := 0

	for y < slope.Height {
		maybeTree, err := slope.At(x, y)
		if err == nil && maybeTree == "tree" {
			trees++
		}

		x += xIncr
		y += yIncr
	}

	fmt.Printf("Trees: %d\n", trees)
}
