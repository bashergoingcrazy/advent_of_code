package days

import (
	"fmt"
	"image/color"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

var frameCounter int = 0
func saveFrame(matrix [][]rune) {
	const cellSize = 20 // Size of each cell in pixels
	rows := len(matrix)
	cols := len(matrix[0])
	width := cols * cellSize
	height := rows * cellSize

	dc := gg.NewContext(width, height)

	// Iterate through the matrix and draw cells
	for i, row := range matrix {
		for j, r := range row {
			switch r {
			case '#':
				dc.SetColor(color.RGBA{0, 255, 255, 255}) // Cyan for walls
			case '[':
				dc.SetColor(color.RGBA{255, 255, 0, 255}) // Yellow for '['
			case ']':
				dc.SetColor(color.RGBA{255, 255, 0, 255}) // Yellow for ']'
			case '@':
				dc.SetColor(color.RGBA{255, 0, 0, 255}) // Red for robot
			default:
				dc.SetColor(color.RGBA{255, 255, 255, 255}) // White for empty
			}
			dc.DrawRectangle(float64(j*cellSize), float64(i*cellSize), float64(cellSize), float64(cellSize))
			dc.Fill()
		}
	}

	// Save the frame to a PNG file
	filename := fmt.Sprintf("frames/frame_%05d.png", frameCounter)
	dc.SavePNG(filename)
	frameCounter++
}

func to_matrix(board string) [][]rune {
    b := strings.Split(board, "\n")
    matrix := make([][]rune, 0, len(b))
    for _, row := range b {
        rr := make([]rune,0,len(row)) 
        for _, r := range row {
            rr = append(rr, r)
        }
        matrix = append(matrix, rr)
    }
    return matrix
}

func d15_print_matrix(matrix [][]rune) {
    for _, row := range matrix {
        for _, r := range row {
            fmt.Printf("%c ",r)
        }
        fmt.Printf("\n")
    }
}

func move(matrix [][]rune, m string) {
    // get initial point of the robot
    z := Pair{-1,-1}
    for i, row := range matrix {
        for j, r := range row {
            if r == '@' {
                z.x = i 
                z.y = j
                break
            }
        }
        if z.x != -1 {break}
    }
    fmt.Println("Index of robot",z)

    dir := []Pair {
        {0,-1},
        {-1,0},
        {0,1},
        {1,0},
    }
    for _, d := range m {
        if d == '\n' {continue}
        var ad Pair
        switch d {
        case '<':
            ad = dir[0] 
        case '^':
            ad = dir[1] 
        case '>':
            ad = dir[2] 
        case 'v':
            ad = dir[3] 
        }

        // if there is a simple empty move make it 
        if matrix[z.x+ad.x][z.y + ad.y] == '.' {
            matrix[z.x][z.y] = '.'
            // move the robot's  index
            z.x = z.x + ad.x
            z.y = z.y + ad.y
            // also make the move in the matrix
            matrix[z.x][z.y] = '@'
            continue
        }

        // for complex movements
        b := Pair{z.x+ad.x, z.y + ad.y}
        // move this index till we have boxes at b
        for matrix[b.x][b.y] == 'O' {
            b.x = b.x + ad.x 
            b.y = b.y + ad.y
        }
        // check if this index is event suitable if it is # we got a return
        if matrix[b.x][b.y] == '#' { continue }
        // fmt.Println(b,z)

        matrix[b.x][b.y] = 'O'
        matrix[z.x][z.y] = '.'
        z.x = z.x + ad.x
        z.y = z.y + ad.y
        matrix[z.x][z.y] = '@'
    }
    d15_print_matrix(matrix)
}

func GpsSum(matrix [][]rune) int {
    count := 0
    for i, row := range matrix {
        for j, val := range row {
            if val == 'O' || val == '['{
                count += (i*100) + j
            }
        }
    }
    return count
}

func to_matrix2(board string) [][]rune {
    b := strings.Split(board, "\n")
    matrix := make([][]rune, 0, len(b))
    for _, row := range b {
        rr := make([]rune,0,2*len(row)) 
        for _, r := range row {
            if r == 'O' {
                rr = append(rr, '[')
                rr = append(rr, ']')
            } else if r == '@' {
                rr = append(rr, '@')
                rr = append(rr, '.')
            } else {
                rr = append(rr, r)
                rr = append(rr, r)
            }
        }
        matrix = append(matrix, rr)
    }
    return matrix
}

func move2(matrix [][]rune, m string) {
    // get initial point of the robot
    z := Pair{-1,-1}
    for i, row := range matrix {
        for j, r := range row {
            if r == '@' {
                z.x = i 
                z.y = j
                break
            }
        }
        if z.x != -1 {break}
    }
    fmt.Println("Index of robot",z)

    dir := []Pair {
        {0,-1},
        {-1,0},
        {0,1},
        {1,0},
    }
    for _, d := range m {
        if d == '\n' {continue}
        var ad Pair
        switch d {
        case '<':
            ad = dir[0] 
        case '^':
            ad = dir[1] 
        case '>':
            ad = dir[2] 
        case 'v':
            ad = dir[3] 
        }
        // saveFrame(matrix)

        // if it is a simple empty move make it 
        if matrix[z.x+ad.x][z.y + ad.y] == '.' {
            matrix[z.x][z.y] = '.'
            // move the robot's  index
            z.x = z.x + ad.x
            z.y = z.y + ad.y
            // also make the move in the matrix
            matrix[z.x][z.y] = '@'
            continue
        }
        b := Pair{z.x+ad.x, z.y + ad.y}

        // for complex left and right movements 
        if d == '<' || d == '>' {
            // move this index till we have boxes at b
            for matrix[b.x][b.y] == '[' || matrix[b.x][b.y] == ']' {
                b.x = b.x + ad.x 
                b.y = b.y + ad.y
            }
            // check if this index is even suitable if it is # we gotta continue
            if matrix[b.x][b.y] == '#' { continue }
            // fmt.Println(b,z)

            if d == '<' {
                matrix[b.x][b.y] = '['
            } else {
                matrix[b.x][b.y] = ']'
            }
            b.x = b.x - ad.x
            b.y = b.y - ad.y
            for b.y!=z.y {
                if matrix[b.x][b.y] == '[' {
                    matrix[b.x][b.y] = ']'
                } else {
                    matrix[b.x][b.y] = '['
                }
                b.x = b.x - ad.x
                b.y = b.y - ad.y
            }
            matrix[z.x][z.y] = '.'
            z.x = z.x + ad.x
            z.y = z.y + ad.y
            matrix[z.x][z.y] = '@'
            continue
        }

        // for complex up and down movements
        valid, boxes := bfsd15(b,ad,matrix)
        if !valid {
            continue
        }

        // iterate the boxes slice from reverse and update matrix
        slices.Reverse(boxes)
        for _, p := range boxes {
            matrix[p.x + ad.x][p.y + ad.y] = '['
            matrix[p.x][p.y] = '.'
            matrix[p.x + ad.x][p.y + ad.y+1] = ']'
            matrix[p.x][p.y+1] = '.'
        }
        // lastly update the position of robot
        matrix[z.x][z.y] = '.'
        z.x = z.x + ad.x
        z.y = z.y + ad.y
        matrix[z.x][z.y] = '@'
    
    }
    // d15_print_matrix(matrix)
}

func bfsd15(b,ad Pair, matrix [][]rune) (bool, []Pair){
    var boxes []Pair
    var queue []Pair
    if matrix[b.x][b.y] == '#' {
        return false, nil
    } else if matrix[b.x][b.y] == ']' {
        queue = append(queue, Pair{b.x,b.y-1})
    } else {
        queue = append(queue, b)
    }

    for len(queue) > 0 {
        currentBox := queue[0]
        boxes = append(boxes, currentBox)
        queue = queue[1:]

        n1 := Pair{currentBox.x + ad.x,currentBox.y + ad.y}
        n2 := Pair{currentBox.x + ad.x,currentBox.y + ad.y + 1}
        if matrix[n1.x][n1.y] == '#' || matrix[n2.x][n2.y] == '#' {
            return false, nil
        } 

        // add corresponding n1
        if matrix[n1.x][n1.y] == '[' {
            queue = append(queue, n1)
        } else if matrix[n1.x][n1.y] == ']' {
            queue = append(queue, Pair{n1.x, n1.y - 1})
        }

        // add corresponding n2
        if matrix[n2.x][n2.y] == '[' {
            queue = append(queue, n2)
        } else if matrix[n2.x][n2.y] == ']' {
            queue = append(queue, Pair{n2.x, n2.y - 1})
        }
    }
    return true, boxes
}



func Day15() {
    f, err := os.ReadFile("inputs/data15.txt")
    if err != nil {
        fmt.Println("Can' Read file ",err)
        os.Exit(2)
    }
    d := string(f)
    d2 := strings.Split(d, "\n\n")
    boardString := strings.TrimSpace(d2[0])
    movements := strings.TrimSpace(d2[1])

    // solution for part 1
    // matrix := to_matrix(boardString)
    // move(matrix, movements)
    // count := GpsSum(matrix)
    // fmt.Println("Count:",count)

    t := time.Now()
    // solution for part 2
    matrix2 := to_matrix2(boardString)
    // d15_print_matrix(matrix2)
    move2(matrix2, movements)
    count2 := GpsSum(matrix2)
    fmt.Println(count2)
    tf := time.Now()
    elapsed := tf.Sub(t)
    fmt.Println(elapsed)


}
