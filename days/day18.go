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


 
// ---------------------- Dijsoint set data structure -------------------

type disjointSet struct {
    parent, rank []int
}

func init_dsj(n int) *disjointSet {
    parent, rank := make([]int, n), make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    return &disjointSet{parent: parent, rank: rank}
}

func (dsj *disjointSet) find(x int) int {
    if dsj.parent[x] == x { return x }
    dsj.parent[x] = dsj.find(dsj.parent[x])
    return dsj.parent[x]
}

func (dsj *disjointSet) union(x,y int) {
    xrep, yrep := dsj.find(x), dsj.find(y)
    if xrep == yrep { return }
    if dsj.rank[x] < dsj.rank[y] {
        dsj.parent[xrep] = yrep
    } else if dsj.rank[y] < dsj.rank[x] {
        dsj.parent[yrep] = xrep
    } else {
        dsj.parent[yrep] = xrep
        dsj.rank[xrep]++
    }
}
// ------------------- End of Disjoint Set data structure ----------------

func d18part2(matrix [][]rune, fallingPoints []string) {
    m,n := len(matrix), len(matrix[0])
    djs := init_dsj(m*n)
    
    // Initialize unions for walls without connecting (0,0) and (m-1,n-1)
    for i := 1; i < m-1; i++ {
        djs.union(i*n, (i+1)*n)                     // Connect adjacent left wall elements
        djs.union((i-1)*n + n-1, i*n + n-1)         // Connect adjacent right wall elements
    }
    for j := 1; j < n-1; j++ {
        djs.union(j, j+1)                           // Connect adjacent top wall elements
        djs.union((m-1)*n+j-1, (m-1)*n+j)          // Connect adjacent bottom wall elements
    }
    // this is done is such a way that 
    // 0,0 and m-1,n-1 are not added anywhere

    // Process all the falling points
    dir := [4]Pair{
        {-1, 0}, {1, 0}, {0, 1}, {0, -1},
    }

    for i, points := range fallingPoints {
        sl := strings.Split(points, ",")
        x := parseInt(strings.TrimSpace(sl[1]))
        y := parseInt(strings.TrimSpace(sl[0]))
        for _, d := range dir {
            nextP := Pair{d.x + x, d.y + y}
            if nextP.x < 0 || nextP.x >= m || nextP.y < 0 || nextP.y >= n {
                continue
            }
            djs.union(x*n + y, nextP.x*n + nextP.y)
        }
        // check if the conditions for closing has been met
        xrep,yrep := djs.find(n-1), djs.find(n)
        if xrep == yrep {
            fmt.Println(fallingPoints[i])
            break
        }
    }
}

func Day18() {
    f, err := os.ReadFile("inputs/data18dummy.txt")
    if err != nil {
        os.Exit(2)
    }
    data := strings.TrimSpace(string(f))
    // fmt.Println(data)
    m,n := 7,7
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
        if i == 12 {break}
        k := strings.Split(row, ",") 
        x := parseInt(k[1])
        y := parseInt(k[0])
        matrix[x][y]='#'
    }
    
    d15_print_matrix(matrix)
    d18part2(matrix, ds)
    // fmt.Println(bfs(matrix))
    // firstPreventingByte(matrix, ds)
}
