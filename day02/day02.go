package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func containsDouble(first string, second string) (int, error) {
	ret := 0
	f, err := strconv.Atoi(first)
	if err != nil {
		return 0, err
	}
	l, err := strconv.Atoi(second)
	if err != nil {
		return 0, err
	}

	if l < f { 
		return 0, errors.New("last number is less than first number: " + first + "-" + second)
	}

	for i := f; i <= l; i++ {
		numStr := strconv.Itoa(i)
		splitSpot := len(numStr) / 2

		firstHalf := numStr[:splitSpot]
		secondHalf := numStr[splitSpot:]

		if firstHalf == secondHalf {
			ret += i
		}
	}

	return ret, nil
}

func containsRepeatingPattern(first string, second string) (int, error) {
	ret := 0
	f, err := strconv.Atoi(first)
	if err != nil {
		return 0, err
	}
	l, err := strconv.Atoi(second)
	if err != nil {
		return 0, err
	}

	if l < f { 
		return 0, errors.New("last number is less than first number: " + first + "-" + second)
	}

	for i := f; i <= l; i++ {
		numStr := strconv.Itoa(i)

		divisors := []int{}
		for d := 1; d < len(numStr); d++ {
			if len(numStr)%d == 0 {
				divisors = append(divisors, d)
			}
		}

		for _, d := range divisors {
			pattern := strings.Join(strings.Split(numStr, "")[0:d], "")
			matches := true
			for j := d; j < len(numStr); j += d {
				if numStr[j:j+d] != pattern {
					matches = false
					break
				}
			}
			if matches {
				ret += i
				break
			}
		}
	}

	return ret, nil
}

func main() {
	pt1 := 0
	pt2 := 0
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, pair := range line {
			first := strings.Split(pair, "-")[0]
			second := strings.Split(pair, "-")[1]

			result1, err := containsDouble(first, second)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			pt1 += result1

			result2, err := containsRepeatingPattern(first, second)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			pt2 += result2
		}
	}

	fmt.Println("Pt1:", pt1)
	fmt.Println("Pt2:", pt2)
}
