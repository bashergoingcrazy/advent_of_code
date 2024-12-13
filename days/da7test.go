package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func CombineNums(a, b int) int {
	power := 1
	for b >= power {
		power *= 10
	}
	return a*power + b
}

func EvalCalibration(nums []int, solution int, combined bool) int {
	var solve func(i int, n int) int
	solve = func(i int, n int) int {
		if i < len(nums) {
			sum := solve(i+1, n+nums[i])
			product := solve(i+1, n*nums[i])
			if combined {
				return sum + product + solve(i+1, CombineNums(n, nums[i]))
			}
			return sum + product
		}

		if n == solution {
			return 1
		}
		return 0
	}

	if solve(1, nums[0]) != 0 {
		return solution
	} else {
		return 0
	}
}

func Day07() (int64, int64) {
	start := time.Now()
	in, err := os.ReadFile("inputs/data7.txt")
	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	var calibration, calibrationWithConcat int64

	for _, line := range strings.Split(string(in), "\n") {
		line := strings.Split(line, ":")
		if len(line) < 2 {
			continue
		}

		nums := make([]int, 0)
		solution, _ := strconv.Atoi(line[0])
		for _, num := range strings.Fields(line[1]) {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}

		calibration += int64(EvalCalibration(nums, solution, false))
		calibrationWithConcat += int64(EvalCalibration(nums, solution, true))
	}

	t := time.Now()
	fmt.Printf("Elapsed time: %v\n", t.Sub(start))
	return calibration, calibrationWithConcat
}
