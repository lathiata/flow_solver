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
// TODO(tanay) might be able to remove Equals method (thanks to serialize)
type state interface {
	GetCell(x, y int) (Cell, error)
	AdjacentCells(x, y int) ([]Cell, error)
	IsSatisfied() bool
	IsSatisfiable() bool
	Equals(s state) bool
	Problem() Problem
	Serialize() string
	Copy() state
}

type stateImplementation struct {
	cells   []Cell
	problem Problem
	// Index into Problem.ColorCoords(). The one we haven't solved yet
	colorIndex int
	frontier   []Cell
}

func NewState(p Problem) *stateImplementation {
	cells := make([]Cell, 0)
	for i := 0; i < p.GridSize(); i++ {
		for j := 0; j < p.GridSize(); j++ {
			cells = append(cells, NewCell(i, j))
		}
	}

	s := &stateImplementation{
		cells:      cells,
		problem:    p,
		colorIndex: 0,
	}

	for val := 0; val < p.NumColors(); val++ {
		cells, err := p.ColorCoords(val)
		if err != nil {
			log.Fatal(err)
		}
		for _, cell := range cells {
			coords := cell.Coords()
			c, cerr := s.GetCell(coords[0], coords[1])
			if cerr != nil {
				log.Fatal(cerr)
			}
			c.Fill(val)
		}
	}

	return s
}

func (s *stateImplementation) inbounds(x, y int) bool {
	if x < 0 || x >= s.Problem().GridSize() || y < 0 || y >= s.Problem().GridSize() {
		return false
	}
	return true
}

func (s *stateImplementation) GetCell(x, y int) (Cell, error) {
	if !s.inbounds(x, y) {
		return nil, errors.New("Cell out of bounds")
	}

	index := x*s.Problem().GridSize() + y
	return s.cells[index], nil
}

func (s *stateImplementation) AdjacentCells(x, y int) ([]Cell, error) {
	adjacentCells := make([]Cell, 0)
	if !s.inbounds(x, y) {
		return nil, errors.New("Cell out of bounds")
	}

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			cell, err := s.GetCell(i, j)
			if err != nil && !(x == i && y == j) {
				adjacentCells = append(adjacentCells, cell)
			}
		}
	}
	return adjacentCells, nil
}

func (s *stateImplementation) IsSatsifiable() bool {
	return true
}

func (s *stateImplementation) IsSatisfied() bool {
	// first check that every cell is filled
	for _, cell := range s.cells {
		if cell.Empty() {
			return false
		}
	}
	// next check that each cell on the frontier is at the "end spot"
	// we don't need to check if each cell is next to an adjacent one
	// of the same color because we don't make illegal moves (for now)
	// TODO(tanay)
	return true
}

// TODO(tanay) can probably deprecate, because of the serialize function
func (s *stateImplementation) Equals(b state) bool {
	if b.Problem().GridSize() != s.Problem().GridSize() ||
		b.Problem().NumColors() != s.Problem().NumColors() {
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

func (s *stateImplementation) Serialize() string {
	serializedState := ""
	for x := 0; x < s.Problem().GridSize(); x++ {
		for y := 0; y < s.Problem().GridSize(); y++ {
			cell, err := s.GetCell(x, y)
			if err != nil {
				log.Fatal(err)
			}
			serializedState += strconv.Itoa(cell.Val())
		}
	}
	return serializedState
}

//TODO(tanay)
func (s *stateImplementation) Copy() *stateImplementation {
	return nil
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
			if !cell.Empty() {
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
