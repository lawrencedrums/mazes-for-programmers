package grid

import (
    "math/rand"

    "mazes/grid/cell"
)

type MaskedGrid struct {
    mask *Mask
    grid [][]*cell.Cell
}

func NewMaskedGrid(mask *Mask) (maskedGrid *MaskedGrid) {
    grid := make([][]*cell.Cell, mask.rows)
    for row := range grid {
        grid[row] = make([]*cell.Cell, mask.cols)
    }

    maskedGrid = &MaskedGrid{mask, grid}
    return
}

func (mg *MaskedGrid) RandomCell() *cell.Cell {
    row := rand.Intn(mg.mask.rows)
    col := rand.Intn(mg.mask.cols)
    return mg.grid[row][col]
}

func (mg *MaskedGrid) Size() int {
    return mg.mask.Count()
}

func prepareMaskedGrid(grid *MaskedGrid) {
    for row := range grid.mask.bits {
        for col := range grid.mask.bits[0] {
            if grid.mask.bits[row][col] {
                grid.grid[row][col] = cell.NewCell(row, col)
            }
        }
    }
}
