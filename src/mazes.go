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

    fmt.Println(swGrid.ToString())

    start := swGrid.grid[0][0]
    swGrid.distances = start.Distances()

    // swGrid.ToPng(80)
    fmt.Println(swGrid.ToString())
}
