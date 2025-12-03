package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const K = 12

func largestKDigits(s string, k int) string {
	n := len(s)
	if k == 0 {
		return "0"
	}
	if k >= n {
		return s
	}

	stack := []byte{}
	toRemove := n - k

	for i := 0; i < n; i++ {
		digit := s[i]

		for toRemove > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			toRemove--
		}

		stack = append(stack, digit)
	}

	if toRemove > 0 {
		stack = stack[:len(stack)-toRemove]
	}

	return string(stack)
}

func maxJoltage(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	best := 0
	maxAfter := 0

	for i := n - 1; i >= 0; i-- {
		curr := int(s[i] - '0')

		if i < n-1 {
			candidate := curr*10 + maxAfter
			if candidate > best {
				best = candidate
			}
		}

		if curr > maxAfter {
			maxAfter = curr
		}
	}

	return best
}

func main() {
	pt1 := 0
	var pt2 int64 = 0

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		pt1 += maxJoltage(line)
		largest := largestKDigits(line, K)
		num, err := strconv.ParseInt(largest, 10, 64)
		if err != nil {
			panic(err)
		}

		pt2 += num
	}

	fmt.Println("Part 1:", pt1)
	fmt.Println("Part 2:", pt2)
}
