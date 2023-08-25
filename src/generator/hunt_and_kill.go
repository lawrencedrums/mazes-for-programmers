package generator

import (
    "math/rand"

    "mazes/models"
    c "mazes/models/cell"
)

func HuntAndKill(grid models.Grider) {
    current := grid.RandomCell()

    for current != nil {
        var unvisitedNeighbors []*c.Cell
        for _, neighbor := range current.Neighbors() {
            if len(neighbor.Links()) == 0 {
                unvisitedNeighbors = append(unvisitedNeighbors, neighbor)
            }
        }

        if len(unvisitedNeighbors) > 0 {
            neighbor := unvisitedNeighbors[rand.Intn(len(unvisitedNeighbors))]
            current.Link(neighbor, true)
            current = neighbor
        } else {
            current = nil

            for cell := range grid.Cells() {
                var visitedNeighbors []*c.Cell
                for _, neighbor := range cell.Neighbors() {
                    if len(neighbor.Links()) > 0 {
                        visitedNeighbors = append(visitedNeighbors, neighbor)
                    }
                }

                if len(cell.Links()) == 0 && len(visitedNeighbors) > 0 {
                    current = cell

                    neighbor := visitedNeighbors[rand.Intn(len(visitedNeighbors))]
                    current.Link(neighbor, true)

                    break
                }
            }
        }
    }
}
