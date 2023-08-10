package main

import (
    "fmt"
)

func main() {
    // btGrid := NewGrid(4, 6)
    swGrid := NewGrid(4, 6)

    // bt := NewBinaryTree()
    // bt.on(btGrid)

    // fmt.Println(btGrid.ToString())

    sw := NewSiderwinder()
    sw.on(swGrid)

    start := swGrid.grid[0][0]

    distances := start.Distances()
    newStart, _ := distances.max()

    newDistances := newStart.Distances()
    goal, _ := newDistances.max()

    swGrid.distances = newDistances.PathTo(goal)

    // swGrid.ToPng(80)
    // swGrid.distances = distances.PathTo(swGrid.grid[swGrid.rows-1][0])

    fmt.Println(swGrid.ToString())
}
