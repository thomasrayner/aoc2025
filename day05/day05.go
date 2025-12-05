package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	low  int
	high int
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].low < ranges[j].low
	})

	merged := []Range{ranges[0]}

	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]

		if r.low <= last.high+1 {
			if r.high > last.high {
				last.high = r.high
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}

func isFresh(c int, ranges []Range) bool {
	for _, r := range ranges {
		if c >= r.low && c <= r.high {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ranges []Range
	var ids []int

	doingRanges := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			doingRanges = false
			continue
		}

		if doingRanges {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				panic("bad range line: " + line)
			}
			low, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			high, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			if high < low {
				low, high = high, low
			}
			ranges = append(ranges, Range{low, high})
		} else {
			v, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ids = append(ids, v)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	merged := mergeRanges(ranges)

	pt1 := 0
	for _, id := range ids {
		if isFresh(id, merged) {
			pt1++
		}
	}

	pt2 := 0
	for _, r := range merged {
		pt2 += (r.high - r.low + 1)
	}

	fmt.Println("Pt1:", pt1)
	fmt.Println("Pt2:", pt2)
}
