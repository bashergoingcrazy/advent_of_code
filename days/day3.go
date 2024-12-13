package days

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day3() {
	data, err := os.ReadFile("inputs/data3.txt")
	if err != nil {
		fmt.Println("Error in reading the file ", err)
		return
	}

	// Regex patterns
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	dore := regexp.MustCompile(`do\(\)`)
	dontre := regexp.MustCompile(`don't\(\)`)

	// Parse data
	stringData := string(data)
	testMatches := re.FindAllStringSubmatchIndex(stringData, -1)
	doMatch := dore.FindAllStringIndex(stringData, -1)
	dontMatch := dontre.FindAllStringIndex(stringData, -1)

	fmt.Println("Input Data:", stringData)
	fmt.Println("mul Matches:", testMatches)
	fmt.Println("do() Matches:", doMatch)
	fmt.Println("don't() Matches:", dontMatch)

	// Variables
	sum := 0
	i, j, k := 0, 0, 0
	a, b, c := len(testMatches), len(doMatch), len(dontMatch)
	enabled := true

	// Process each character in stringData
	for x := 0; x < len(stringData); x++ {
		// Debug: Show current enabled state
		if enabled {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}

		// Check for don't() at current index
		if k < c && dontMatch[k][0] == x {
			enabled = false
			fmt.Println("\nDon't() encountered at index", x, "=> Disabled mul()")
			k++
		}

		// Check for do() at current index
		if j < b && doMatch[j][0] == x {
			enabled = true
			fmt.Println("\nDo() encountered at index", x, "=> Enabled mul()")
			j++
		}

		// Check for mul() at current index
		if i < a && testMatches[i][0] == x {
			if enabled {
				num1, _ := strconv.Atoi(stringData[testMatches[i][2]:testMatches[i][3]])
				num2, _ := strconv.Atoi(stringData[testMatches[i][4]:testMatches[i][5]])
				sum += num1 * num2
				fmt.Printf("\nProcessing mul(%d,%d) at index %d => Result: %d\n", num1, num2, x, num1*num2)
			} else {
				fmt.Printf("\nSkipping mul at index %d => Disabled\n", x)
			}
			i++
		}
	}

	// Final result
	fmt.Println("\nFinal Sum:", sum)
}
