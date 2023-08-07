package main

import (
    "fmt"
)

func main() {
    btGrid := NewGrid(4, 6)
    swGrid := NewGrid(4, 6)

    bt := NewBinaryTree()
    bt.on(btGrid)
    fmt.Println(btGrid.ToString())

    sw := NewSiderwinder()
    sw.on(swGrid)
    fmt.Println(swGrid.ToString())
}
