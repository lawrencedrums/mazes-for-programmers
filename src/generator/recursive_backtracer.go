package generator

import (
    "math/rand"

    "mazes/grid"
    "mazes/grid/cell"
)

func RecursiveBacktracker(grid grid.Grider) {
    stack := []*cell.Cell{grid.RandomCell()}

    for len(stack) > 0 {
        current := stack[len(stack)-1]

        var neighborsNoLinks []*cell.Cell
        for _, neighbor := range current.Neighbors() {
            if len(neighbor.Links()) == 0 {
                neighborsNoLinks = append(neighborsNoLinks, neighbor)
            }
        }

        if len(neighborsNoLinks) == 0 {
            _, stack = pop(stack)
        } else {
            neighbor := neighborsNoLinks[rand.Intn(len(neighborsNoLinks))]
            current.Link(neighbor, true)
            stack = append(stack, neighbor)
        }
    }
}

func pop(stack []*cell.Cell) (poppedCell *cell.Cell, poppedStack []*cell.Cell) {
    lastIdx := len(stack) - 1
    poppedCell, poppedStack = stack[lastIdx], stack[:lastIdx]
    return
}
