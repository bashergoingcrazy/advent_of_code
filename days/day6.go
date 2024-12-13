package days

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func d6_printMatrix(bMatrix [][]string) {
	for _, rows := range bMatrix {
		fmt.Printf("%v\n", rows)
	}
	fmt.Printf("%d,%d\n", len(bMatrix), len(bMatrix[0]))
}

func printDynamicGrid(grid [][]string) {
	fmt.Print("\033[H\033[2J") // Clear terminal
	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}
	time.Sleep(100 * time.Millisecond) // Adjust delay for smooth visualization
}

func d6_find_guard_index(bMatrix [][]string) (int, int) {
	for i, row := range bMatrix {
		for j, val := range row {
			if val == "^" {
				return i, j
			}

		}
	}
	return -1, -1
}

func d6_traverse(bMatrix [][]string) int {
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	mp := make(map[string]bool)
	dIndex := 0
	i, j := d6_find_guard_index(bMatrix)
	m, n := len(bMatrix), len(bMatrix[0])

	// copy of the first
	cMatrix := make([][]string, 0, len(bMatrix))
	for _, k := range bMatrix {
		copiedSlices := make([]string, len(k))
		copy(copiedSlices, k)
		cMatrix = append(cMatrix, copiedSlices)
	}

	for i >= 0 && i < m && j >= 0 && j < n {
		if bMatrix[i][j] == "#" {
			// an obstacle encountered
			i, j = i-directions[dIndex][0], j-directions[dIndex][1]
			dIndex = (dIndex + 1) % 4
			i, j = i+directions[dIndex][0], j+directions[dIndex][1]
		}

		key := fmt.Sprintf("%d,%d", i, j)
		mp[key] = true
		cMatrix[i][j] = "X"
		printDynamicGrid(cMatrix)
		cMatrix[i][j] = "."
		i, j = i+directions[dIndex][0], j+directions[dIndex][1]
	}
	d6_printMatrix(cMatrix)
	return len(mp)
}

func d6_parseInput(data string) [][]string {
	d := strings.Split(strings.TrimSpace(data), "\n")
	matrix := make([][]string, 0, len(d))

	for _, k := range d {
		row := strings.Split(k, "")
		matrix = append(matrix, row)
	}
	return matrix
}

func Day6() {
	d, err := os.ReadFile("inputs/data6dummy.txt")
	if err != nil {
		fmt.Println("Err while opening the file", err)
	}
	data := string(d)

	matrix := d6_parseInput(data)
	d6_printMatrix(matrix)
	fmt.Printf("%d\n", d6_traverse(matrix))
	d6_printMatrix(matrix)

}
