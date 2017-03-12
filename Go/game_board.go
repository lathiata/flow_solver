package main

import "errors"

const (
	Empty = -1
)
// TODO(tanay) document interface methods
type Board interface {
	GetCell(x, y int) (Cell, error)
	Explore(c Cell) ([]Cell, error)
	IsSatisfied() bool
}

type BoardImplementation struct {
	gridsize int
	cells    []Cell
}

func NewBoard(gridsize int) *BoardImplementation {
	numCells := gridsize * gridsize
	cells := make([]Cell, numCells)
	for i := 0; i < numCells; i++ {
		cells[i] = NewCell()
	}
	return &BoardImplementation{
		gridsize: gridsize,
		cells:    cells,
	}
}

func (b *BoardImplementation) GetCell(x, y int) (Cell, error) {
	if x >= b.gridsize || y >= b.gridsize {
		return nil, errors.New("Cell out of bounds")
	}
	index := x * b.gridsize + y
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