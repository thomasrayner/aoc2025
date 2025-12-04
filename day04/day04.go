package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = []struct{di, dj int}{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1},           {0, 1},
	{1, -1},  {1, 0},  {1, 1},
}

func countAccessible(grid []string) int {
	h := len(grid)
	if h == 0 {
		fmt.Println("Empty grid")
		return 0
	}
	w := len(grid[0])

	count := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] != '@' {
				continue
			}

			neighborRolls := 0
			for _, d := range dirs {
				ni, nj := i+d.di, j+d.dj
				if ni >= 0 && ni < h && nj >= 0 && nj < w && grid[ni][nj] == '@' {
					neighborRolls++
				}
			}

			if neighborRolls < 4 {
				count++
			}
		}
	}

	return count
}

func countRemovable(grid []string) int {
	h, w := len(grid), len(grid[0])
	g := make([][]byte, h)
	for i := range grid {
		g[i] = []byte(grid[i])
	}

	deg := make([][]int, h)
	for i := range deg {
		deg[i] = make([]int, w)
	}

	type pos struct{ i, j int }
	var q []pos

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if g[i][j] != '@' {
				continue
			}

			cnt := 0
			for _, d := range dirs {
				ni, nj := i+d.di, j+d.dj
				if ni >= 0 && ni < h && nj >= 0 && nj < w && g[ni][nj] == '@' {
					cnt++
				}
			}
			deg[i][j] = cnt
			if cnt < 4 {
				q = append(q, pos{i, j})
			}
		}
	}

	removed := 0
	for len(q) > 0 {
		size := len(q)
		for k := 0; k < size; k++ {
			p := q[k]
			i, j := p.i, p.j

			g[i][j] = '.'
			removed++

			for _, d := range dirs {
				ni, nj := p.i+d.di, p.j+d.dj
				if ni >= 0 && ni < h && nj >= 0 && nj < w && g[ni][nj] == '@' {
					deg[ni][nj]--
					if deg[ni][nj] == 3 {
						q = append(q, pos{ni, nj})
					}
				}
			}
		}
		q = q[size:]
	}

	return removed
}

func main() {
	pt1 := 0
	pt2 := 0

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	pt1 = countAccessible(grid)
	pt2 = countRemovable(grid)

	fmt.Println("Part 1:", pt1)
	fmt.Println("Part 2:", pt2)
}
