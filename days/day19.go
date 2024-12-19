package days

import (
	"fmt"
	"os"
	"strings"
	"time"
)
func d19ParseInput(f []byte) (map[string]struct{}, []string) {
    data :=strings.Split(strings.TrimSpace(string(f)), "\n\n") 
    dictionary := data[0]
    wordSpace := data[1]

    dict := make(map[string]struct{})

    maxLengthInDict = 0
    ds := strings.Split(dictionary, ",")
    for _, str := range ds {
        str := strings.TrimSpace(str)
        maxLengthInDict = max(len(str), maxLengthInDict)
        dict[str] = struct{}{}
    }
    qwords := strings.Split(wordSpace, "\n")
    return dict, qwords 
}

var maxLengthInDict int
// recursive function to check if the design is possible
func isDesignPossibleP1(design string, index int, 
dict map[string]struct{}, memo map[int]bool) bool {
    // base case
    if index >= len(design) {
        return true
    }

    if result, exists := memo[index]; exists {
        return result
    }

    for i := index+1; i <= index + maxLengthInDict && i <= len(design); i++ {
        possibleMatch := design[index:i]
        // fmt.Printf("Checking design[%d:%d] = %q\n", index, i, possibleMatch)
        if _, exists := dict[possibleMatch]; exists {
            if isDesignPossibleP1(design, i, dict, memo) { 
                memo[index] = true
                // fmt.Printf("Match found for design[%d:%d] = %q\n", index, i, possibleMatch)
                return true
            }
        }
    }
    memo[index] = false
    return false
}

func d19Part1(dict map[string]struct{}, designSpace []string) int {
    count := 0
    for _, design := range designSpace {
        memo := make(map[int]bool)
        if isDesignPossibleP1(design, 0, dict, memo) {
            count++
        }
    }
    return count
}


func isDesignPossibleP2(design string, index int, 
dict map[string]struct{}, memo map[int]int) int {
    // base case
    if index >= len(design) {
        return 1
    }

    if result, exists := memo[index]; exists {
        return result
    }


    total := 0
    for i := index+1; i <= index + maxLengthInDict && i <= len(design); i++ {
        possibleMatch := design[index:i]
        // fmt.Printf("Checking design[%d:%d] = %q\n", index, i, possibleMatch)
        if _, exists := dict[possibleMatch]; exists {
            result := isDesignPossibleP2(design, i, dict, memo)
            if result > 0 { 
                total += result
                // fmt.Printf("Match found for design[%d:%d] = %q\n", index, i, possibleMatch)
            }
        }
    }

    memo[index] = total
    return total
}

func d19Part2(dict map[string]struct{}, designSpace []string) int {
    count := 0
    for _, design := range designSpace {
        memo := make(map[int]int)
        count += isDesignPossibleP2(design, 0, dict, memo)
    }
    return count
}

func Day19() {
    ts := time.Now()
    debug := 0
    f,_ := os.ReadFile("inputs/data19.txt")
    if debug == 1 {
        f,_ = os.ReadFile("inputs/data19dummy.txt")
    }
    dict, designSpace := d19ParseInput(f)
    // fmt.Println(dict, designSpace) 

    fmt.Println(d19Part1(dict, designSpace))
    fmt.Println(d19Part2(dict, designSpace))
    tp := time.Now()
    elapsed := tp.Sub(ts)
    fmt.Println("Elapsed:",elapsed)
}
