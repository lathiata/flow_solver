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
type state interface {
	GetCell(x, y int) (Cell, error)
	IsSatisfied() bool
	IsSatisfiable() bool
	Equals(s state) bool
	Problem() Problem
}

type stateImplementation struct {
	cells   []Cell
	problem Problem
}

func NewState(p Problem) *stateImplementation {
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

	return &stateImplementation{
		cells:   cells,
		problem: p,
	}
}

func (s *stateImplementation) GetCell(x, y int) (Cell, error) {
	if x < 0 || x >= s.Problem().GridSize() || y < 0 || y >= s.Problem().GridSize() {
		return nil, errors.New("Cell out of bounds")
	}
	index := x*s.Problem().GridSize() + y
	return s.cells[index], nil
}

func (s *stateImplementation) IsSatsifiable() bool {
	return true
}

// TODO(tanay) will only be useful in CSP search where you
// can make non-legal moves
func (s *stateImplementation) IsSatisfied() bool {
	return true
}

func (b *stateImplementation) Equals(s state) bool {
	if b.Problem().GridSize() != s.Problem().GridSize() {
		return false
	}

	for x := 0; x < b.Problem().GridSize(); x++ {
		for y := 0; y < b.Problem().GridSize(); y++ {
			cell1, err := b.GetCell(x, y)
			log.Fatal(err)
			cell2, err := s.GetCell(x, y)
			log.Fatal(err)
			if cell1.Val() != cell2.Val() {
				return false
			}
		}
	}

	return true
}

func (s *stateImplementation) Problem() Problem {
	return s.problem
}

func (s *stateImplementation) String() string {
	reprString := "  "
	// column headers
	for i := 0; i < s.Problem().GridSize(); i++ {
		reprString += strconv.Itoa(i)
		if i != s.Problem().GridSize()-1 {
			reprString += "|"
		}
	}
	reprString += "\n"

	// fill in rest of grid
	for i := 0; i < s.Problem().GridSize(); i++ {
		// row headers
		reprString += strconv.Itoa(i) + "|"
		for j := 0; j < s.Problem().GridSize(); j++ {
			cell, err := s.GetCell(i, j)
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
