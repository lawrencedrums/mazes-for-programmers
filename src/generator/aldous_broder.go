package generator

import (
    "math/rand"

    "mazes/grid"
)

func AldousBroder(grid grid.Grider) {
    cell := grid.RandomCell()
    unvisited := grid.Size() - 1

    for unvisited > 0 {
        neighbors := cell.Neighbors()
        randomNeighbor := neighbors[rand.Intn(len(neighbors))]

        if len(randomNeighbor.Links()) == 0 {
            cell.Link(randomNeighbor, true)
            unvisited -= 1
        }

        cell = randomNeighbor
    }
}
