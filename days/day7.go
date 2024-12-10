package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


func parseInt(data string) int{
    d,err := strconv.Atoi(data)
    if err != nil {
        fmt.Println("Error converting to int: ",err)
    }
    return d
}

func parseIntSlice(data []string) []int {
    res := make([]int,0,len(data))
    for _, val := range data {
        res = append(res, parseInt(val))
    }
    return res
}

func d7_parseInput(data string)([]int, [][]int) {
    dlist := strings.Split(strings.TrimSpace(data), "\n")

    results := make([]int, 0, len(dlist))
    queries := make([][]int, 0, len(dlist))

    
    for _, lines := range dlist {
        sections := strings.Split(lines , ":")

        querry := strings.Split(strings.TrimSpace(sections[1]), " ")
        

        
        results = append(results, parseInt(sections[0]))
        queries = append(queries, parseIntSlice(querry))
    }
    return results,queries
}

func concat(x,y int) int {
    power := 1
    for y >= power {
        power *= 10
    }
    return x*power + y
}

func isValid(query []int, index,currentVal, result int, comb bool) bool {
    // prune
    if currentVal > result {
        return false
    }

    // base case
    if index == len(query) {
        // fmt.Println(currentVal)
        if result == currentVal {
            return true
        }
        return false
    }

    if comb {
        return isValid(query, index+1, currentVal+query[index], result, comb) ||
        isValid(query, index+1, currentVal*query[index], result, comb) ||
        isValid(query, index+1, concat(currentVal,query[index]), result, comb)
    }

    return isValid(query, index+1, currentVal+query[index], result, comb) ||
        isValid(query, index+1, currentVal*query[index], result, comb) 
}

func d7_part1(result []int, queries [][]int, comb bool) int {
    n := len(result)

    total := 0
    
    for i:=0; i<n; i++ {
        if len(queries[i]) == 0 {
            continue
        }

        currentVal := queries[i][0] 
        // fmt.Printf("%d:%v\n",result[i], queries[i])
        if isValid(queries[i],1,currentVal,result[i], comb) {
            total += result[i]
        }
    }
    return total
}

func Day7() {
    start_time := time.Now()
    d, err := os.ReadFile("inputs/data7.txt")
    if err != nil {
        fmt.Println("Error while opening the file: ",err)
        return
    }
    data := string(d)
    results, queries := d7_parseInput(data)
    fmt.Println(d7_part1(results, queries, false))
    fmt.Println(d7_part1(results, queries, true))

    t := time.Now()
    elapsed_time := t.Sub(start_time)
    fmt.Println(elapsed_time)

}
