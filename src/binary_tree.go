package main

import "math/rand"

type BinaryTree struct {}

func NewBinaryTree() (bt *BinaryTree) {
    bt = &BinaryTree{}
    return
}

func (bt *BinaryTree) on(grid *Grid) {
    for i := 0; i < grid.rows; i++ {
        for j := 0; j <grid.cols; j++ {
            cell := grid.grid[i][j]

            neighbors := make([]*Cell, 0)
            if cell.North != (&Cell{}) {
                neighbors = append(neighbors, cell.North)
            }
            if cell.East != (&Cell{}) {
                neighbors = append(neighbors, cell.East)
            }

            if neighborLen := len(neighbors); neighborLen > 0 {
                randNeighborIdx := rand.Intn(neighborLen)
                neighbor := neighbors[randNeighborIdx]
                cell.Link(neighbor, false)
            }
        }
    }
}
