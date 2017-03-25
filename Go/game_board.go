package main

import (
	"errors"
	"log"
	"math"
	"strconv"

	"github.com/fatih/color"
)

var (
	Colors = []func(format string, a ...interface{}) string{
		color.BlueString,
		color.RedString,
		color.WhiteString,
		color.GreenString,
		color.YellowString,
		color.MagentaString,
		color.CyanString,
	}
)

// TODO(tanay) document interface methods
type Board interface {
	GetCell(x, y int) (Cell, error)
	IsSatisfied() bool
	Equals(board Board) bool
	Problem() Problem
}

type BoardImplementation struct {
	cells   []Cell
	problem Problem
}

func NewBoard(p Problem) *BoardImplementation {
	numCells := p.GridSize() * p.GridSize()
	cells := make([]Cell, numCells)
	for i := 0; i < numCells; i++ {
		cells[i] = NewCell()
	}

	val := 0
	for i := 0; i < p.NumColors()*2; i++ {
		coords, err := p.ColorCoords(i)
		if err != nil {
			log.Fatal(err)
		}
		x := coords[0]
		y := coords[1]
		index := x*p.GridSize() + y
		err = cells[index].Fill((val))
		if err != nil {
			log.Fatal("Cell was already filled in - there are repeated values in the problem")
		}
		if math.Mod(float64(i), float64(2)) != float64(0) {
			val += 1
		}
	}

	return &BoardImplementation{
		cells:   cells,
		problem: p,
	}
}

func (b *BoardImplementation) GetCell(x, y int) (Cell, error) {
	if x < 0 || x >= b.gridSize || y < 0 || y >= b.gridSize {
		return nil, errors.New("Cell out of bounds")
	}
	index := x*b.gridSize + y
	return b.cells[index], nil
}

// TODO(tanay) will only be useful in CSP search where you
// can make non-legal moves
func (b *BoardImplementation) IsSatisfied() bool {
	return true
}

func (b *BoardImplementation) Equals(board Board) bool {
	if b.Problem().GridSize() != board.Problem().GridSize() {
		return false
	}

	for x := 0; x < b.Problem().GridSize(); x++ {
		for y := 0; y < b.Problem().GridSize(); y++ {
			if b.GetCell(x, y).Val() != board.GetCell(x, y).Val() {
				return false
			}
		}
	}

	return true
}

func (b *BoardImplementation) Problem() Problem {
	return b.problem
}

func (b *BoardImplementation) String() string {
	reprString := "  "
	// column headers
	for i := 0; i < b.gridSize; i++ {
		reprString += strconv.Itoa(i)
		if i != b.gridSize-1 {
			reprString += "|"
		}
	}
	reprString += "\n"

	// fill in rest of grid
	for i := 0; i < b.gridSize; i++ {
		// row headers
		reprString += strconv.Itoa(i) + "|"
		for j := 0; j < b.gridSize; j++ {
			cell, err := b.GetCell(i, j)
			if err != nil {
				return err.Error()
			}
			if cell.Val() != Empty {
				colorInd := int(math.Mod(float64(cell.Val()), float64(len(Colors))))
				reprString += Colors[colorInd](cell.String()) + "|"
			} else {
				reprString += cell.String() + "|"
			}
		}
		reprString += "\n"
	}
	return reprString
}
