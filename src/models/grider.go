package models

import cell "mazes/models/cell"

type Grider interface {
    RandomCell() cell.Celler
    ToPng(filename string, cellSize int, background bool)

    Cells() <-chan cell.Celler

    prepareGrid()
}
