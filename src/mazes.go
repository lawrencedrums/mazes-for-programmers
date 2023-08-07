package main

import (
    "fmt"
)

func main() {
    grid := NewGrid(4, 6)

    bt := NewBinaryTree()
    bt.on(grid)

    maze := grid.ToString()
    fmt.Println(maze)
}
