package main

import (
	"errors"
	"fmt"
)

type Problem interface {
	GridSize() int
	NumColors() int
	// ColorCoords returns a length two, two-d array
	// the first array is the "starting coord" for that color
	// the second is the "ending coord" for that color
	// pass in the index of the color you want to see (its value)
	ColorCoords(i int) ([][]int, error)
}

type ProblemImplementation struct {
	gridSize  int
	numColors int
	coords    [][]int
}

// Instantiate a Problem struct to represent the game board

func NewProblem(gridSize, numColors int, coords [][]int) (*ProblemImplementation, error) {
	p := &ProblemImplementation{
		gridSize:  gridSize,
		numColors: numColors,
		coords:    coords,
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

func (p *ProblemImplementation) ColorCoords(i int) ([][]int, error) {
	if i >= p.numColors || i < 0 {
		return nil, errors.New(fmt.Sprintf("index i, %d, out of range", i))
	}
	coords := make([][]int, 2)
	coords[0] = p.coords[i*2]
	coords[1] = p.coords[i*2+1]
	return coords, nil
}

// TODO(tanay) fix this based on what is available
func (p *ProblemImplementation) Validate() error {
	if len(p.coords)/2 != p.numColors {
		return errors.New(fmt.Sprintf("Number of coordinates, %d, and number of colors, %d, must be equal",
			len(p.coords)/2, p.numColors))
	}

	if p.numColors >= int(p.gridSize*p.gridSize/2) {
		return errors.New(fmt.Sprintf("Too many colors (%d) comapred "+
			"to number of cells (%d)", p.numColors, p.gridSize*p.gridSize))
	}

	outOfBounds := func(coord []int) bool {
		return coord[0] < 0 || coord[0] >= p.gridSize || coord[1] < 0 || coord[1] >= p.gridSize
	}

	for i := 0; i < p.numColors; i++ {
		if len(p.coords[i]) != 2 {
			return errors.New("Every coord should be a 2 length integer array")
		}

		if outOfBounds(p.coords[i]) {
			return errors.New(fmt.Sprintf("Coordinates %v are out of bounds", p.coords[i]))
		}
	}

	return nil
}
