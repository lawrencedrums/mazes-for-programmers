package generator

import (
    "math/rand"

    "mazes/models"
    c "mazes/models/cell"
)

func BinaryTree(grid models.Grider) {
    for cell := range grid.Cells() {
        var neighbors []*c.Cell
        if cell.North() != nil {
            neighbors = append(neighbors, cell.North())
        }
        if cell.East() != nil {
            neighbors = append(neighbors, cell.East())
        }

        if len := len(neighbors); len > 0 {
            neighbor := neighbors[rand.Intn(len)]
            cell.Link(neighbor, true)
        }
    }
}
