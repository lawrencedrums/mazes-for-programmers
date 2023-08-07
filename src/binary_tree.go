package main

import "math/rand"

type BinaryTree struct {}

func NewBinaryTree() (bt *BinaryTree) {
    bt = &BinaryTree{}
    return
}

func (bt *BinaryTree) on(grid *Grid) {
    for row := range grid.grid {
        for col := range grid.grid[0] {
            cell := grid.grid[row][col]

            var neighbors []*Cell
            if cell.North != nil {
                neighbors = append(neighbors, cell.North)
            }
            if cell.East != nil {
                neighbors = append(neighbors, cell.East)
            }

            if len := len(neighbors); len > 0 {
                neighbor := neighbors[rand.Intn(len)]
                cell.Link(neighbor, true)
            }
        }
    }
}
