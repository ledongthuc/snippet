package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const preambleLength = 25

func main() {
	content, err := ParseInput("./numbers.txt")
	if err != nil {
		panic(err)
	}
	numbers := ParseNumbers(content)
	invalidNumber := findFirstInvalidNumber(numbers)
	contiguousNumbers := findContiguousContiguousNumbers(numbers, invalidNumber)
	sort.Slice(contiguousNumbers, func(i, j int) bool {
		return contiguousNumbers[i] < contiguousNumbers[j]
	})
	fmt.Println(contiguousNumbers[0]+contiguousNumbers[len(contiguousNumbers)-1], contiguousNumbers)
}

func ParseInput(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ParseNumbers(s string) []int64 {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	numbers := make([]int64, 0, len(lines))
	for _, line := range lines {
		v, err := strconv.ParseInt(line, 10, 64)
		if err == nil {
			numbers = append(numbers, v)
		}
	}
	return numbers
}

func findFirstInvalidNumber(numbers []int64) int64 {
	if len(numbers) <= preambleLength {
		return -1
	}
	for i := preambleLength; i < len(numbers); i++ {
		valid, _, _ := isCombinationOf(numbers[i], numbers[i-preambleLength:i])
		if !valid {
			return numbers[i]
		}
	}
	return -1
}

func isCombinationOf(number int64, preambleNumbers []int64) (bool, int64, int64) {
	for i := 0; i < len(preambleNumbers)-1; i++ {
		for j := i + 1; j < len(preambleNumbers); j++ {
			if preambleNumbers[i] != preambleNumbers[j] && preambleNumbers[i]+preambleNumbers[j] == number {
				return true, preambleNumbers[i], preambleNumbers[j]
			}
		}
	}
	return false, 0, 0
}

func findContiguousContiguousNumbers(numbers []int64, invalidNumber int64) []int64 {
	for i := 0; i < len(numbers)-1; i++ {
		value := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			value += numbers[j]
			if value == invalidNumber {
				return numbers[i : j+1]
			}

			if value > invalidNumber {
				break
			}
		}
	}
	return []int64{}
}
