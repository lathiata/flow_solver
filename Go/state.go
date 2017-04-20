package main

import (
	"errors"
	"log"
	"math"
	"strconv"

	"github.com/fatih/color"
	"reflect"
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

// TODO(tanay) implement a hueristic method (should be in interface probably)
// TODO(tanay) might be able to remove Equals method (thanks to serialize)
// The state interface actually wraps two things:
// 1) How you want to store internals (which cells are filled, etc
// 2) Transition model (NextStates)
// The methods defined are useful for implementing the pruned, threaded BFS
type state interface {
	IsSatisfied() bool
	IsSatisfiable() bool
	Problem() Problem
	Serialize() string
	Copy() state
	NextStates() []state
	Heuristic() int
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
	frontier := make([]Cell, p.NumColors())
	s := &stateImplementation{
		cells:      cells,
		problem:    p,
		colorIndex: 0,
		frontier:   frontier,
	}

	for val := 0; val < p.NumColors(); val++ {
		cells, err := p.ColorCoords(val)
		if err != nil {
			log.Fatal(err)
		}
		for i, cell := range cells {
			coords := cell.Coords()
			c, cerr := s.getCell(coords[0], coords[1])
			if cerr != nil {
				log.Fatal(cerr)
			}
			c.Fill(val)
			if i == 0 {
				// "start cell"
				s.frontier[val] = c
			}
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

func (s *stateImplementation) getCell(x, y int) (Cell, error) {
	if !s.inbounds(x, y) {
		return nil, errors.New("Cell out of bounds")
	}

	index := x*s.Problem().GridSize() + y
	return s.cells[index], nil
}

// Get a list of cells that are empty and next to the input cell
func (s *stateImplementation) adjacentCells(x, y int) ([]Cell, error) {
	adjacentCells := make([]Cell, 0)
	if !s.inbounds(x, y) {
		return nil, errors.New("Cell out of bounds")
	}

	possibleCoords := [][]int{
		[]int{x, y + 1},
		[]int{x, y - 1},
		[]int{x + 1, y},
		[]int{x - 1, y},
	}

	for _, coords := range possibleCoords {
		i := coords[0]
		j := coords[1]
		if s.inbounds(i, j) {
			cell, err := s.getCell(i, j)
			if err != nil {
				log.Fatal(err)
			}
			if !reflect.DeepEqual(coords, []int{x, y}) && cell.Empty() {
				adjacentCells = append(adjacentCells, cell)
			}

		}
	}
	return adjacentCells, nil
}

// TODO(tanaylathia)
func (s *stateImplementation) IsSatisfiable() bool {
	// for each cell, do a DFS (astar?) search
	// from its current position on the frontier
	// to the end state. If, the end state is reached for
	// each cell, return true. If any color is unable to
	// reach its end cell, return false.
	return true
}

func (s *stateImplementation) IsSatisfied() bool {
	// first check all cells filled
	for _, cell := range s.cells {
		if cell.Empty() {
			return false
		}
	}

	// TODO(tanay) now check all cells on frontier are adjacent to end cell for that color
	return true
}

func distance(c1, c2 Cell) float64 {
	coords1 := c1.Coords()
	coords2 := c2.Coords()
	return math.Sqrt(math.Pow(float64(coords1[0]-coords2[0]), 2.0) +
		math.Pow(float64(coords1[1]-coords2[1]), 2.0))
}

func (s *stateImplementation) Heuristic() int {
	totalDistance := 0
	for i, frontierCell := range s.frontier {
		cells, err := s.Problem().ColorCoords(i)
		if err != nil {
			log.Fatal(err)
		}
		totalDistance += int(distance(frontierCell, cells[1]))
	}

	return totalDistance
}

func (s *stateImplementation) Problem() Problem {
	return s.problem
}

func (s *stateImplementation) Serialize() string {
	serializedState := ""
	for x := 0; x < s.Problem().GridSize(); x++ {
		for y := 0; y < s.Problem().GridSize(); y++ {
			cell, err := s.getCell(x, y)
			if err != nil {
				log.Fatal(err)
			}
			serializedState += strconv.Itoa(cell.Val())
		}
	}
	return serializedState
}

func (s *stateImplementation) Copy() state {
	cellCopy := make([]Cell, len(s.cells))
	for i, cell := range s.cells {
		coords := cell.Coords()
		cellCopy[i] = NewCell(coords[0], coords[1])
		cellCopy[i].Fill(cell.Val())
	}

	copy := &stateImplementation{
		cells:      cellCopy,
		problem:    s.Problem(),
		colorIndex: s.colorIndex,
	}

	frontierCopy := make([]Cell, len(s.frontier))
	for i, cell := range s.frontier {
		coords := cell.Coords()
		copyCell, err := s.getCell(coords[0], coords[1])
		if err != nil {
			log.Fatal(err)
		}
		frontierCopy[i] = copyCell
	}

	copy.frontier = frontierCopy
	return copy
}

func (s *stateImplementation) String() string {
	reprString := "\n  "
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
			cell, err := s.getCell(i, j)
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

func (s *stateImplementation) areCellsAdjacent(c1, c2 Cell) bool {
	return distance(c1, c2) == 1.0
}

// The transition model this implements will "solve" one color
// at a time (explores every legal move until finds "end" cell)
//
// TODO(tanaylathia) need to initialize frontier/colorIndex in NewState - increment color index for already solved colors
func (s *stateImplementation) NextStates() []state {
	// First get the "frontier" cell and the
	// end cell for the corresponding color
	frontierCell := s.frontier[s.colorIndex]
	frontierCellCoords := frontierCell.Coords()
	colorCells, err := s.problem.ColorCoords(s.colorIndex)
	endCell := colorCells[1]
	if err != nil {
		log.Fatal(err)
	}
	if s.areCellsAdjacent(frontierCell, endCell) {
		s.colorIndex += 1
		frontierCell = s.frontier[s.colorIndex]
		frontierCellCoords = frontierCell.Coords()

	}

	// Based on the correct frontierCell (if above protocol is correct)
	// we go through all possible moves and create next states for them
	possibleMoves, err := s.adjacentCells(frontierCellCoords[0], frontierCellCoords[1])
	if err != nil {
		log.Fatal(err)
	}
	nextStates := make([]state, len(possibleMoves))
	for i, move := range possibleMoves {
		nextState := s.Copy()
		castedNextState := nextState.(*stateImplementation)
		moveCoords := move.Coords()
		cell, merr := castedNextState.getCell(moveCoords[0], moveCoords[1])
		if merr != nil {
			log.Fatal(err)
		}
		// Update info in nextState
		cell.Fill(s.colorIndex)
		castedNextState.frontier[s.colorIndex] = cell
		nextStates[i] = castedNextState
	}

	return nextStates
}
