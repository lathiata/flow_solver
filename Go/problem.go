package main

import (
	"errors"
	"fmt"
)

// TODO(tanay) should this be in its own file
type Problem struct {
	gridSize  int
	numColors int
	xCoords   []int
	yCoords   []int
}

// Instantiate a Problem struct to represent the game board

func NewProblem(gridSize, numColors int, xCoords, yCoords []int) *Problem {
	return &Problem{
		gridSize:  gridSize,
		numColors: numColors,
		xCoords:   xCoords,
		yCoords:   yCoords,
	}
}

func (p *Problem) Validate() error {
	if !(len(p.xCoords) == 2*p.numColors && len(p.yCoords) == 2*p.numColors) {
		if p.numColors >= int(p.gridSize*p.gridSize/2) {
			return errors.New(fmt.Sprintf("Too many colors (%d) comapred "+
				"to number of cells (%d)", p.numColors, p.gridSize*p.gridSize))
		}
		return errors.New(fmt.Sprintf("number of x coordinates (%d), "+
			"y coordinates (%d), and colors (%d) must be equal", len(p.xCoords), len(p.yCoords), p.numColors))
	}
	for i := 0; i < p.numColors; i++ {
		if p.xCoords[i] < 0 || p.xCoords[i] >= p.numColors || p.yCoords[i] < 0 || p.yCoords[i] >= p.numColors {
			return errors.New(fmt.Sprintf("Cell at (%d, %d) out of bounds", p.xCoords[i], p.yCoords[i]))
		}
	}

	return nil
}
