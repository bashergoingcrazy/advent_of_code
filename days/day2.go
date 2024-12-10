package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	// Read the file
	file, err := os.ReadFile("inputs/data.txt")
	if err != nil {
		fmt.Println("Error in opening file: ", err)
		return
	}

	// Split the file content into lines
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	// Solve the problem
	result := solveDay2(lines)
	fmt.Println("Day 2 result:", result)
}

// Main solving function for Day 2
func solveDay2(lines []string) int {
	safeCount := 0

	// Process each line
	for _, line := range lines {
		// Convert the line from a string to a slice of integers
		numStrs := strings.Fields(line) // Split by spaces
		var intLine []int
		for _, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				continue
			}
			intLine = append(intLine, num)
		}

		// Check if the line is safe or can be made safe with one removal
		if isSafe(intLine) || canBeSafeWithRemoval(intLine) {
			safeCount++
		}
	}

	return safeCount
}

// Function to check if a line is "safe"
func isSafe(line []int) bool {
	if len(line) < 2 {
		return false
	}
	isIncreasing := line[1] > line[0]
	for i := 1; i < len(line); i++ {
		diff := abs(line[i] - line[i-1])
		if diff < 1 || diff > 3 {
			return false
		}

		if (line[i] > line[i-1]) != isIncreasing {
			return false
		}
	}
	return true
}

// Function to check if a line can be made "safe" by removing one element
func canBeSafeWithRemoval(line []int) bool {
	n := len(line)
	if n < 3 {
		return false // Removing one element leaves too few elements
	}

	for i := 0; i < n; i++ {
		// Create a new slice excluding the current element
		newLine := make([]int, n)
		copy(newLine, line)
		modifiedLine := append(newLine[:i], newLine[i+1:]...)
		if isSafe(modifiedLine) {
			return true
		}
	}
	return false
}

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
