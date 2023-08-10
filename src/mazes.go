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
    swGrid.distances = distances

    // swGrid.ToPng(80)
    fmt.Println(swGrid.ToString())

    swGrid.distances = distances.PathTo(swGrid.grid[swGrid.rows-1][0])
    fmt.Println(swGrid.ToString())
}
