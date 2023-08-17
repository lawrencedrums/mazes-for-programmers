package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"strings"
)

type Grid struct {
    rows, cols int
    grid [][]*Cell
    distances *Distances
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
        distances: &Distances{},
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

func (g *Grid) contentsOf(cell *Cell) string {
    if val, ok := g.distances.cells[cell]; ok {
        return fmt.Sprintf("%X", val) // base-16
    }
    return " " // single space
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

            body := fmt.Sprintf(" %s ", g.contentsOf(cell)) // distance between 2 spaces
            corner := "+"

            eastBoundary := "│"
            if cell.Linked(cell.east) {
                eastBoundary = " " // single space
            }
            top += body + eastBoundary

            southBoundary := "---"
            if cell.Linked(cell.south) {
                southBoundary = "   " // three spaces
            }
            bot += southBoundary + corner
        }
        output = append(output, top + "\n")
        output = append(output, bot + "\n")
    }
    return output
}

func (g *Grid) backgroundColorFor(cell *Cell) color.RGBA {
    distance, ok := g.distances.cells[cell]
    if !ok {
        return color.RGBA{0, 100, 0, 255}
    }

    _, maxDist := g.distances.max()
    intensity := float64(maxDist - distance) / float64(maxDist)
    dark := uint8(255 * intensity)
    bright := 128 + uint8(127 * intensity)

    return color.RGBA{dark, bright, dark, 255}
}

func (g *Grid) ToPng(background bool, cellSize int, filepath string) {
    f, err := os.Create(filepath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    imgWidth := g.cols * cellSize
    imgHeight := g.rows * cellSize
    imgMargin := 20

    backgroundClr := color.RGBA{255, 255, 255, 255}
    wallsClr := color.RGBA{0, 0, 0, 255}

    img := image.NewRGBA(
        image.Rect(
            -imgMargin,
            -imgMargin,
            imgWidth+imgMargin,
            imgHeight+imgMargin,
        ),
    )

    draw.Draw(img, img.Bounds(), &image.Uniform{backgroundClr}, image.ZP, draw.Src)

    drawMaze:
    for row := range g.grid {
        for col := range g.grid[0] {
            cell := g.grid[row][col]

            x0 := cell.col * cellSize
            y0 := cell.row * cellSize
            x1 := (cell.col + 1) * cellSize
            y1 := (cell.row + 1) * cellSize

            if background {
                cellBackgroundClr := g.backgroundColorFor(cell)
                drawRect(img, x0, y0, x1, y1, cellBackgroundClr)
                continue
            }

            if cell.north == nil {
                drawRect(img, x0, y0, x1, y0, wallsClr)
            }
            if cell.west == nil {
                drawRect(img, x0, y0, x0, y1, wallsClr)
            }
            if !cell.Linked(cell.east) {
                drawRect(img, x1, y0, x1, y1, wallsClr)
            }
            if !cell.Linked(cell.south) {
                drawRect(img, x0, y1, x1, y1, wallsClr)
            }
        }
    }

    if background {
        background = false
        goto drawMaze
    }

    if err = png.Encode(f, img); err != nil {
        fmt.Printf("Failed to encode: %v", err)
    }
}

func drawRect(img draw.Image, x0, y0, x1, y1 int, clr color.Color) {
    width := 6
    rect := image.Rect(x0, y0, x1+width, y1+width)
    draw.Draw(img, rect, &image.Uniform{clr}, image.ZP, draw.Src)
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
            cell := grid.grid[row][col]
            if row-1 >= 0 {
                cell.north = grid.grid[row-1][col]
            }
            if row+1 < grid.rows {
                cell.south = grid.grid[row+1][col]
            }
            if col-1 >= 0 {
                cell.west = grid.grid[row][col-1]
            }
            if col+1 < grid.cols {
                cell.east = grid.grid[row][col+1]
            }
        }
    }
}
