package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Passport is a passport.
type Passport struct {
	byr string // (Birth Year)
	iyr string // (Issue Year)
	eyr string // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid string // (Country ID)
}

// NewPassport constructs a Passport from raw data.
func NewPassport(rawPassport string) Passport {
	pp := Passport{}
	// Clean up the newline vs space delimiters.
	rawPassport = strings.ReplaceAll(rawPassport, "\n", " ")
	fields := strings.Split(rawPassport, " ")
	for _, field := range fields {
		key, value, err := parseField(field)
		if err != nil {
			// Ignore unparsable fields.
			continue
		}
		switch key {
		case "byr":
			pp.byr = value
		case "iyr":
			pp.iyr = value
		case "eyr":
			pp.eyr = value
		case "hgt":
			pp.hgt = value
		case "hcl":
			pp.hcl = value
		case "ecl":
			pp.ecl = value
		case "pid":
			pp.pid = value
		case "cid":
			pp.cid = value
		}
	}
	return pp
}

// Returns key, value, maybeError
func parseField(field string) (string, string, error) {
	sl := strings.Split(field, ":")
	if len(sl) != 2 {
		return "", "", errors.New("Invalid field, more than one ':'")
	}
	return sl[0], sl[1], nil
}

// Valid returns true if passport is valid
func (p Passport) Valid() bool {
	return p.validByr() && p.validIyr() && p.validEyr() && p.validHgt() && p.validHcl() && p.validEcl() && p.validPid() && p.validCid()
}

func (p Passport) validByr() bool {
	return p.byr != ""
}

func (p Passport) validIyr() bool {
	return p.iyr != ""
}

func (p Passport) validEyr() bool {
	return p.eyr != ""
}

func (p Passport) validHgt() bool {
	return p.hgt != ""
}

func (p Passport) validHcl() bool {
	return p.hcl != ""
}

func (p Passport) validEcl() bool {
	return p.ecl != ""
}

func (p Passport) validPid() bool {
	return p.pid != ""
}

func (p Passport) validCid() bool {
	return true
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

	rawPassports := strings.Split(string(data), "\n\n")
	validPassports := 0
	for _, rawPassport := range rawPassports {
		pp := NewPassport(rawPassport)
		if pp.Valid() {
			validPassports++
		}
	}

	fmt.Printf("Result: %d\n", validPassports)
}
