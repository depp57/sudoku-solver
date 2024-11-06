package lib

import (
	"fmt"
	"slices"
	"time"
)

func Solve(grid [][]int) (bool, [][]int) {
	defer timer("backtracking")()
	return backTracking(grid, 0, 0)
}

func backTracking(grid [][]int, row int, col int) (bool, [][]int) {
	if row == 9 {
		return true, grid
	} else if col == 9 {
		return backTracking(grid, row+1, 0)
	} else if grid[row][col] != 0 {
		return backTracking(grid, row, col+1)
	} else {
		for currentNum := 1; currentNum <= 9; currentNum++ {
			if isValid(grid, row, col, currentNum) {
				grid[row][col] = currentNum

				completed, _ := backTracking(grid, row, col+1)
				if completed {
					return true, grid
				} else {
					grid[row][col] = 0
				}
			}
		}

		return false, grid
	}
}

func isValid(grid [][]int, row, col, num int) bool {
	// is cell empty?
	if grid[row][col] != 0 {
		return false
	}

	// is row valid?
	if slices.Contains(grid[row], num) {
		return false
	}

	// is column valid?
	for i := 0; i < len(grid); i++ {
		if grid[i][col] == num {
			return false
		}
	}

	// is box valid?
	boxRow := row / 3
	boxCol := col / 3

	boxRowLowerBouding := boxRow * 3
	boxRowUpperBouding := boxRow*3 + 2
	boxColLowerBouding := boxCol * 3
	boxColUpperBouding := boxCol*3 + 2

	for row := boxRowLowerBouding; row <= boxRowUpperBouding; row++ {
		for col := boxColLowerBouding; col <= boxColUpperBouding; col++ {
			if grid[row][col] == num {
				return false
			}
		}
	}

	return true
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
