package days

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Equation struct {
	a, b, c int
}

func d13parse_data(data string) []soe {
	reg_val := regexp.MustCompile(`\+(\d+)`)
	reg_eq := regexp.MustCompile(`=(\d+)`)

	values := reg_val.FindAllStringSubmatch(data, -1)
	equates := reg_eq.FindAllStringSubmatch(data, -1)
	equations := make([]soe, 0, len(equates)/2)

	i, j := 0, 0
	for j < len(equates) {
		e1, e2 := Equation{}, Equation{}
		e1.a = to_int(values[i][1])
		e1.b = to_int(values[i+2][1])
		e1.c = -to_int(equates[j][1])

		e2.a = to_int(values[i+1][1])
		e2.b = to_int(values[i+3][1])
		e2.c = -to_int(equates[j+1][1])
		s := soe{e1, e2}
		equations = append(equations, s)
		i += 4
		j += 2
	}
	return equations
}

type soe struct {
	e1, e2 Equation
}

func to_int(d string) int {
	i, _ := strconv.Atoi(d)
	return i
}

func intersection(e1, e2 Equation) (int, int) {
	denominator := e1.a*e2.b - e2.a*e1.b
	xnumerator := e1.b*e2.c - e2.b*e1.c
	ynumerator := e2.a*e1.c - e1.a*e2.c

	if denominator == 0 {
		return -1, -1
	}

	if xnumerator%denominator == 0 && ynumerator%denominator == 0 {
		x, y := xnumerator/denominator, ynumerator/denominator
		// fmt.Println(e1,e2,x,y)
		return x, y
	}
	return -1, -1
}

func solve1(eqs []soe) int {
	count := 0
	for _, eq := range eqs {
		x, y := intersection(eq.e1, eq.e2)
		// fmt.Printf("x: %d, y: %d\n",x,y)
		if x == -1 && y == -1 {
			continue
		}
		if x > 100 || y > 100 || x < 0 || y < 0 {
			continue
		}
		count += (x * 3) + y
	}
	return count
}

func solve2(eqs []soe) int {
	count := 0
	for _, eq := range eqs {
		eq.e1.c -= int(math.Pow(10, 13))
		eq.e2.c -= int(math.Pow(10, 13))
		// fmt.Println(eq)
		x, y := intersection(eq.e1, eq.e2)
		// fmt.Printf("x: %d, y: %d\n",x,y)
		if x == -1 && y == -1 {
			continue
		}
		if x < 0 || y < 0 {
			continue
		}
		count += (x * 3) + y
	}
	return count
}

func Day13() {
	t := time.Now()
	d, err := os.ReadFile("inputs/data13.txt")
	if err != nil {
		fmt.Println("Can't read file: ", err)
		return
	}
	data := string(d)
	equations := d13parse_data(data)
	fmt.Println(solve1(equations))
	fmt.Println(solve2(equations))
	ts := time.Now()
	elapsed := ts.Sub(t)
	fmt.Println("Time taken:", elapsed)
}
