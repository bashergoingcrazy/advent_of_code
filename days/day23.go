package days

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Graph [676][676]bool

func indexOf(cs string) int {
    xOffset := int(cs[0]-'a')
    yOffset := int(cs[1]-'a')
    return xOffset*26 + yOffset
}
func btStr(x int) string {
    barr := make([]byte, 2) 
    barr[0] = 'a' + byte((x/26))
    barr[1] = 'a' + byte((x%26))
    return string(barr)
}

func create_graph(f []byte) Graph {
    graph := Graph{}
    for i:=0; i<676; i++ {
        for j:=0; j<676; j++ {
            graph[i][j] = false
        }
    }

    data := string(f) 
    d := strings.Split(strings.TrimSpace(data), "\n")

    for _, ab := range d {
        x := indexOf(ab[0:2])
        y := indexOf(ab[3:])
        graph[x][y] = true
        graph[y][x] = true
    }
    return graph
}

func neighbors_of(graph Graph, f int) []int {
    res := []int{}
    for i:=0; i<676; i++ {
        if graph[f][i] {res = append(res, i)}
    }
    return res
}

func d23part1(graph Graph) {
    set := make(map[[3]int]struct{})
    
    // Indices corresponding to 'ta' to 'tz' (adjust based on `indexOf` implementation)
    for index := indexOf("ta"); index <= indexOf("tz"); index++ {
        nbrs := neighbors_of(graph, index) // Neighbors of current node

        // Check all pairs of neighbors to find triangles
        for i := 0; i < len(nbrs); i++ {
            for j := i + 1; j < len(nbrs); j++ {
                if graph[nbrs[i]][nbrs[j]] { // Check if nbrs[i] and nbrs[j] are connected
                    // Create a sorted triangle to store in the set
                    temp := [3]int{index, nbrs[i], nbrs[j]}
                    sort.Ints(temp[:]) // Sort the triangle nodes
                    set[temp] = struct{}{}
                }
            }
        }
    }

    // Print the results
    for triad := range set {
        fmt.Println(btStr(triad[0]), btStr(triad[1]), btStr(triad[2]))
    }
    fmt.Println("Total triangles:", len(set))
}

func bronKerbosch(graph Graph, r, p, x []int, cliques *[][]int) {
    if len(p) == 0 && len(x) == 0 {
        clique := make([]int, len(r))
        copy(clique, r)
        *cliques = append(*cliques, clique)
        return
    }

    for i := 0; i < len(p); i++ {
        v := p[i]
        newR := append(r, v)
        newP := intersect(p, neighbors_of(graph, v))
        newX := intersect(x, neighbors_of(graph, v))
        bronKerbosch(graph, newR, newP, newX, cliques)
        p = remove(p, v)
        x = append(x, v)
    }
}

func intersect(a, b []int) []int {
    set := make(map[int]bool)
    for _, v := range b {
        set[v] = true
    }
    res := []int{}
    for _, v := range a {
        if set[v] {
            res = append(res, v)
        }
    }
    return res
}

func remove(slice []int, elem int) []int {
    res := []int{}
    for _, v := range slice {
        if v != elem {
            res = append(res, v)
        }
    }
    return res
}

func findMaximalCliques(graph Graph) [][]int {
    nodes := []int{}
    for i := 0; i < 676; i++ {
        nodes = append(nodes, i)
    }

    cliques := [][]int{}
    bronKerbosch(graph, []int{}, nodes, []int{}, &cliques)
    return cliques
}

func largestClique(graph Graph) []int {
    cliques := findMaximalCliques(graph)
    maxClique := []int{}
    for _, clique := range cliques {
        if len(clique) > len(maxClique) {
            maxClique = clique
        }
    }
    return maxClique
}

func formatPassword(clique []int) string {
    names := []string{}
    for _, index := range clique {
        names = append(names, btStr(index))
    }
    sort.Strings(names)
    return strings.Join(names, ",")
}

func d23part2(graph Graph) {
    maxClique := largestClique(graph)
    password := formatPassword(maxClique)
    fmt.Println("LAN party password:", password)
}

func Day23() {
    debug := 0
    f := []byte{}
    if debug == 0 {
        f, _ = os.ReadFile("inputs/data23.txt")
    } else {
        f, _ = os.ReadFile("inputs/data23dummy.txt")
    }

    graph := create_graph(f)
    d23part1(graph)
    d23part2(graph)
}
