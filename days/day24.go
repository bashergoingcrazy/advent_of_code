package days

import (
	"fmt"
	"os"
	"strings"
)

// ---------- Implementation of Wire ----------------------
type Wire struct {
    isDefined bool     // To check if the signal has reached this wire
    state int          // Active is 1 Ground is 0
    fromGate *Gate     // From which gate the wire initially originates probably don't need this
    toGates []*Gate    // To which gates the wire act as an input
}

func NewWire(isDefined bool, preState int) *Wire {
    return &Wire{isDefined:isDefined ,state: preState}
}

func (w *Wire) addFromGate(fromGate *Gate) {
    w.fromGate = fromGate
}

func (w *Wire) appendToGate(toGate *Gate) {
    w.toGates = append(w.toGates, toGate)
}

// ---------- Implementation of Gate ----------------------
type Gate struct {
    typeOf int          // AND=0, OR=1, XOR=2
    input1 *Wire
    input2 *Wire
    output *Wire
}


func d24_parse_input(f []byte) ([]*Gate, map[string]*Wire) {
    strf := strings.Split(strings.TrimSpace(string(f)),"\n\n")

    // Parse all the pre defined wires
    preDefinedWires := strf[0]
    // fmt.Print(preDefinedWires)
    ds := strings.Split(preDefinedWires, "\n")
    allWires := make(map[string]*Wire)
    for _, d := range ds {
        flds := strings.Fields(d)
        wire := flds[0][0:3] 
        state := parseInt(flds[1])
        allWires[wire] = NewWire(true, state)
    }

    // Parse all the gates information
    definedGates := strf[1]
    dsg := strings.Split(definedGates, "\n")
    allGates := make([]*Gate,0,len(dsg))
    for _, d := range dsg {
        flds := strings.Fields(d)
        input1 := flds[0]
        input2 := flds[2]
        gateType := flds[1]
        output := flds[4]


        newGate := &Gate{}
        // correctly initialize the wires
        if _,exists := allWires[input1]; !exists {
            allWires[input1] = NewWire(false, -1)
        }
        if _,exists := allWires[input2]; !exists {
            allWires[input2] = NewWire(false, -1)
        }
        if _,exists := allWires[output]; !exists {
            allWires[output] = NewWire(false, -1)
        }

        // defining gate type
        switch gateType {
        case "AND":
            newGate.typeOf = AND
        case "OR":
            newGate.typeOf = OR
        case "XOR":
            newGate.typeOf = XOR
        }
        // defining all the wires connections
        allWires[input1].appendToGate(newGate)
        allWires[input2].appendToGate(newGate)
        allWires[output].fromGate = newGate
        
        // defining all the gate connections
        newGate.input1 = allWires[input1]
        newGate.input2 = allWires[input2]
        newGate.output = allWires[output]

        allGates = append(allGates, newGate)
    }
    return allGates, allWires
}

func topological_processing(allGates []*Gate, allWires map[string]*Wire) int {
    q := []*Gate{}

    // fill the queue with the gates where in degree is defined
    for _, gate := range allGates {
        if gate.input1.isDefined && gate.input2.isDefined {
            q = append(q, gate)
        }
    }

    for len(q) > 0 {
        topGate := q[0]
        q = q[1:]

        // process the gate output wire
        if !topGate.input1.isDefined || !topGate.input2.isDefined {
            panic("Not Defined gates are being processed help!")
        }
        switch topGate.typeOf {
        case AND:  // AND gate case
            topGate.output.state = topGate.input1.state & topGate.input2.state
        case OR:  // OR gate case
            topGate.output.state = topGate.input1.state | topGate.input2.state
        case XOR:  // XOR gate case
            topGate.output.state = topGate.input1.state ^ topGate.input2.state
        }
        topGate.output.isDefined = true

        // Process all the gates where this wire act as an input
        // If both their input is now defined then append them to the queue
        for _, gate := range topGate.output.toGates {
            if gate.input1.isDefined && gate.input2.isDefined {
                q = append(q, gate)
            }
        }
    }

    res := 0
    for k,v := range allWires {
        if !v.isDefined {
            fmt.Println("All the wires should have been proccessed")
            panic("something is going horribly wrong")
        }
        if k[0]=='z' {
            shiftIndex := parseInt(k[1:])
            res |= (v.state<<shiftIndex)       
        }
    }
    return res
}

const (
    AND = 0
    OR = 1
    XOR = 2
)

func whichWire(allWires map[string]*Wire,w *Wire) string {
    for k, v := range allWires {
        if w == v {
            return k
        }
    }
    return "Wire not defined in the map\n"
}
func whichGate(mappedGates map[string]*Gate, g* Gate) string {
    for k, v := range mappedGates {
        if g == v {
            return k
        }
    }
    return "Gate not defined in the map\n"
}

func swapped_wires_config(allGates []*Gate, allWires map[string]*Wire) {
    // Mapping and finding the discrepencies in initial x00 and y00 wires
    gateMapping := make(map[string]*Gate)  // xor10,xor11, and10,and11,or1
    for i := range 45 {
        xi := fmt.Sprintf("x%02d",i)
        yi := fmt.Sprintf("y%02d",i)
        xwi := allWires[xi]
        ywi := allWires[yi]
        var xori0 *Gate
        var andi0 *Gate
        // lock in the xor gate 
        for _, gate := range xwi.toGates {
            if gate.typeOf == XOR  {  // 2 for XOR
                xori0 = gate 
            }
            if gate.typeOf == AND  {  // 0 for AND
                andi0 = gate 
            }
            if gate.typeOf == OR {
                fmt.Printf("Why is there an OR gate connecting the inputs\n")
            }
        }

        // check if the yi are also the same gates
        for _, gate := range ywi.toGates {
            if gate.typeOf == 2 && xori0 != gate {
                fmt.Printf("Discrepancy found at xor gate of %02d\n",i)
            }
            if gate.typeOf == 0 && andi0 != gate {
                fmt.Printf("Discrepancy found at and gate of %02d\n",i)
            }
        }
        xorGateKey := fmt.Sprintf("xor%02d|0",i)
        andGateKey := fmt.Sprintf("and%02d|0",i)
        gateMapping[xorGateKey] = xori0
        gateMapping[andGateKey] = andi0
    }
    if gateMapping["xor00|0"].output == allWires["z00"] {
        fmt.Printf("Correctly matched z00 output\n")
    }

    // Now we have correctly verified all the starting xor and and gates
    // I will now move on to checking the next XOR gate 
    // all the rest of the starting xor gates except from the 0th 
    // will have output to an AND gate and an XOR gate 
    // both these xor and and gates must have the output of previous and gate 
    // as their second input if not then we have found a discrepancy
    for i:=1; i<45; i++ {
        xorKey := fmt.Sprintf("xor%02d|0",i)
        firstXorOutputWire := gateMapping[xorKey]
        if firstXorOutputWire == nil {
            fmt.Printf("Something is going wrong \n")
            break
        }
        w01 := firstXorOutputWire.output

        // something is going on with the 39th wire
        var and02, xor02 *Gate
        for _, gate := range w01.toGates {
            if gate.typeOf == OR {
                fmt.Printf("Why is xor%02d|0 going to an OR gate\n",i)
            }
            if gate.typeOf == AND {
                and02 = gate
            }
            if gate.typeOf == XOR {
                xor02 = gate
            }
        }
        if and02 == nil {
            fmt.Printf("and gate is not reachable from xor%02d\n",i)
        }
        if xor02 == nil {
            fmt.Printf("xor2 gate is not reachable from xor%02d\n",i)
        }
        if xor02 != nil && and02 != nil {
            xorKey := fmt.Sprintf("xor%02d|1",i)
            andKey := fmt.Sprintf("and%02d|1",i)
            gateMapping[xorKey] = xor02
            gateMapping[andKey] = and02
            // output of the second xor must be mapped to the correct output wire
            outputWireKey := fmt.Sprintf("z%02d",i)
            if outputWireKey != whichWire(allWires,xor02.output) {
                fmt.Printf("Wrong output mapped to the %02d\n",i)
                fmt.Println(whichWire(allWires, xor02.output))
            }
            
            // hmt, and z18 have been swapped and confirmed by me
            // bfq, and z27 have been swapped and confirmed by me
            // hkh, and z31 have been swapped and confirmed by me
                


            gateMapping[xorKey] = xor02
            gateMapping[andKey] = and02
        }
        

        if i==39 {
            fmt.Println("Malfunctioning wire1",whichWire(allWires,w01)) // it is fjp
            fmt.Println("Length of toGates of fjp",len(w01.toGates))
            // for _, gate := range w01.toGates {
            //     
            // }
        }


        // done checking for xor discrepecies now will check for and
        

    } 

    gatestring := whichGate(gateMapping, allWires["jss"].fromGate)
    fmt.Println(gatestring)


    
    fmt.Printf("Reached here\n")
    
}

func Day24(){
    debug := 0
    f := []byte{}
    if debug == 0 {
        f, _ = os.ReadFile("inputs/data24.txt")
    } else {
        f, _ = os.ReadFile("inputs/data24dummy.txt")
    }
    allGates, allWires := d24_parse_input(f)
    res1 := topological_processing(allGates, allWires)
    fmt.Println(res1)
    swapped_wires_config(allGates, allWires)

    // a, o, x := 0, 0, 0
    // for _, v := range allGates {
    //     if v.typeOf == 0 {
    //         a++
    //     } else if v.typeOf == 1 {
    //         o++
    //     } else {
    //         x++
    //     }
    // }
    // fmt.Printf("AND gates:%d, OR gates:%d, XOR gates:%d\n",a,o,x)
    // fmt.Println(len(allGates), len(allWires))
}
