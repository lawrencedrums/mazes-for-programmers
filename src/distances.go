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
