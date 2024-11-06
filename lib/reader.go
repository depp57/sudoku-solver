package lib

import (
	"encoding/csv"
	"os"
	"strconv"
)

type GridReader interface {
	ReadGrid() ([][]int, error)
}

type CsvGridReader struct {
	FileName string
}

func (r CsvGridReader) ReadGrid() ([][]int, error) {
	file, err := os.Open(r.FileName)

	if err != nil {
		panic("Error while opening grid file")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil || len(records) != 9 || len(records[0]) != 9 {
		panic("Error while parsing grid file")
	}

	grid := make([][]int, len(records))

	for row := 0; row < len(records); row++ {
		for col := 0; col < len(records[row]); col++ {
			cellAsInt, err := strconv.Atoi(records[row][col])
			if err != nil {
				panic("Error while parsing cell")
			}

			grid[row] = append(grid[row], cellAsInt)
		}
	}

	return grid, nil
}
