package main

import "math/rand"

type HuntAndKill struct {}

func NewHuntAndKill() *HuntAndKill {
    return &HuntAndKill{}
}

func (h *HuntAndKill) on(grid *Grid) {
    current := grid.RandomCell()

    for current != nil {
        var unvisitedNeighbors []*Cell
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

            for row := range grid.grid {
                for col := range grid.grid[0] {
                    cell := grid.grid[row][col]

                    var visitedNeighbors []*Cell
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
}
