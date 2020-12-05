package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("can't find file")
	}

	data := make([]byte, 64000)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	strValues := strings.Split(string(data), "\n")
	values := make([]int, 0)
	for _, v := range strValues {
		num, err := strconv.Atoi(v)
		if err == nil {
			values = append(values, num)
		}

	}
	sort.Ints(values)
	end := len(values) - 1

	for i := 0; i < end; {
		small := values[i]
		large := values[end]
		total := small + large
		if total == 2020 {
			fmt.Printf("found %d & %d. Total: %d\n", small, large, small*large)
			break
		} else if total > 2020 {
			end--
		} else {
			i++
		}
	}
}
