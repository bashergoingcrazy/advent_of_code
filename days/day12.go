package days

import (
	"fmt"
	"os"
	"strings"
	"time"
)
func parse_input(d []string)([][]rune) {
    matrix := make([][]rune,0,len(d))
    for _, row := range d {
        line := make([]rune, 0,len(row))
        for _, r := range row {
            line = append(line, r)
        }
        matrix = append(matrix, line)
    }
    return matrix
}

func print_matrix(matrix [][]rune) {
    for _, row := range matrix {
        fmt.Println(row)
    }
}

func traverse(p Pair, r rune, matrix [][]rune, vis map[Pair]struct{},
x *Polynomial, flood map[*Polynomial][]Pair) (int, int) {
    // Base case
    m, n := len(matrix), len(matrix[0])
    if p.x < 0 || p.x >= m || p.y < 0 || p.y >= n || matrix[p.x][p.y] != r {
        return 0, 1 // area, perimeter
    }

    if _, exists := vis[p]; exists {
        return 0, 0 // Already visited
    }

    vis[p] = struct{}{} // Mark as visited
    flood[x] = append(flood[x], p)

    // Directions: up, right, down, left
    directions := []Pair{
        {p.x - 1, p.y},
        {p.x, p.y + 1},
        {p.x + 1, p.y},
        {p.x, p.y - 1},
    }

    area, perimeter := 1, 0
    for _, dir := range directions {
        a, per := traverse(dir, r, matrix, vis,x , flood)
        area += a
        perimeter += per
    }
    return area, perimeter
}

type Polynomial struct {
    area, perimeter, side int
}

func initialize_Polynomial() *Polynomial {
    return &Polynomial{-1,-1,-1}
}

func flooding_polynomial(matrix [][]rune) map[*Polynomial][]Pair{
    m,n := len(matrix), len(matrix[0])
    visited := make(map[Pair]struct{})
    
    flood := make(map[*Polynomial][]Pair)


    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if _,exists := visited[Pair{i,j}]; !exists {
                // traverse the current region
                // calculate area and perimeter
                p1 := initialize_Polynomial()
                area, perimeter := traverse(Pair{i,j}, matrix[i][j], matrix, visited, p1, flood)
                p1.area = area
                p1.perimeter = perimeter
            }
        }
    }
    return flood
}

func is_in(p Pair, set map[Pair]struct{}) bool {
    _, exists := set[p]
    return exists
}

const (
    TOPLEFT = iota
    LEFT = iota
    DOWNLEFT = iota
    DOWN = iota
    DOWNRIGHT = iota
    RIGHT = iota
    TOPRIGHT = iota
    TOP = iota
)

func count_corner(i Pair, set map[Pair]struct{}) int {
    // chcek for outer corners
    d := []Pair{
        {i.x-1, i.y-1},
        {i.x, i.y-1},
        {i.x+1, i.y-1},
        {i.x+1, i.y},
        {i.x+1, i.y+1},
        {i.x, i.y+1},
        {i.x-1, i.y+1},
        {i.x-1, i.y},
    }
    res := 0

    
    // check outer corners
    // for top left
    if !is_in(d[TOPLEFT], set) && !is_in(d[LEFT], set) && !is_in(d[TOP],set) {
        res ++
    }
    // for top right
    if !is_in(d[TOPRIGHT], set) && !is_in(d[RIGHT], set) && !is_in(d[TOP],set) {
        res ++
    }
    // for down left
    if !is_in(d[DOWNLEFT], set) && !is_in(d[LEFT], set) && !is_in(d[DOWN],set) {
        res ++
    }
    // for down right
    if !is_in(d[DOWNRIGHT], set) && !is_in(d[DOWN], set) && !is_in(d[RIGHT],set) {
        res ++
    }

    // for inner corners
    if !is_in(d[TOPLEFT], set) && is_in(d[LEFT], set) && is_in(d[TOP],set) {
        res ++
    }
    // for top right
    if !is_in(d[TOPRIGHT], set) && is_in(d[RIGHT], set) && is_in(d[TOP],set) {
        res ++
    }
    // for down left
    if !is_in(d[DOWNLEFT], set) && is_in(d[LEFT], set) && is_in(d[DOWN],set) {
        res ++
    }
    // for down right
    if !is_in(d[DOWNRIGHT], set) && is_in(d[DOWN], set) && is_in(d[RIGHT],set) {
        res ++
    }
    
    return res
}

func fill_sides(flood map[*Polynomial][]Pair) {
    for poly, pairs := range flood {
        set := make(map[Pair]struct{})
        for _, p := range pairs {
            set[p] = struct{}{}
        }
        
        poly.side = 0
        for _, p:= range pairs {
            poly.side += count_corner(p, set)
        }
    }
}

func solve(flood map[*Polynomial][]Pair) (int, int){
    part1, part2 := 0, 0
    for poly := range flood {
        part1 += poly.area*poly.perimeter
        part2 += poly.area*poly.side
    }
    return part1, part2
}


func Day12() {
    t := time.Now()
    f ,err := os.ReadFile("inputs/data12.txt") 
    if err != nil {
        fmt.Println("exx while reading input ",err)
        return
    }
    d := strings.Split(strings.TrimSpace(string(f)),"\n")

    matrix := parse_input(d)
    

    // print_matrix(matrix)
    flood := flooding_polynomial(matrix)
    fill_sides(flood)
    x,y := solve(flood)
    fmt.Printf("Part1: %d  Part2: %d\n",x,y)
    t2 := time.Now()
    elapsed := t2.Sub(t)
    fmt.Printf("Time taken: %v\n",elapsed)
}
