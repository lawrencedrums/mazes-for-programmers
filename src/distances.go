package main

type Distances struct {
    root *Cell
    cells map[*Cell]int
}

func NewDistances(root *Cell) (distances *Distances) {
    cells := make(map[*Cell]int)
    cells[root] = 0

    distances = &Distances{root: root, cells: cells}
    return
}

func (d *Distances) SetDistance(distance int ,cell *Cell) {
    d.cells[cell] = distance
}

// Cells returns list of all keys in cells hashmap
func (d *Distances) Cells() (keys []*Cell) {
    for k := range d.cells {
        keys = append(keys, k)
    }
    return
}

func (d *Distances) PathTo(goal *Cell) *Distances {
    current := goal

    breadcrumbs := NewDistances(d.root)
    breadcrumbs.cells[current] = d.cells[current]

    for current != d.root {
        for _, link := range current.Links() {
            if d.cells[link] < d.cells[current] {
                breadcrumbs.cells[link] = d.cells[link]
                current = link
                break
            }
        }
    }
    return breadcrumbs
}

func (d *Distances) Max() (*Cell, int) {
    maxDist := 0
    maxCell := d.root

    for cell, distance := range d.cells {
        if distance > maxDist {
            maxDist = distance
            maxCell = cell
        }
    }
    return maxCell, maxDist
}
