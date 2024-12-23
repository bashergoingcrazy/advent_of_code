package days

import (
	"fmt"
	"os"
	"strings"
)

const MODM = 16777216


func secret_iter(secret int) int {
    secret = (((secret*64)^secret) % MODM)
    secret = (((secret/32)^secret) % MODM)
    secret = (((secret*2048)^secret) % MODM)
    return secret
}

func d22part1(data []int) int {
    count := 0
    for _, val := range data {
        x := val
        for i:=0; i<2000; i++ {
            x = secret_iter(x) 
        }
        fmt.Println(val,":",x)
        count += x
    }
    return count
}

func fill_collection(bananaList []int, startVal int, collection map[[4]int]int) {
    set := make(map[[4]int]struct{})
    temp := [4]int{bananaList[0], bananaList[1],bananaList[2],bananaList[3]}
    val := 0
    val += startVal + temp[0] + temp[1] + temp[2] + temp[3]
    collection[temp] += val
    set[temp] = struct{}{}

    for i:=4; i<2000; i++ {
        temp := [4]int{bananaList[i-3], bananaList[i-2],bananaList[i-1],
        bananaList[i]}
        val += temp[3]

        if _, exists := set[temp]; !exists {
            collection[temp] += val
            set[temp] = struct{}{}
        }
    }
}

func d22part2(data []int) int {
    collection := make(map[[4]int]int)
    for _, val := range data {
        x := val 
        bananaList := make([]int,0,2000)
        for i:=0; i<2000; i++ {
            t := secret_iter(x)
            bananaVal := (t%10) - (x%10) 
            bananaList = append(bananaList, bananaVal)
            x = t
        }
        fill_collection(bananaList, val%10, collection)
    }

    ans := -1
    for _,v := range collection {
        if v > ans {
            ans = v
        }
    }
    return ans
}


func Day22() {
    debug := 0
    f := []byte{}
    if debug == 1 {
        f,_ = os.ReadFile("inputs/data22dummy.txt")
    } else {
        f,_ = os.ReadFile("inputs/data22.txt")
    }
    data := strings.Split(strings.TrimSpace(string(f)), "\n")
    di := parseIntSlice(data)

    fmt.Println(d22part1(di))
    fmt.Println(d22part2(di))
}
