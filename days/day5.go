package days

import (
	"fmt"
	"os"
	"strconv"
    "strings"
)

func createGraph(rules []string) map[string][]string {
	graph := make(map[string][]string)

	for _, line := range rules {
		nodes := strings.Split(line, "|")
		if len(nodes) != 2 {
			fmt.Println("Error in format while creating graph n1|n2 ",nodes)
			return graph
		}
		from, to := nodes[0], nodes[1]
		graph[from] = append(graph[from], to)
		if _, exists := graph[to]; !exists {
			graph[to]=[]string{}
		}
	}
	return graph
}

func createSubGraph(graph map[string][]string, nodes []string) map[string][]string {
	subG := make(map[string][]string)
	set := make(map[string]bool)

	for _, node := range nodes {
		set[node] = true
	}

	for _, node := range nodes {
		for _, neighbor := range graph[node] {
			if _, exists := set[neighbor]; exists {
				subG[node] = append(subG[node], neighbor)
			}
		}
		if _, exists := subG[node]; !exists {
			subG[node] = []string{}
		}
	}
	return subG
}

func topologicalSort(graph map[string][]string) ([] string, bool) {
	inDegree := make(map[string]int)

	for node := range graph {
		inDegree[node] = 0
	}

	for _, neighbors := range graph {
		for _, nbr := range neighbors {
			inDegree[nbr]++	
		}
	}

	var queue []string

	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	var sortedOrder []string
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		sortedOrder = append(sortedOrder, curr)

		for _, nbr := range graph[curr] {
			inDegree[nbr]--
			if inDegree[nbr] == 0 {
				queue = append(queue, nbr)
			}
		}
	}

	if len(sortedOrder) != len(graph){
		fmt.Println("Graph has a cycle !! topological sort not possible ")
		return nil,false
	}
	return sortedOrder, true
}

func parseInput(data string) ([]string, []string) {
	sections := strings.Split(strings.TrimSpace(data), "\n")
	rulesSection, quesSection := []string{},[]string{}

	for _, line := range sections {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		} else if strings.Contains(line, "|") {
			rulesSection = append(rulesSection, line)
		} else {
			quesSection = append(quesSection, line)
		}
	}
	return rulesSection,quesSection
}

func isEqualTo(sl1, sl2 []string) bool {
	if len (sl1) != len (sl2) {
		return false
	}

	for i:=0; i<len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

func midValueOf(sl1 [] string) int {
	index := len(sl1)/2
	value, err := strconv.Atoi(sl1[index])
	if err != nil {
		fmt.Println("Not an integer value at the middle index something's gone horribly wrong")
		return 0
	}
	return value
}

func Day5() {
	fileData, err := os.ReadFile("inputs/data5real.txt")
	if err!=nil {
		fmt.Println("Error occured while reading file : ",err)
		return
	}
	data := string(fileData)
	rules, ques := parseInput(data)

	graph := createGraph(rules)

	middleSum := 0
	for _, q := range ques {
		qs := strings.Split(q, ",")
		subG := createSubGraph(graph, qs)
		actualOrder, isPossible := topologicalSort(subG)
		if !isPossible {
			continue
		}

		fmt.Println(actualOrder,qs)

		if isEqualTo(actualOrder, qs) {
			middleSum += midValueOf(actualOrder)	
		}
	}
	fmt.Println("Output : ",middleSum)
}
