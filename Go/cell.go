package main

import (
	"errors"
	"strconv"
)

const (
	Empty = -1
)

// TODO(tanay) document interface methods
type Cell interface {
	Fill(val int) error
	Val() int
	String() string
}

type CellImplementation struct {
	val int
}

func NewCell() *CellImplementation {
	return &CellImplementation{val: Empty}
}

func (c *CellImplementation) Val() int {
	return c.val
}

func (c *CellImplementation) Fill(val int) error {
	if c.val != Empty {
		return errors.New("Can't fill a non-empty cell")
	}
	c.val = val
	return nil
}

func (c *CellImplementation) String() string {
	if c.val == Empty {
		return " "
	}
	return strconv.Itoa(c.val)
}
