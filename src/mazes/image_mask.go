package main

import (
    "mazes/generator"
    "mazes/grid"
)

func main() {
    mask := grid.MaskFromPng("masks/mask.png")
    grid := grid.NewMaskedGrid(mask)

    generator.RecursiveBacktracker(grid)

    grid.ToPng("png/image_mask.png", 60, false)
}
