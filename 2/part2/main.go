package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type policy struct {
	// 0-indexed
	FirstPos int
	// 0-indexed
	LastPos int
	Value   string
}

func (p policy) Valid(password string) bool {
	first := string(password[p.FirstPos])
	last := string(password[p.LastPos])
	fmt.Printf("password: '%s', policyValue: %s, positions: %d/%d, first: %s, last: %s\n", password, p.Value, p.FirstPos, p.LastPos, first, last)
	// XOR
	return (first == p.Value) != (last == p.Value)
}

var re *regexp.Regexp = regexp.MustCompile(`(\d+)-(\d+) (.): (.+)`)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("can't find file")
	}

	data := make([]byte, 30000)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	strValues := strings.Split(string(data), "\n")
	validPasswords := 0
	for _, v := range strValues {
		x := re.FindStringSubmatch(v)
		if len(x) != 5 {
			fmt.Printf("Ignoring line: '%s'\n", v)
			continue
		}
		firstPos, err1 := strconv.Atoi(x[1])
		lastPos, err2 := strconv.Atoi(x[2])
		if err1 == nil && err2 == nil {
			p := policy{
				FirstPos: firstPos - 1,
				LastPos:  lastPos - 1,
				Value:    x[3],
			}
			password := x[4]
			if p.Valid(password) {
				validPasswords++
			}
		}
	}
	fmt.Printf("found %d valid\n", validPasswords)
}
