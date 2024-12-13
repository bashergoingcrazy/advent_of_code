package days

import (
	"fmt"
	"os"
	"strings"
)

func d10print_map(matrix [][]rune, karr ...Key) {
	var k Key
	if len(karr) >= 1 {
		k = karr[0]
	} else {
		k = Key{-1, -1, 'c'}
	}

	for i, row := range matrix {
		for j, val := range row {
			if i == k.x && j == k.y {
				s := string(val)
				fmt.Printf(Color(s, Red))
				fmt.Printf(" ")
			} else {
				fmt.Printf("%c ", val)
			}
		}
		fmt.Printf("\n")
	}
	// fmt.Printf("%d,%d\n",len(matrix), len(matrix[0]))
}

var direct = [][]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

type Key struct {
	x, y int
	r    rune
}

var memo map[Key]int

// this whole thing can be memoed for search of value for diffrent trailheads
func d10traverse(curr rune, i, j int, matrix [][]rune, vis map[Key]bool) int {
	k := Key{i, j, curr}

	// Memo check
	if val, exists := memo[k]; exists {
		return val
	}

	// Base case checks
	m, n := len(matrix), len(matrix[0])
	if i < 0 || i >= m || j < 0 || j >= n || curr != matrix[i][j] {
		return 0
	}

	// Final case: Reached '9'
	if curr == '9' {
		memo[k] = 1 // Cache the result
		return 1
	}

	// Mark as visited for this path

	// Traverse to adjacent neighbors
	count := 0
	for _, dir := range direct {
		newX, newY := i+dir[0], j+dir[1]
		if newX >= 0 && newX < m && newY >= 0 && newY < n && matrix[newX][newY] == curr+1 {
			count += d10traverse(curr+1, newX, newY, matrix, vis)
		}
	}

	// Unmark visited to allow other paths to use this node

	// Cache the result
	memo[k] = count
	return count
}

func d10parse_input(data string) [][]rune {
	ds := strings.Split(strings.TrimSpace(data), "\n")
	matrix := make([][]rune, 0, len(ds))
	for _, row := range ds {
		// rs := strings.Split(row, "")
		rrow := []rune(row)
		matrix = append(matrix, rrow)
	}
	return matrix
}

func Day10() {
	d, err := os.ReadFile("inputs/data10dummy.txt")
	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return
	}
	data := string(d)
	matrix := d10parse_input(data)

	totalCount := 0
	memo = make(map[Key]int)
	for i, row := range matrix {
		for j, r := range row {
			if r == '0' {
				vis := make(map[Key]bool)
				totalCount += d10traverse(r, i, j, matrix, vis)
			}
		}
	}
	d10print_map(matrix)
	fmt.Println(Color(totalCount))
}
