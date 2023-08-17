package main

import "math/rand"

type BinaryTree struct {}

func NewBinaryTree() *BinaryTree {
    return &BinaryTree{}
}

func (bt *BinaryTree) on(grid *Grid) {
    for row := range grid.grid {
        for col := range grid.grid[0] {
            cell := grid.grid[row][col]

            var neighbors []*Cell
            if cell.north != nil {
                neighbors = append(neighbors, cell.north)
            }
            if cell.east != nil {
                neighbors = append(neighbors, cell.east)
            }

            if len := len(neighbors); len > 0 {
                neighbor := neighbors[rand.Intn(len)]
                cell.Link(neighbor, true)
            }
        }
    }
}
