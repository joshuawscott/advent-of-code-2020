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

	func() {
		for i := 0; i < len(values); i++ {
			for j := 0; j < len(values); j++ {
				if j <= i {
					j = i + 1
				}
				targetNumber := 2020 - (values[i] + values[j])
				if targetNumber < values[j] {
					// We will already have evaluated this combination in reverse.
					break
				}
				found, idx := indexOf(targetNumber, values)
				if found && idx != i && idx != j {
					fmt.Printf("Found: %d, %d, %d. Total: %d\n", values[i], values[j], values[idx], values[i]*values[j]*values[idx])
					return // All done
				}
			}
		}
	}()
}

// Takes a sorted slice of ints and determines if a number is in the slice.
func indexOf(target int, list []int) (bool, int) {
	idx := sort.SearchInts(list, target)
	fmt.Println(idx)
	if idx < len(list) && list[idx] == target {
		return true, idx
	}
	return false, -1
}
