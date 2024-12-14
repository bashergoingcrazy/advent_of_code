package days


import (
	"fmt"
	"os"
	"regexp"
    "github.com/bashergoingcrazy/goviz/color"
)

type Robot struct{
    point Pair
    vel Pair 
}

func print_pos(ps []Pair, m,n int) {
    for i:=0; i<m; i++ {
        for j:=0; j<n;j++ {
            count := 0
            for _, p:= range ps {
                if p.x == j && p.y == i {
                    count++
                }
            }
            switch count {
            case 0:
                fmt.Printf(color.Color(count, color.Blue))
            default:
                fmt.Printf(color.Color(count, color.Red))
            }
        }
        fmt.Println()
    }
}

func after_iterations(i int, r Robot, m,n int) Pair {
    av1 := (r.point.x +((r.vel.x * i)%n)+n)%n
    av2 := (r.point.y +((r.vel.y * i)%m)+m)%m
    
    // fmt.Printf("av1: %d, av2: %d\n",av1,av2)
    return Pair{av1,av2}
}

func d14part1(robots []Robot) int {
    m,n := 103,101
    ps := make([]Pair,0,len(robots))
    for _,r := range robots {
        ps = append(ps, after_iterations(100,r,m,n))
    }
    print_pos(ps, m,n)

    quad := make([]int,4)
    for _, p:= range ps {
        if p.x < n/2 && p.y < m/2 {
            quad[0]++
        } 
        if p.x > n/2 && p.y < m/2 {
            quad[1]++
        } 
        if p.x > n/2 && p.y > m/2 {
            quad[2]++
        } 
        if p.x < n/2 && p.y > m/2 {
            quad[3]++
        } 
    }
    safety := 1
    for i:= 0; i<4; i++ {
        safety *= quad[i]
        fmt.Println(quad[i])
    }
    return safety
}

func find_easter_egg(ps []Pair, m, n int, iteration int) bool {
    matrix := make([][]int, m)
    for i := range matrix {
        matrix[i] = make([]int, n)
    }

    for _, p := range ps {
        matrix[p.y][p.x]++
    }

    // Define your easter egg pattern (example: tree-shaped pattern)
    for i := 1; i < m-4; i++ {
        for j := 1; j < n-1; j++ {
            // Check for a specific pattern (e.g., "tree" structure)
            if matrix[i][j] != 0 &&
                matrix[i-1][j] != 0 &&
                matrix[i+1][j] != 0 &&
                matrix[i+2][j] != 0 &&
                matrix[i+3][j] != 0 &&
                matrix[i][j-1] != 0 &&
                matrix[i][j+1] != 0 {
                fmt.Printf("Easter egg found at iteration %d at position (%d, %d)\n", iteration, j, i)
                return true
            }
        }
    }
    return false
}

func d14part2_with_detection(robots []Robot) {
    m, n := 103, 101
    for i := 0; i < 10000000; i++ {
        ps := make([]Pair, 0, len(robots))
        for _, r := range robots {
            ps = append(ps, after_iterations(i, r, m, n))
        }

        if find_easter_egg(ps, m, n, i) {
            fmt.Scanln()
            print_pos(ps, m, n)
            fmt.Printf("Seconds passed: %d\n\n\n", i)
        }

    }
}



func Day14() {
    d,err := os.ReadFile("inputs/data14.txt")
    if err != nil {
        fmt.Println("Not opening file correctly: ",err)
        os.Exit(2)
    }
    data := string(d)
    re := regexp.MustCompile(`(-?\d+),(-?\d+)`) 
    matches := re.FindAllStringSubmatch(data, -1)
    // fmt.Println(matches)
    eqs := make([]Robot, 0, len(matches)/2)
    for i:=0; i<len(matches);{
        r := Robot{}
        r.point.x = to_int(matches[i][1]) 
        r.point.y = to_int(matches[i][2]) 
        r.vel.x = to_int(matches[i+1][1])
        r.vel.y = to_int(matches[i+1][2])
        eqs = append(eqs, r)
        i+=2
    }

    // fmt.Println(eqs)
    // fmt.Println(d14part1(eqs))
    d14part2_with_detection(eqs)
}
