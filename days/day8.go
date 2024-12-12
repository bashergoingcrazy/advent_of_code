package days

import (
	"fmt"
	"os"
	"strings"
)

func d8_parse_input(data string) [][]string{

    sp := strings.Split(strings.TrimSpace(data), "\n")

    matrix := make([][]string, 0, len(sp))
    for _, row := range sp {
        rowD := strings.TrimSpace(row) 
        f := strings.Split(rowD, "")
        matrix = append(matrix, f)
    }
    return matrix
}

func d8_print_matrix(matrix [][]string) {
    for _, row := range matrix {
        fmt.Println(row)
    }
    fmt.Printf("%d,%d\n\n", len(matrix), len(matrix[0]))
}

type Pair struct {
    x, y int
}

const NN int = 12
const M int = 50

func d8_is_valid_pair_dummy(p Pair)bool {
    if p.x >= 0 && p.x < NN && p.y >= 0 && p.y < NN {
        return true
    }
    return false
}

func d8_is_valid_pair(p Pair)bool {
    if p.x >= 0 && p.x < M && p.y >= 0 && p.y < M {
        return true
    }
    return false
}

func d8_anti_points_of(px,py Pair, mp2 map[Pair]bool) {
    diffX := py.x - px.x
    diffY := px.y - py.y

    anti1 := Pair{px.x - diffX, px.y + diffY}
    for d8_is_valid_pair(anti1) {
        mp2[anti1] = true
        anti1 = Pair{anti1.x - diffX, anti1.y + diffY}
    }

    anti2 := Pair{py.x + diffX, py.y - diffY}
    for d8_is_valid_pair(anti2) {
        mp2[anti2] = true
        anti2 = Pair{anti2.x + diffX, anti2.y - diffY}
    }

    
}

func d8_total_anti_points(px Pair, points []Pair, mp2 map[Pair]bool) {
    if len(points) == 0 {return}
    mp2[px] = true
    for _, py := range points {
        mp2[py] = true
        d8_anti_points_of(py, px, mp2)
    }
    return
}

func Day8() {
    d, err := os.ReadFile("inputs/data8.txt")
    if err != nil {
        fmt.Println("error opening the file: ",err)
        return
    }
    data := string(d)
    matrix := d8_parse_input(data)

    d8_print_matrix(matrix)

    mp := make(map[string][]Pair)
    mp2 := make(map[Pair]bool)

    for i, rows := range matrix {
        for j, s := range rows {
            if s == "." {continue}
            p := Pair{i,j} 
            fmt.Printf("%s, %v\n",s,p)
            d8_total_anti_points(p, mp[s],mp2)

            mp[s] = append(mp[s], p)

            // fmt.Scanln()
        }
    }

    fmt.Println(len(mp2))
    // print matrix
    
    // for p := range mp2 {
    //     if matrix[p.x][p.y] == "." {
    //         matrix[p.x][p.y] = "X"
    //     }
    // } 
    // d8_print_matrix(matrix)
}
