package main

import (
    "mazes/generator"
    "mazes/models"
)

func main() {
    mask := models.MaskFromPng("masks/mask.png")
    grid := models.NewMaskedGrid(mask)

    generator.RecursiveBacktracker(grid)

    grid.ToPng("png/image_mask.png", 60, false)
}
