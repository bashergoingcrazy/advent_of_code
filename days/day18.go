package days

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

func bfs(matrix [][]rune) int {
    m, n := len(matrix), len(matrix[0])
    start := Pair{0, 0}
    goal := Pair{m - 1, n - 1}

    q := &MazeQueue{}
    heap.Init(q)
    heap.Push(q, &MazeNode{
        coord: start,
        cost:  0,
        minFactor: heuristic(start, goal), // Start with the heuristic
    })

    dir := [4]Pair{
        {-1, 0}, {1, 0}, {0, 1}, {0, -1},
    }

    visited := make(map[Pair]bool) // Visited nodes

    for len(*q) > 0 {
        top := heap.Pop(q).(*MazeNode)

        // Check if we reached the goal
        if top.coord == goal {
            return top.cost
        }

        // Skip if this node was already processed
        if visited[top.coord] {
            continue
        }
        visited[top.coord] = true

        for _, d := range dir {
            nextD := Pair{top.coord.x + d.x, top.coord.y + d.y}

            // Check bounds
            if nextD.x < 0 || nextD.y < 0 || nextD.x >= m || nextD.y >= n {
                continue
            }

            // Check obstacles
            if matrix[nextD.x][nextD.y] == '#' {
                continue
            }

            // Skip if visited
            if visited[nextD] {
                continue
            }

            heap.Push(q, &MazeNode{
                coord:     nextD,
                cost:      top.cost + 1,
                minFactor: top.cost + 1 + heuristic(nextD, goal),
            })
        }
    }

    return -1 // No path found
}

func firstPreventingByte(matrix [][]rune, ds []string) {
    for i:=1024; i<len(ds); i++ {
        k := strings.Split(ds[i], ",")
        x := parseInt(k[1])
        y := parseInt(k[0])
        matrix[x][y] = '#'
        if bfs(matrix) == -1 {
            fmt.Println(ds[i])
            break
        }
    }
}

func Day18() {
    f, err := os.ReadFile("inputs/data18.txt")
    if err != nil {
        os.Exit(2)
    }
    data := strings.TrimSpace(string(f))
    // fmt.Println(data)
    m,n := 71,71
    matrix := make([][]rune,0,m)
    for i:=0; i<m; i++ {
        row := make([]rune, 0, n) 
        for j:=0; j<n; j++ {
            row = append(row, '.')
        }
        matrix = append(matrix, row)
    }

    ds := strings.Split(data, "\n")
    for i, row := range ds {
        if i == 1024 {break}
        k := strings.Split(row, ",") 
        x := parseInt(k[1])
        y := parseInt(k[0])
        matrix[x][y]='#'
    }
    
    // d15_print_matrix(matrix)
    fmt.Println(bfs(matrix))
    firstPreventingByte(matrix, ds)
}
