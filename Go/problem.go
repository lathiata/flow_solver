package main

import (
	"errors"
	"fmt"
	"log"
)

type Problem interface {
	GridSize() int
	NumColors() int
	// ColorCoords returns two cells, the "start" cell
	// and the "end" cell for color i
	ColorCoords(i int) ([]Cell, error)
}

type ProblemImplementation struct {
	gridSize  int
	numColors int
	coordMap  map[int][]Cell
}

// Instantiate a Problem struct to represent the game board
// Coords is a two-d array where
// 	Coords[i] =   (x1, y1)
//  Coords[i+1] = (x2, y2)
// for i%2 == 0
// these coordinates are the start, end for color i
func NewProblem(gridSize, numColors int, coords [][]int) (*ProblemImplementation, error) {
	coordMap := make(map[int][]Cell, numColors)
	for i := 0; i < numColors; i++ {
		x1 := coords[i*2][0]
		y1 := coords[i*2][1]
		cell1 := NewCell(x1, y1)
		err := cell1.Fill(i)
		if err != nil {
			log.Fatal(err)
		}
		x2 := coords[i*2+1][0]
		y2 := coords[i*2+1][1]
		cell2 := NewCell(x2, y2)
		err = cell2.Fill(i)
		if err != nil {
			log.Fatal(err)
		}

		coordMap[i] = []Cell{cell1, cell2}
	}
	p := &ProblemImplementation{
		gridSize:  gridSize,
		numColors: numColors,
		coordMap:  coordMap,
	}
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *ProblemImplementation) GridSize() int {
	return p.gridSize
}

func (p *ProblemImplementation) NumColors() int {
	return p.numColors
}

func (p *ProblemImplementation) ColorCoords(i int) ([]Cell, error) {
	if i >= p.numColors || i < 0 {
		return nil, errors.New(fmt.Sprintf("index i, %d, out of range", i))
	}
	return p.coordMap[i], nil
}

// TODO(tanay) fix this based on what is available
func (p *ProblemImplementation) Validate() error {
	if len(p.coordMap) != p.numColors {
		return errors.New(fmt.Sprintf("Number of coordinates, %d, and number of colors, %d, must be equal",
			len(p.coordMap), p.numColors))
	}

	if p.numColors >= int(p.gridSize*p.gridSize/2) {
		return errors.New(fmt.Sprintf("Too many colors (%d) comapred "+
			"to number of cells (%d)", p.numColors, p.gridSize*p.gridSize))
	}

	outOfBounds := func(coord []int) bool {
		return coord[0] < 0 || coord[0] >= p.gridSize || coord[1] < 0 || coord[1] >= p.gridSize
	}

	for i := 0; i < p.numColors; i++ {
		if len(p.coordMap[i]) != 2 {
			return errors.New("Every coord should be a 2 length integer array")
		}

		if outOfBounds(p.coordMap[i][0].Coords()) ||
			outOfBounds(p.coordMap[i][1].Coords()) {
			return errors.New(fmt.Sprintf("Coordinates %v are out of bounds", p.coordMap[i]))
		}
	}

	return nil
}
