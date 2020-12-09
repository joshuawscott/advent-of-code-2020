package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var re *regexp.Regexp = regexp.MustCompile(" *[0-9] ")

type void struct{}
type stringset map[string]void
type legalContainers map[string]stringset

var member void

// returns true if added, false otherwise
func add(ss stringset, value string) bool {
	_, found := ss[value]
	if found {
		return false
	}
	ss[value] = member
	return true
}

func parseRule(rawRule string) (outer string, inners []string) {
	rawRule = strings.ReplaceAll(rawRule, "bags", "")
	rawRule = strings.ReplaceAll(rawRule, "bag", "")
	// Remove numbers
	rawRule = strings.TrimSuffix(rawRule, ".")
	rawRule = strings.TrimSpace(rawRule)

	s := strings.Split(rawRule, " contain ")
	if len(s) != 2 {
		return outer, inners
	}
	outer = strings.TrimSpace(s[0])
	if s[1] == "no other bags." {
		inners = make([]string, 0)
	} else {
		inners = strings.Split(s[1], ",")
	}
	for i := range inners {
		inners[i] = re.ReplaceAllString(inners[i], "")
		inners[i] = strings.TrimSpace(inners[i])
	}
	return outer, inners
}

func addContainer(lc legalContainers, outer string, inner string) {
	if lc[inner] == nil {
		lc[inner] = make(stringset)
	}
	add(lc[inner], outer)
}

func findContainersFor(lc legalContainers, bag string, ss stringset) {
	for outer := range lc[bag] {
		if add(ss, outer) {
			findContainersFor(lc, outer, ss)
		}
	}
}

const myBag string = "shiny gold"

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("can't find file")
	}

	data := make([]byte, 50000)
	bytesRead, err := file.Read(data)
	data = data[:bytesRead]
	if err != nil {
		log.Fatal(err)
	}

	rawRules := strings.Split(string(data), "\n")
	lc := make(legalContainers)
	for _, rawRule := range rawRules {
		outer, allInner := parseRule(rawRule)
		for _, inner := range allInner {
			addContainer(lc, outer, inner)
		}
	}
	ss := make(stringset)
	findContainersFor(lc, myBag, ss)
	fmt.Println(len(ss))
}
