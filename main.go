package main

import (
	"fmt"
	"sudoku-solver/lib"
)

func main() {
	grid, _ := lib.CsvGridReader{FileName: "data/3.csv"}.ReadGrid()

	fmt.Println("Initial grid:")
	prettyPrintGrid(grid)
	fmt.Println()

	fmt.Println("Solved grid:")
	isSolvable, grid := lib.Solve(grid)
	if !isSolvable {
		fmt.Println("The grid is not solvable")
		return
	}
	prettyPrintGrid(grid)
}

func prettyPrintGrid(grid [][]int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			fmt.Print(grid[row][col], "  ")
		}
		fmt.Println()
	}
}
