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
    for k := range c.links {
        keys = append(keys, k)
    }
    return
}

func (c *Cell) Linked(targetCell *Cell) bool {
    for _, cell := range c.Links() {
        if cell == targetCell {
            return true
        }
    }
    return false
}

func (c *Cell) Neighbors() (neighbors []*Cell) {
    if c.North != nil {
        neighbors = append(neighbors, c.North)
    }
    if c.East != nil {
        neighbors = append(neighbors, c.East)
    }
    if c.South != nil {
        neighbors = append(neighbors, c.South)
    }
    if c.West != nil {
        neighbors = append(neighbors, c.West)
    }
    return
}
