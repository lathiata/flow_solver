package main

import "errors"

// TODO(tanay) document interface methods
type Cell interface {
	Fill(val int) error
}

type CellImplementation struct {
	val int
}

func NewCell() *CellImplementation {
	return &CellImplementation{val: Empty}
}

func (c *CellImplementation) Fill(val int) error {
	if c.val != Empty {
		return errors.New("Can't fill a non-empty cell")
	}
	c.val = val
	return nil
}
