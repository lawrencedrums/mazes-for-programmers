package generator

import (
    "math/rand"

    "mazes/models"
    c "mazes/models/cell"
)

func RecursiveBacktracker(grid models.Grider) {
    stack := []*c.Cell{grid.RandomCell()}

    for len(stack) > 0 {
        current := stack[len(stack)-1]

        var neighborsNoLinks []*c.Cell
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

func pop(stack []*c.Cell) (poppedCell *c.Cell, poppedStack []*c.Cell) {
    lastIdx := len(stack) - 1
    poppedCell, poppedStack = stack[lastIdx], stack[:lastIdx]
    return
}
