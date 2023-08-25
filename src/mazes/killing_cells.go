package main

import (
    "fmt"

	"mazes/generator"
	"mazes/models"
)

func main() {
    grid := models.NewGrid(8, 8)

    grid.Cell(0, 0).East().SetWest(nil)
    grid.Cell(0, 0).South().SetNorth(nil)

    grid.Cell(7, 7).West().SetEast(nil)
    grid.Cell(7, 7).North().SetSouth(nil)

    generator.RecursiveBacktracker(grid)
    fmt.Println(grid.ToString())
}
