package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type void struct{}
type stringset map[string]void

var member void

func add(ss stringset, s string) {
	re := regexp.MustCompile("[a-z]")
	if !re.MatchString(s) {
		return
	}
	ss[s] = member
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

	groups := strings.Split(string(data), "\n\n")

	totalQuestions := 0

	for _, group := range groups {
		ss := make(stringset, 0)
		group = strings.ReplaceAll(group, "\n", "")
		for _, r := range strings.Split(group, "") {
			add(ss, r)
		}
		fmt.Printf("%d: %s\n", len(ss), group)
		totalQuestions += len(ss)
	}
	fmt.Println(totalQuestions)

}
