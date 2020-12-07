package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.

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
	validatorFunctions := []func() bool{p.validByr, p.validIyr, p.validEyr, p.validHgt, p.validHcl, p.validEcl, p.validPid, p.validCid}
	for _, fn := range validatorFunctions {
		if fn() != true {
			//fmt.Printf("INV#%d -- %v\n", i, p)
			return false
		}
	}
	fmt.Printf("VALID -- %v\n", p)
	return true
}

func (p Passport) String() string {
	return fmt.Sprintf("byr: %s, iyr: %s, eyr: %s, hgt: %s, hcl: %s, ecl: %s, pid: %s, cid: %s",
		p.byr, p.iyr, p.eyr, p.hgt, p.hcl, p.ecl, p.pid, p.cid)
}

func (p Passport) validByr() bool {
	return validateNum(p.byr, 1920, 2002)
}

func (p Passport) validIyr() bool {
	return validateNum(p.iyr, 2010, 2020)
}

func (p Passport) validEyr() bool {
	return validateNum(p.eyr, 2020, 2030)
}

func (p Passport) validHgt() bool {
	switch {
	case strings.HasSuffix(p.hgt, "cm"):
		return validateNum(p.hgt[:len(p.hgt)-2], 150, 193)
	case strings.HasSuffix(p.hgt, "in"):
		return validateNum(p.hgt[:len(p.hgt)-2], 59, 76)
	default:
		return false
	}
}

func (p Passport) validHcl() bool {
	re := regexp.MustCompile("#[0-9a-f]{6}")
	return re.MatchString(p.hcl)
}

var validEyes = []string{"amb", "blu", "brn", "grn", "gry", "hzl", "oth"}

func (p Passport) validEcl() bool {
	idx := sort.SearchStrings(validEyes, p.ecl)
	return idx < len(validEyes) && validEyes[idx] == p.ecl
}

func (p Passport) validPid() bool {
	re := regexp.MustCompile("^[0-9]{9}$")
	return re.MatchString(p.pid)
}

func (p Passport) validCid() bool {
	return true
}

func validateNum(n string, min int, max int) bool {
	num, err := strconv.Atoi(n)
	return err == nil && num >= min && num <= max
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
