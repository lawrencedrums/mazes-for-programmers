package main

import (
	"math/rand"
	"strings"
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

func (g *Grid) ToString() []string {
    var output []string
    output = append(output, "+" + strings.Repeat("---+", g.cols) + "\n")

    for row := range g.grid {
        top := "|"
        bot := "+"

        for col := range g.grid[0] {
            cell := g.grid[row][col]
            if cell == nil {
                cell = NewCell(-1, -1)
            }

            body := "   " // three spaces
            corner := "+"

            eastBoundary := "|"
            if cell.Linked(cell.East) {
                eastBoundary = " "
            }
            top += body + eastBoundary

            southBoundary := "---"
            if cell.Linked(cell.South) {
                southBoundary = "   " // three spaces
            }
            bot += southBoundary + corner
        }
        output = append(output, top + "\n")
        output = append(output, bot + "\n")
    }
    return output
}

func prepareGrid(grid *Grid) {
    for row := range grid.grid {
        for col := range grid.grid[0] {
            grid.grid[row][col] = NewCell(row, col)
        }
    }
}

func configureCells(grid *Grid) {
    for row := range grid.grid {
        for col := range grid.grid[0] {
            setNeighbors(row, col, grid)
        }
    }
}

func setNeighbors(row, col int, grid *Grid) {
    cell := grid.grid[row][col]
    if row-1 >= 0 {
        cell.North = grid.grid[row-1][col]
    }
    if row+1 < grid.rows {
        cell.South = grid.grid[row+1][col]
    }
    if col-1 >= 0 {
        cell.West = grid.grid[row][col-1]
    }
    if col+1 < grid.cols {
        cell.East = grid.grid[row][col+1]
    }
}

