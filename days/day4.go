package days

import (
	"fmt"
	"os"
	"strings"
)

func wordSearch(matrix []string, x,y int) int {
	rows, cols := len(matrix), len(matrix[0])
	toFind := "XMAS"

	directions := [][]int {
		{0, 1},   // Horizontal right
		{0, -1},  // Horizontal left
		{1, 0},   // Vertical down
		{-1, 0},  // Vertical up
		{1, 1},   // Diagonal down-right
		{-1, -1}, // Diagonal up-left
		{1, -1},  // Diagonal down-left
		{-1, 1},  // Diagonal up-right
	}


	count := 0
	for _, d := range directions {
		found := true
		for k:=0;  k < len(toFind); k++ {
			newX, newY := x + k*d[0] , y + k*d[1]

			if newX < 0 || newY < 0 || newX>=rows || newY >= cols || matrix[newX][newY] != toFind[k] {
				found = false
				break
			}
		}
		if found {
			count++
		}
	}

	return count
}

func findX_MAS(matrix []string, x,y int) bool {
	rows, cols := len(matrix), len(matrix[0])
	if y == rows-1 {
		return false	
	}
	directions := [][] int {
		{x-1, y-1},
		{x-1, y+1},
		{x+1, y+1},
		{x+1, y-1},
	}

	// check for outer bounds 
	for _, dir := range directions {
		if dir[0] < 0 || dir[0] >= rows || dir[1] < 0 || dir[0] >= cols || matrix[dir[0]][dir[1]] == 'A' || matrix[dir[0]][dir[1]] == 'X' {
			return false
		}
	}



	// check for the X condition
	if matrix[directions[0][0]][directions[0][1]] == matrix[directions[2][0]][directions[2][1]] || matrix[directions[1][0]][directions[1][1]] == matrix[directions[3][0]][directions[3][1]] {
		return false
	}
	return true
}

func Day4() {
	data, err := os.ReadFile("inputs/data4dummy.txt")
	if err != nil {
		fmt.Printf("Error occured : %v\n",err)
	}

	strData := string(data)
	matrix := strings.Split(strings.TrimSpace(strData), "\n")



	totalCount := 0



	for i,row := range matrix {
		for j,c := range row {
			if c == 'A' && findX_MAS(matrix, i, j) {
				totalCount++
			}
		} 
	}
	fmt.Printf("%v\n",matrix[1][10])
	fmt.Printf("total XMAS count: %d\n",totalCount)
}
