package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// SeatAssignment is the seat you are assigned.
type SeatAssignment struct {
	code    string
	rowCode string
	colCode string
	row     int64
	column  int64
	SeatID  int64
}

// NewSeatAssignment parses the raw code and returns a seat assignment.
func NewSeatAssignment(code string) (SeatAssignment, error) {
	// Cheating.
	rowCode := code[:7]
	rowCode = strings.ReplaceAll(rowCode, "B", "1")
	rowCode = strings.ReplaceAll(rowCode, "F", "0")
	row, err := strconv.ParseInt(rowCode, 2, 64)
	if err != nil {
		return SeatAssignment{}, errors.New("Failed parsing row code")
	}
	colCode := code[7:]
	colCode = strings.ReplaceAll(colCode, "R", "1")
	colCode = strings.ReplaceAll(colCode, "L", "0")
	col, err := strconv.ParseInt(colCode, 2, 64)
	if err != nil {
		return SeatAssignment{}, errors.New("Failed parsing column code")
	}

	seatID := row*8 + col

	return SeatAssignment{
		code:    code,
		rowCode: rowCode,
		colCode: rowCode,
		row:     row,
		column:  col,
		SeatID:  seatID,
	}, nil
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

	rawBoardingCodes := strings.Split(string(data), "\n")
	highestSeat := int64(0)
	lowestSeat := int64(math.MaxInt64)
	var seats = make([]SeatAssignment, 0)
	for _, rawBoardingCode := range rawBoardingCodes {
		sa, err := NewSeatAssignment(rawBoardingCode)
		if err != nil {
			continue // Move on if we have an unparsable line
		}
		if sa.SeatID > highestSeat {
			highestSeat = sa.SeatID
		}
		if sa.SeatID < lowestSeat {
			lowestSeat = sa.SeatID
		}
		seats = append(seats, sa)
	}

	fmt.Printf("highest: %d\n", highestSeat)
	fmt.Printf("lowest: %d\n", lowestSeat)
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].SeatID < seats[j].SeatID
	})

	var mySeat int64
	prevSeat := seats[0].SeatID
	for _, sa := range seats {
		if sa.SeatID-2 == prevSeat {
			mySeat = sa.SeatID - 1
			break
		}
		prevSeat = sa.SeatID
	}
	fmt.Printf("My seat: %d", mySeat)

}
