package days

import (
	"fmt"
	"os"
	"strings"
)
func parse_input(d []string)([][]rune, map[rune][]Pair) {
    matrix := make([][]rune,0,len(d))
    mp := make(map[rune][]Pair)
    for i, row := range d {
        line := make([]rune, 0,len(row))
        for j, r := range row {
            line = append(line, r)
            mp[r] = append(mp[r], Pair{i,j})
        }
        matrix = append(matrix, line)
    }
    return matrix,mp
}

func d12_print(matrix [][]rune) {
    for _, row := range matrix {
        fmt.Println(row)
    }
}

func total_perimeter(p Pair, r rune, matrix [][]rune ,vis map[Pair]struct{}) int {
    m,n := len(matrix), len(matrix[0])

    // base case
    if p.x < 0 || p.x >= m || p.y < 0 || p.y >=n || matrix[p.x][p.y] != r {
        return 1
    }
    
    vis[p] = struct{}{}

    dir := []Pair{
        {p.x-1, p.y},
        {p.x, p.y+1},
        {p.x+1, p.y},
        {p.x, p.y-1},
    }

    count := 0
    // recursive case
    for _, d := range dir {
        if _, exists := vis[d]; !exists{
            count += total_perimeter(d, r, matrix, vis)
        }
    }
    return count
}

func Day12() {
    f ,err := os.ReadFile("inputs/data12.txt") 
    if err != nil {
        fmt.Println("exx while reading input ",err)
        return
    }
    d := strings.Split(strings.TrimSpace(string(f)),"\n")

    matrix, mp := parse_input(d)
    

    ans := 0
    for r, region := range mp {
        // area := len(region)
        visited := make(map[Pair]struct{})
        prevVisitedSize := 0
        for _, block := range region {
            if _,exists := visited[block]; !exists {
                perimeter := total_perimeter(block, r, matrix ,visited)
                visitedSize := len(visited)
                ans += perimeter*(visitedSize-prevVisitedSize)
                fmt.Printf("block: %c  area: %d  perimeter: %d\n",r,visitedSize-prevVisitedSize,perimeter)
                prevVisitedSize = visitedSize
            }
        }
    }

    fmt.Println(ans)
    // fmt.Println(mp)
    // d12_print(matrix)
}
