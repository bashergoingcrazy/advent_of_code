package days

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Block struct {
    isEmpty bool
    id int
    fileSize int
}

func d9continuous_block(data string) []Block {
    res := []Block{}
    isFile := true
    id := 0
    for _, r := range data {
       runeLength := int(r-'0')
        for i:=0; i<runeLength; i++ {
            if isFile {
                res = append(res, Block{false,id,runeLength})
            } else {
                res = append(res, Block{true, -1, runeLength})
            }
        } 
        if isFile {
            id++
        }
        isFile = !isFile
    } 
    return res
}

// 90907254371
func checksum(data []Block) int {
    total := 0
    for i, b := range data {
        if b.isEmpty {continue}
        total += i * b.id
    }
    return total
}

func d9Compact(data []Block) []Block {
    n := len(data)
    
    // Identify files and their properties
    files := []struct {
        id       int
        start    int
        fileSize int
    }{}
    
    // Collect file metadata
    for i := 0; i < n; {
        if !data[i].isEmpty {
            id := data[i].id
            fileSize := data[i].fileSize
            files = append(files, struct {
                id       int
                start    int
                fileSize int
            }{id, i, fileSize})
            i += fileSize
        } else {
            i++
        }
    }

    // Sort files by ID in descending order
    sort.Slice(files, func(i, j int) bool {
        return files[i].id > files[j].id
    })

    // Move each file
    for _, file := range files {
        start, size := file.start, file.fileSize

        // Find the leftmost contiguous free space
        for i := 0; i < start; i++ {
            validSpan := true
            for j := 0; j < size; j++ {
                if i+j >= start || !data[i+j].isEmpty {
                    validSpan = false
                    break
                }
            }

            if validSpan {
                // Move the file
                for j := 0; j < size; j++ {
                    data[i+j] = data[start+j]
                    data[start+j] = Block{true, -1, 0}
                }
                break
            }
        }
    }

    return data
}


func Day9() {
    d,err := os.ReadFile("inputs/data9.txt")
    if err != nil {
        fmt.Println("os can't read thef file: ", err)
        return
    }

    data := strings.TrimSpace(string(d))
    fmt.Println(len(data))

    r := d9continuous_block(data)
    fmt.Println(r)
    fmt.Println()
    cr := d9Compact(r)
    // fmt.Println(cr)
    res := checksum(cr)
    fmt.Println(res)
    // fmt.Println(data)
    // fmt.Println(ids)
    // fmt.Println(c)
    // fmt.Println(res)


}
