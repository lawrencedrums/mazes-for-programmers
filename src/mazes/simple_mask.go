package main

import (
    "fmt"

    "mazes/generator"
    "mazes/grid"
)

func main() {
    mask := grid.NewMask(5, 5)

    mask.SetBit(0, 0, false)
    mask.SetBit(2, 2, false)
    mask.SetBit(4, 4, false)

    grid := grid.NewMaskedGrid(mask)
    generator.RecursiveBacktracker(grid)

    fmt.Println(grid.ToString())
}
