package main

import (
	"fmt"
)

func main() {
	grid := []string{"abc", "ade", "efg"}
	fmt.Println(gridChallenge(grid))
}

func gridChallenge(grid []string) string {
	var column = make([]string, len(grid))
	var matrix = make([][]string, 0)

	if len(grid) == 1 {
		return "YES"
	} else {
		for i := 0; i < len(grid); i++ {
			row := grid[i]
			column = make([]string, 0)
			for j := 0; j < len(row); j++ {
				column = append(column, row[j:j+1])
			}
			matrix = append(matrix, column)
		}

		//fmt.Println("matrix: ", matrix)

		for i := 1; i < len(matrix); i++ {

			for j := 0; j < len(matrix); j++ {
				fmt.Println("comparisson ", matrix[i][j], matrix[i-1][j])

				if matrix[i][j] < matrix[i-1][j] {
					return "NO"
				}
			}
		}
		return "YES"
	}
}
