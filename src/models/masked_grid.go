package models

import (
    "math/rand"

    c "mazes/models/cell"
)

type MaskedGrid struct {
    *Grid
    mask *Mask
}

func NewMaskedGrid(mask *Mask) (maskedGrid *MaskedGrid) {
    maskedGrid = &MaskedGrid{
        &Grid{
            rows: mask.rows,
            cols: mask.cols,
        },
        mask,
    }
    maskedGrid.prepareGrid()
    maskedGrid.configureCells()
    return
}

func (mg *MaskedGrid) RandomCell() *c.Cell {
    row := rand.Intn(mg.mask.rows)
    col := rand.Intn(mg.mask.cols)
    return mg.grid[row][col]
}

func (mg *MaskedGrid) Size() int {
    return mg.mask.Count()
}

func (mg *MaskedGrid) prepareGrid() {
    grid := make([][]*c.Cell, mg.rows)
    for row := range grid {
        column := make([]*c.Cell, mg.cols)
        grid[row] = column

        for col := range column {
            if mg.mask.bits[row][col] {
                grid[row][col] = c.NewCell(row, col)
            }
        }
    }
    mg.grid = grid
}
