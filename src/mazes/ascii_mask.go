package main

import (
    "mazes/generator"
    "mazes/models"
)

func main() {
    mask := models.MaskFromTxt("masks/mask.txt")
    grid := models.NewMaskedGrid(mask)

    generator.RecursiveBacktracker(grid)

    grid.ToPng("png/ascii_mask.png", 60, false)
}
