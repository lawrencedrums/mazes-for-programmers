package main

import "math/rand"

type RecursiveBacktracker struct {}

func NewRecursiveBacktracker() *RecursiveBacktracker {
    return &RecursiveBacktracker{}
}

func (rb *RecursiveBacktracker) on(grid *Grid) {
    var stack []*Cell
    stack = append(stack, grid.RandomCell())

    for len(stack) > 0 {
        current := stack[len(stack)-1]

        var neighborsNoLinks []*Cell
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

func pop(stack []*Cell) (poppedCell *Cell, poppedStack []*Cell) {
    lastIdx := len(stack) - 1
    poppedCell, poppedStack = stack[lastIdx], stack[:lastIdx]
    return
}
