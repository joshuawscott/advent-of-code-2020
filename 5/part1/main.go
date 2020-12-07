package main

import (
	"fmt"
	"log"
	"os"
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
func NewSeatAssignment(code string) SeatAssignment {
	// Cheating.
	rowCode := code[:7]
	rowCode = strings.ReplaceAll(rowCode, "B", "1")
	rowCode = strings.ReplaceAll(rowCode, "F", "0")
	row, err := strconv.ParseInt(rowCode, 2, 64)
	if err != nil {
		//fmt.Printf("Failed parsing row from %s", code)
	}
	colCode := code[7:]
	colCode = strings.ReplaceAll(colCode, "R", "1")
	colCode = strings.ReplaceAll(colCode, "L", "0")
	col, err := strconv.ParseInt(colCode, 2, 64)
	if err != nil {
		//fmt.Printf("Failed parsing column from %s", code)
	}

	seatID := row*8 + col

	return SeatAssignment{
		code:    code,
		rowCode: rowCode,
		colCode: rowCode,
		row:     row,
		column:  col,
		SeatID:  seatID,
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

	rawBoardingCodes := strings.Split(string(data), "\n")
	highestSeat := int64(0)
	for _, rawBoardingCode := range rawBoardingCodes {
		sa := NewSeatAssignment(rawBoardingCode)
		if sa.SeatID > highestSeat {
			highestSeat = sa.SeatID
		}
	}

	fmt.Printf("Result: %d\n", highestSeat)
}
