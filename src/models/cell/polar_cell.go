package models

import "slices"

type PolarCell struct {
    // clock-wise, counter clock-wise
    row, col int
    cw, ccw Celler
    inward Celler
    outward []Celler
    links map[*PolarCell] bool
}

func NewPolarCell(row, col int) *PolarCell {
    return &PolarCell{
        row,
        col,
        nil,
        nil,
        nil,
        []Celler{},
        make(map[*PolarCell]bool),
    }
}

func (p *PolarCell) Link(cell Celler, bidi bool) {}

func (p *PolarCell) Unlink(cell Celler, bidi bool) {}

func (p *PolarCell) Links() (keys []Celler) {
    for k := range p.links {
        keys = append(keys, k)
    }
    return
}

func (p *PolarCell) IsLinked(targetCell Celler) bool {
    return slices.Contains(p.Links(), targetCell)
}

func (p *PolarCell) Neighbors() (neighbors []Celler) {
    if p.cw != nil {
        neighbors = append(neighbors, p.cw)
    }
    if p.ccw != nil {
        neighbors = append(neighbors, p.ccw)
    }
    if p.inward != nil {
        neighbors = append(neighbors, p.inward)
    }
    neighbors = append(neighbors, p.outward...)
    return
}

func (p *PolarCell) Row() int {
    return p.row
}

func (p *PolarCell) Col() int {
    return p.col
}

func (p *PolarCell) Cw() Celler {
    return p.cw
}

func (p *PolarCell) Inward() Celler {
    return p.inward
}

func (p *PolarCell) SetCw(cell Celler) {
    p.cw = cell
}

func (p *PolarCell) SetCcw(cell Celler) {
    p.ccw = cell
}

func (p *PolarCell) SetInward(cell Celler) {
    p.inward = cell
}

func (p *PolarCell) AddToOutward(cell Celler) {
    p.outward = append(p.outward, cell)
}
