package days

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day11to_int(nums []string) []int {
	arr := make([]int, 0, len(nums))
	for _, val := range nums {
		f, _ := strconv.Atoi(val)
		arr = append(arr, f)
	}
	return arr
}

func day11num_digits(num int) (int, bool) {
	x := float64(num)
	n := math.Floor(math.Log10(x)) + 1
	y := int(n)
	isEven := y%2 == 0
	return y, isEven
}
func day11split_half(num, n int) (int, int) {
	pow := int(math.Pow10(n / 2))
	lsbNum := num % pow
	msbNum := num / pow
	return msbNum, lsbNum
}

func blink(nums []int) []int {
	arr := []int{}
	for _, val := range nums {
		if val == 0 {
			arr = append(arr, 1)
		} else if n, isEven := day11num_digits(val); isEven {
			f, s := day11split_half(val, n)
			arr = append(arr, f)
			arr = append(arr, s)
		} else {
			arr = append(arr, 2024*val)
		}
	}
	return arr
}

func d11blink_iterations(m map[int]int, iter int) map[int]int {
	for range iter {
		nm := make(map[int]int)
		for key, val := range m {
			if key == 0 {
				nm[1] += val
			} else if n, isEven := day11num_digits(key); isEven {
				f, s := day11split_half(key, n)
				nm[f] += val
				nm[s] += val
			} else {
				nm[2024*key] += val
			}
		}
		m = nm
	}
	return m
}

func Day11() {
	d, err := os.ReadFile("inputs/data11.txt")
	if err != nil {
		fmt.Println("error in opening file: ", err)
		return
	}
	data := string(d)
	nu := strings.Fields(data)
	nums := day11to_int(nu)
	fmt.Println(nums)
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	f := d11blink_iterations(m, 25)
	count := 0
	for _, val := range f {
		count += val
	}
	fmt.Println(len(f))
	fmt.Println(count)
}
