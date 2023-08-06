package main

type Cell struct {
    row, col int
    North, South, East, West *Cell
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
    keys = make([]*Cell, 0, len(c.links))
    for k := range c.links {
        keys = append(keys, k)
    }
    return
}

func (c *Cell) Linked(targetCell *Cell) bool {
    for i, cell := range c.Links() {
        _ = i
        if cell == targetCell {
            return true
        }
    }
    return false
}

func (c *Cell) Neighbors() []*Cell {
    neighbors := []*Cell{}
    if c.North != (&Cell{}) {
        neighbors = append(neighbors, c.North)
    }
    if c.East != (&Cell{}) {
        neighbors = append(neighbors, c.East)
    }
    if c.South != (&Cell{}) {
        neighbors = append(neighbors, c.South)
    }
    if c.West != (&Cell{}) {
        neighbors = append(neighbors, c.West)
    }
    return neighbors
}
