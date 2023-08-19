package main

import (
	"math/rand"
	"slices"
)

type Cell struct {
    row, col int
    north, south, east, west *Cell
    links map[*Cell]bool
}

func NewCell(row, col int) (cell *Cell) {
    links := make(map[*Cell]bool)
    cell = &Cell{row: row, col: col, links: links}
    return
}

func (c *Cell) Link(cell *Cell, bidi bool) {
    c.links[cell] = true
    if bidi {
        cell.Link(c, false)
    }
}

func (c *Cell) Unlink(cell *Cell, bidi bool) {
    delete(c.links, cell)
    if bidi {
        cell.Unlink(c, false)
    }
}

func (c *Cell) Links() (keys []*Cell) {
    for k := range c.links {
        keys = append(keys, k)
    }
    return
}

func (c *Cell) Linked(targetCell *Cell) bool {
    return slices.Contains(c.Links(), targetCell)
}

func (c *Cell) Neighbors() (neighbors []*Cell) {
    if c.north != nil {
        neighbors = append(neighbors, c.north)
    }
    if c.east != nil {
        neighbors = append(neighbors, c.east)
    }
    if c.south != nil {
        neighbors = append(neighbors, c.south)
    }
    if c.west != nil {
        neighbors = append(neighbors, c.west)
    }
    return
}

func (c *Cell) RandomNeighbor() *Cell {
    neighbors := c.Neighbors()
    return neighbors[rand.Intn(len(neighbors))]
}

func (c *Cell) Distances() *Distances {
    dist := NewDistances(c)

    var frontier []*Cell
    frontier = append(frontier, c)

    for len(frontier) > 0 {
        var newFrontier []*Cell

        for _, cell := range frontier {
            for _, linked := range cell.Links() {
                if _, ok := dist.cells[linked]; !ok {
                    dist.cells[linked] = dist.cells[cell] + 1
                    newFrontier = append(newFrontier, linked)
                }
            }
        }
        frontier = newFrontier
    }
    return dist
}
