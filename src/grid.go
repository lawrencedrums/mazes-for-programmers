package main

import (
	"math/rand"
)

type Grid struct {
    rows, cols int
    grid [][]*Cell
}

func NewGrid(rows, cols int) *Grid {
    cells := make([][]*Cell, rows)
    for i := range cells {
        cells[i] = make([]*Cell, cols)
    }

    grid := &Grid{
        rows: rows,
        cols: cols,
        grid: cells,
    }
    prepareGrid(grid)
    configureCells(grid)
    return grid
}

func (g *Grid) RandomCell() *Cell {
    randRow := rand.Intn(g.rows)
    randCol := rand.Intn(g.cols)
    return g.grid[randRow][randCol]
}

func (g *Grid) Size() int {
    return g.rows * g.cols
}

func prepareGrid(grid *Grid) {
    for i := 0; i < grid.rows; i++ {
        for j := 0; j < grid.cols; j++ {
            grid.grid[i][j] = NewCell(i, j)
        }
    }
}

func configureCells(grid *Grid) {
    for row := 0; row < grid.rows; row++ {
        for col := 0; col < grid.cols; col++ {
            setNeighbors(row, col, grid)
        }
    }
}

func setNeighbors(row, col int, grid *Grid) {
    cell := grid.grid[row][col]
    if row-1 > 0 {
        cell.North = grid.grid[row-1][col]
    }
    if row+1 < grid.rows {
        cell.South = grid.grid[row+1][col]
    }
    if col-1 > 0 {
        cell.West = grid.grid[row][col-1]
    }
    if col+1 < grid.cols {
        cell.East = grid.grid[row][col+1]
    }
}

