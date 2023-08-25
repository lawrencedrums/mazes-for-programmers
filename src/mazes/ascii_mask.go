package main

import (
    "mazes/generator"
    "mazes/grid"
)

func main() {
    mask := grid.MaskFromTxt("masks/mask.txt")
    grid := grid.NewMaskedGrid(mask)

    generator.RecursiveBacktracker(grid)

    grid.ToPng("png/ascii_mask.png", 60, false)
}
