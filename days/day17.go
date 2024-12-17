package days

import (
	"fmt"
	"math"
	"strings"
)

var A,B,C,IP int

func combo(operand int) int {
    if operand <= 3 {
        return operand
    } 
    switch operand {
    case 4:
        return A
    case 5:
        return B
    case 6: 
        return C
    }
    panic("Operand out of index")
}

func adv(operand int) {
    operand = combo(operand)
    floatA, floatOp := float64(A), float64(operand)
    A = int(floatA/math.Pow(2,floatOp))
}

func bxl(operand int) {
    ans := B ^ operand
    B = ans 
}

func bst(operand int) {
    operand = combo(operand)
    B = operand % 8
}

func jnz(operand int) {
    if A != 0 {
        IP = operand - 2
    }
}

func bxc() {
    B = B ^ C
}

func out2(operand int)int {
    operand = combo(operand) 
    out := operand %8
    fmt.Printf("%d,",out)
    return out 
}
func out(operand int)int {
    operand = combo(operand) 
    out := operand %8
    // fmt.Printf("%d,",out)
    return out 
}

func bdv(operand int) {
    operand = combo(operand)
    floatA, floatOp := float64(A), float64(operand)
    B = int(floatA/math.Pow(2,floatOp))
}

func cdv(operand int) {
    operand = combo(operand)
    floatA, floatOp := float64(A), float64(operand)
    C = int(floatA/math.Pow(2,floatOp))
}

func pState(){
    fmt.Printf("A:%d, B:%d, C:%d, IP:%d",A,B,C,IP)
    fmt.Scanln()
}

func d17_part1(instructions []int) {
    for IP < len(instructions) {
        // pState()
        opcode := instructions[IP]
        operand := instructions[IP+1]
        switch opcode {
        case 0:
            adv(operand)
        case 1:
            bxl(operand)
        case 2:
            bst(operand)
        case 3:
            jnz(operand)
        case 4:
            bxc()
        case 5:
            out2(operand)
        case 6:
            bdv(operand)
        case 7:
            cdv(operand)
        default:
            panic("Something terrible happened")
        }
        IP += 2
    }
}

func compareIntSlice(s1,s2[]int) bool{
    if len(s1) != len(s2) {
        return false
    }
    for i:=0; i<len(s1); i++ {
        if s1[i]!=s2[i] {
            return false
        }
    }
    return true
}

func debugginStuf (instructions []int) {
    for i:=164540892147389; i<164540892147390; i++ {
        A = i
        B,C,IP = 0,0,0
        fmt.Printf("\ni:%d ",i)
        for IP < len(instructions) {
            // pState()
            opcode := instructions[IP]
            operand := instructions[IP+1]
            switch opcode {
            case 0:
                adv(operand)
            case 1:
                bxl(operand)
            case 2:
                bst(operand)
            case 3:
                jnz(operand)
            case 4:
                bxc()
            case 5:
                out2(operand)
            case 6:
                bdv(operand)
            case 7:
                cdv(operand)
            default:
                panic("Something terrible happened")
            }
            IP += 2
        }
    }
}

func d17_part2(instructions []int) {
    // don't know the value of the regist
    val := 0
    for i:=len(instructions)-1; i>=0; i-- {
        val = val << 3
        for {
            A = val
            B,C,IP = 0,0,0
            found := false
            for IP < len(instructions) {
                opcode := instructions[IP]
                operand := instructions[IP+1]
                k := -1
                switch opcode {
                case 0:
                    adv(operand)
                case 1:
                    bxl(operand)
                case 2:
                    bst(operand)
                case 3:
                    jnz(operand)
                case 4:
                    bxc()
                case 5:
                    k = out(operand) 
                    fmt.Println(k)
                case 6:
                    bdv(operand)
                case 7:
                    cdv(operand)
                default:
                    panic("Something terrible happened")
                }
                IP += 2
                if k != -1 {
                    if k == instructions[i] {
                        found = true
                    }
                    break
                }
            }
            if found {
                fmt.Printf("ins:%d val:%d ",instructions[i],val)
                pState()
                break
            }
            val++
        }

    }
    fmt.Println("A",val)

}

func Day17() {
    fmt.Println("Hello world")
    IP = 0
    debug := 1 
    var data string
    // parse your inputs
    
    dc := strings.Split(data, ",")
    di := parseIntSlice(dc)
    if debug ==1{
        debugginStuf(di)
    }
    fmt.Println(di)
    d17_part2(di)
    // d17_part1(di)
}

