package main

import (
	"errors"
	"log"
	"math"
	"strconv"

	"github.com/fatih/color"
)

const (
	Empty = -1
)

var (
	Colors = []func(format string, a ...interface{})string{
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
	Explore(c Cell) ([]Cell, error)
	IsSatisfied() bool
}

type BoardImplementation struct {
	gridSize int
	cells    []Cell
}

func NewBoard(p *Problem) *BoardImplementation {
	numCells := p.gridSize * p.gridSize
	cells := make([]Cell, numCells)
	for i := 0; i < numCells; i++ {
		cells[i] = NewCell()
	}

	val := 0
	for i := 0; i < len(p.xCoords); i++ {
		x := p.xCoords[i]
		y := p.yCoords[i]
		index := x*p.gridSize + y
		err := cells[index].Fill((val))
		if err != nil {
			log.Fatal("Cell was already filled in - there are repeated values in the problem")
		}
		if math.Mod(float64(i), float64(2)) != float64(0) {
			val += 1
		}
	}

	return &BoardImplementation{
		gridSize: p.gridSize,
		cells:    cells,
	}
}

func (b *BoardImplementation) GetCell(x, y int) (Cell, error) {
	if x < 0 || x >= b.gridSize || y < 0 || y >= b.gridSize {
		return nil, errors.New("Cell out of bounds")
	}
	index := x*b.gridSize + y
	return b.cells[index], nil
}

// TODO(tanay)
func (b *BoardImplementation) Explore(c Cell) ([]Cell, error) {
	return nil, nil
}

// TODO(tanay)
func (b *BoardImplementation) IsSatisfied() bool {
	return true
}

func (b *BoardImplementation) String() string {
	reprString := "  "
	// ind := 0 //used for keeping track of colors

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
