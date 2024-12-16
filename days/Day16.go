package days

import (
	"container/heap"
	"fmt"
	"os"
)


type MazeNode struct {
    coord Pair
    cost int
    direction Pair
    minFactor int
}
type MazeQueue []*MazeNode
// -------------- functions for heap interface ---------------
func (pq MazeQueue) Len()int {return len(pq)}
func (pq MazeQueue) Swap(i,j int){pq[i],pq[j] = pq[j], pq[i]}
func (pq MazeQueue) Less(i, j int)bool {
    return pq[i].minFactor < pq[j].minFactor
}
func (pq *MazeQueue) Push(x any) {
    node := x.(*MazeNode)
    *pq = append(*pq, node)
}
func (pq *MazeQueue) Pop() any {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    *pq = old[0:n-1]
    return item
}
// -------------------end for heap interface -----------------


func findInMatrix(matrix [][]rune, find rune)Pair {
    z := Pair{-1,-1}
    for i, row := range matrix {
        for j, r := range row {
            if r == find {
                z.x = i 
                z.y = j
                break
            }
        }
        if z.x != -1 {break}
    }
    return z
}
const h_mul int = 0 
func heuristic(currState, goalState Pair) int {
    xval := abs(goalState.x - currState.x)
    yval := abs(goalState.y - currState.y)
    return (xval + yval)*h_mul
}

func bfsa(matrix [][]rune) int {
    dir := [4]Pair {
        {0, 1},  // Right
        {-1, 0}, // Top
        {0, -1}, //Left
        {1, 0}, // Bottom
    }

    startState := findInMatrix(matrix, 'S')
    goalState := findInMatrix(matrix, 'E')
    pq := &MazeQueue{}
    heap.Init(pq)
    heap.Push(pq, &MazeNode{
        coord: startState,
        direction: Pair{0,1},
        cost: 0,
        minFactor: heuristic(startState,goalState),
    })

    fmt.Println(startState,goalState)
    visited := make(map[Pair]struct{})

    for len(*pq) > 0 {
        currState := heap.Pop(pq).(*MazeNode)

        if _, exists := visited[currState.coord]; exists {
            continue
        }
        visited[currState.coord] = struct{}{}

        if currState.coord.x == goalState.x &&
        currState.coord.y == goalState.y {
            return currState.cost
        }

        
        // matrix[currState.coord.x][currState.coord.y] = 'O'
        //
        // fmt.Println()
        // d15_print_matrix(matrix)
        // fmt.Println(*currState)
        // fmt.Scanln()

        // iterate over all directions and push for valid states
        for _, d := range dir {
            
            
            nextCord := Pair{d.x + currState.coord.x, d.y + currState.coord.y}
            if matrix[nextCord.x][nextCord.y] == '#' { continue }

            extraCost := 0
            if d.x != currState.direction.x ||
            d.y != currState.direction.y {
               extraCost += 1000 
            }

            nextNode := &MazeNode{
                coord: nextCord,
                direction: d,
                cost: 1 + extraCost + currState.cost,
            }
            nextNode.minFactor = nextNode.cost + heuristic(nextCord, goalState)

            heap.Push(pq, nextNode)
        }
    }
    
    return 0
}




func Day16(){
    f, err := os.ReadFile("inputs/data16.txt") 
    if err != nil {
        fmt.Println("Error occured while reading file: ",err)
        os.Exit(2)
    }
    data := string(f)
    fmt.Println(data)
    matrix := d10parse_input(data)
    d15_print_matrix(matrix)
    cost := bfsa(matrix)
    fmt.Println(cost)
}
// 102508

