package days

import (
	"container/heap"
	"fmt"
	"math"
	"os"
)



func dijkstra(matrix [][]rune) ([][]int, []*MazeNode) {
    m,n := len(matrix), len(matrix[0])
    start := findInMatrix(matrix, 'S')
    end := findInMatrix(matrix, 'E')
    
    dist := make([][]int, 0, m)
    for i:=0; i<m; i++ {
        t := make([]int, 0, n) 
        for j:=0; j<n; j++ {
            t = append(t, math.MaxInt)
        }
        dist = append(dist, t)
    }

    dir := []Pair {
        {-1,0},
        {0,1},
        {1,0},
        {0,-1},
    }

    pq := &MazeQueue{}
    heap.Init(pq)
    heap.Push(pq, &MazeNode{coord: start, cost: 0, minFactor: 0})
    dist[start.x][start.y] = 0
    shortestPaths := []*MazeNode{}

    for pq.Len() > 0 {
        top := heap.Pop(pq).(*MazeNode)
        shortestPaths = append(shortestPaths, top)
        // check if reached the end
        if top.coord.x == end.x && top.coord.y == end.y {
            break
        }
        if matrix[top.coord.x][top.coord.y] == '.' {
            matrix[top.coord.x][top.coord.y] = 'O'
        }
        for _, d := range dir {
            newP := Pair{d.x + top.coord.x, d.y + top.coord.y}
            if newP.x < 0 || newP.y < 0 || newP.x >= m || newP.y >=n {
                continue
            }
            if matrix[newP.x][newP.y] == '#' {
                continue
            }
            if dist[newP.x][newP.y] > dist[top.coord.x][top.coord.y] + 1 {
                dist[newP.x][newP.y] = dist[top.coord.x][top.coord.y] + 1
                heap.Push(pq, &MazeNode{coord: newP,
                cost: top.cost + 1,
                minFactor: top.minFactor + 1,})
            }
        }
    }
    // fmt.Println()
    // d15_print_matrix(matrix)
    // fmt.Println("Something happened")
    return dist, shortestPaths
}

func d20Part1(matrix [][]rune) {
    m,n := len(matrix), len(matrix[0])
    dist, nodes := dijkstra(matrix)
    dir := []Pair {
        {-1,0},
        {0,1},
        {1,0},
        {0,-1},
    }

    // Process every MazeNode
    // map[saved time] number of cheats
    mp := make(map[int]int)
    for _, node := range nodes {
        for _, d := range dir {
            newP := Pair{2*d.x + node.coord.x, 2*d.y + node.coord.y}
            if newP.x < 0 || newP.y < 0 || newP.x >= m || newP.y >=n {
                continue
            }
            if matrix[newP.x][newP.y] == '#' {
                continue
            }
            if dist[newP.x][newP.y] - dist[node.coord.x][node.coord.y] > 2 {
                mp[dist[newP.x][newP.y] - dist[node.coord.x][node.coord.y]-2]++
            }
        } 
    }

    for k,v := range mp {
        fmt.Printf("time saved:%d, total cheats:%d\n",k,v)
    }

    count := 0 
    for k,v := range mp {
        if k >= 100 {
            count += v
        }
    }
    fmt.Println("Cheats with at least 100picoseconds saved",count)
}

func Day20() {
    f, _ := os.ReadFile("inputs/data20.txt")
    data := string(f) 
    matrix := d10parse_input(data)
    d15_print_matrix(matrix)
    d20Part1(matrix)
}
