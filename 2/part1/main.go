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
	MinTimes int
	MaxTimes int
	Value    string
}

func (p policy) Valid(password string) bool {
	occurrences := 0
	for _, v := range password {
		if p.Value == string(v) {
			occurrences++
		}
	}
	return occurrences >= p.MinTimes && occurrences <= p.MaxTimes
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
		minTimes, err1 := strconv.Atoi(x[1])
		maxTimes, err2 := strconv.Atoi(x[2])
		if err1 == nil && err2 == nil {
			p := policy{
				MinTimes: minTimes,
				MaxTimes: maxTimes,
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
