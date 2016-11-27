package sudoku

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Puzzle represents the 9x9 grid.
type Puzzle [9][9]Cell

// Cell represents a cell in the grid and encapsulates a value and a list of potential values as a bit map.
type Cell struct {
	rowIdx          int
	columnIdx       int
	value           *int
	potentialValues []bool
}

// Load ...
func (p *Puzzle) Load(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	rowIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, ",")

		if len(row) != 9 {
			return fmt.Errorf("Malformed row data: [%s]", line)
		}

		for columnIdx := 0; columnIdx < 9; columnIdx++ {
			val, err := strconv.Atoi(row[columnIdx])
			if err == nil {
				if val < 1 || val > 9 {
					return fmt.Errorf("Malformed row data, %d is not between 1 and 9: [%s]", val, line)
				}
				p[rowIdx][columnIdx] = Cell{rowIdx, columnIdx, &val, []bool{false, false, false, false, false, false, false, false, false}}
			} else {
				p[rowIdx][columnIdx] = Cell{rowIdx, columnIdx, nil, []bool{true, true, true, true, true, true, true, true, true}}
			}
		}

		rowIdx++

		if rowIdx == 9 {
			break
		}
	}

	if rowIdx < 9 {
		return fmt.Errorf("Invalid puzzle data.  Expected 9 rows, found %d row(s)", rowIdx)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Print prints out the puzzle to stdout with 'x's representing unknown values.
func (p *Puzzle) Print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if p[i][j].value != nil {
				fmt.Printf("%d", *p[i][j].value)
			} else {
				fmt.Printf("x")
			}
			if j < 8 {
				fmt.Printf(" ")
				if j == 2 || j == 5 {
					fmt.Printf("|| ")
				}
			}
		}
		fmt.Println()
		if i == 2 || i == 5 {
			fmt.Println("=======================")
		}
	}
}

// Solve attempts to solve the Puzzle by scanning updating the values of the cells.  If the puzzle is solvable
// this will return true, false otherwise.
func (p *Puzzle) Solve() bool {
	reduced := true
	// Loop until we can't reduce the potential values of any cell.
	for reduced {
		reduced = p.reducePossibilities()
	}
	return p.isSolved()
}

// isSolved returns false if any cell in the Puzzle has an unknown value, returns true otherwise.
func (p *Puzzle) isSolved() bool {
	solved := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if p[i][j].value == nil {
				solved = false
			}
		}
	}
	return solved
}

// reducePossibilities iterates the cells in the puzzle and attempts to reduce the possible values of the unknown cells
// by scanning rows, columns and 3x3 grids.
func (p *Puzzle) reducePossibilities() bool {
	reduced := false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := &p[i][j]
			if cell.value == nil {
				if p.reduceRow(cell) {
					reduced = true
				}
				if p.reduceColumn(cell) {
					reduced = true
				}
				if p.reduceGrid(cell) {
					reduced = true
				}
			}
		}
	}
	return reduced
}

// reduceRow will attempt to reduce the potential values of the given Cell by walking the other cells
// in the same row.
func (p *Puzzle) reduceRow(c *Cell) bool {
	reduced := false
	for i := 0; i < 9; i++ {
		val := p[c.rowIdx][i].value
		if val != nil && c.potentialValues[*val-1] {
			c.potentialValues[*val-1] = false
			reduced = true
		}
	}
	if reduced {
		c.updateValue()
	}
	return reduced
}

// reduceColumn will attempt to reduce the potential values of the given Cell by walking the other cells
// in the same column.
func (p *Puzzle) reduceColumn(c *Cell) bool {
	reduced := false
	for i := 0; i < 9; i++ {
		val := p[i][c.columnIdx].value
		if val != nil && c.potentialValues[*val-1] {
			c.potentialValues[*val-1] = false
			reduced = true
		}
	}
	if reduced {
		c.updateValue()
	}
	return reduced
}

// reduceGrid will attempt to reduce the potential values of the given Cell by walking the other cells
// in the 3x3 grid the given Cell belongs to.
func (p *Puzzle) reduceGrid(c *Cell) bool {
	reduced := false
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rowIdx := i + (3 * (c.rowIdx / 3))       // Offset to the correct 3x3 grid.
			columnIdx := j + (3 * (c.columnIdx / 3)) // Offset to the correct 3x3 grid.
			val := p[rowIdx][columnIdx].value
			if val != nil && c.potentialValues[*val-1] {
				c.potentialValues[*val-1] = false
				reduced = true
			}
		}
	}
	if reduced {
		c.updateValue()
	}
	return reduced
}

// updateValue updates the value of the cell if there is only one possibility remaining.
func (c *Cell) updateValue() {
	val := 0
	count := 0
	for k := 0; k < 9; k++ {
		if c.potentialValues[k] {
			val = k + 1
			count++
		}
	}
	if c.value == nil && count == 1 {
		c.value = &val
	}
}
