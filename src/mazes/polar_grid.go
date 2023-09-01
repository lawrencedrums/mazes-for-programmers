package main

import "mazes/models"

func main() {
    rows, cols := 20, 20
    g := models.NewPolarGrid(rows, cols)

    g.ToPng("png/polar_grid.png", 60, false)
}
