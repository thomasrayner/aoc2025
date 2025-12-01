package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	pointer := 50
	countzero := 0
	countpastzero := 0

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}

		// if the pointer passes zero or a 100 multiple, count it with countpastzero
		if line[0] == 'L' {
			for i := 0; i < num; i++ {
				if (pointer - i) == 0 || (pointer - i) % 100 == 0 {
					countpastzero++
				}
			}
			pointer -= num
		} else if line[0] == 'R' {
			for i := 0; i < num; i++ {
				if (pointer + i) == 0 || (pointer + i) % 100 == 0 {
					countpastzero++
				}
			}
			pointer += num
		}

		if pointer == 0 || pointer % 100 == 0 {
			countzero++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Pt1: %d\n", countzero)
	fmt.Printf("Pt2: %d\n", countpastzero)
}
