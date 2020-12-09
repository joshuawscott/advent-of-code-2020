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

func has(ss stringset, value string) bool {
	_, found := ss[value]
	return found
}

func intersection(s1, s2 stringset) stringset {
	result := make(stringset)
	for k := range s1 {
		_, found := s2[k]
		if found {
			add(result, k)
		}
	}
	return result
}

func intersect(ss []stringset) stringset {
	if len(ss) == 0 {
		return make(stringset, 0)
	}

	if len(ss) == 1 {
		return ss[0]
	}

	result := []stringset{intersection(ss[0], ss[1])}
	return intersect(append(result, ss[2:]...))
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("can't find file")
	}

	data := make([]byte, 20000)
	bytesRead, err := file.Read(data)
	data = data[:bytesRead]
	if err != nil {
		log.Fatal(err)
	}

	groups := strings.Split(string(data), "\n\n")

	totalQuestions := 0

	for _, group := range groups {
		// list of people's answers
		rawGroupAnswers := strings.Split(strings.TrimSpace(group), "\n")
		fmt.Println(rawGroupAnswers)
		commonAnswers := groupToCommonAnswers(rawGroupAnswers)

		totalQuestions += len(commonAnswers)

	}
	fmt.Println(totalQuestions)

}

// groupToCommonAnswers returns a stringset of common answers among all group memebers.
func groupToCommonAnswers(rawGroupAnswers []string) stringset {
	groupAnswers := make([]stringset, 0)
	for _, rawPersonAnswers := range rawGroupAnswers {
		ss := make(stringset)
		for _, r := range strings.Split(rawPersonAnswers, "") {
			add(ss, r)
		}
		groupAnswers = append(groupAnswers, ss)
	}
	commonAnswers := intersect(groupAnswers)
	fmt.Println(commonAnswers)
	return commonAnswers
}
