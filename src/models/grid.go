package models

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"strings"

    c "mazes/models/cell"
)

type Grider interface {
    Deadends() []*c.Cell
    RandomCell() *c.Cell
    Size() int
    ToPng(filename string, cellSize int, background bool)
    ToString() []string

    Rows() <-chan []*c.Cell
    Cells() <-chan *c.Cell

    prepareGrid()
}

type Grid struct {
    rows, cols int
    grid [][]*c.Cell
    distances *c.Distances
}

func NewGrid(rows, cols int) *Grid {
    grid := &Grid{
        rows: rows,
        cols: cols,
        distances: &c.Distances{},
    }
    grid.prepareGrid()
    grid.configureCells()
    return grid
}

func (g *Grid) Cell(row, col int) *c.Cell {
    if (row >= 0 && row < g.rows) && (col >= 0 && col < g.cols) {
        return g.grid[row][col]
    }
    return nil
}

func (g *Grid) Cells() <-chan *c.Cell {
    ch := make(chan *c.Cell, 1)
    go func() {
        for _, row := range g.grid {
            for _, cell := range row {
                if cell != nil {
                    ch <- cell
                }
            }
        }
        close(ch)
    }()
    return ch
}

func (g *Grid) Rows() <-chan []*c.Cell {
    ch := make(chan []*c.Cell)
    go func() {
        for _, row := range g.grid {
            ch <- row
        }
        close(ch)
    }()
    return ch
}

func (g *Grid) Deadends() (list []*c.Cell) {
    for row := range g.grid {
        for col := range g.grid[0] {
            cell := g.grid[row][col]

            if len(cell.Links()) == 1 {
                list = append(list, cell)
            }
        }
    }
    return
}

func (g *Grid) RandomCell() *c.Cell {
    randRow := rand.Intn(g.rows)
    randCol := rand.Intn(g.cols)
    return g.grid[randRow][randCol]
}

func (g *Grid) Size() int {
    return g.rows * g.cols
}

func (g *Grid) contentsOf(cell *c.Cell) string {
    if val, ok := g.distances.Cell(cell); ok {
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
                cell = c.NewCell(-1, -1)
            }

            var body string
            if g.distances != nil {
                body = fmt.Sprintf(" %s ", g.contentsOf(cell)) // distance between 2 spaces
            } else {
                body = "   " // three spaces
            }
            corner := "+"

            eastBoundary := "│"
            if cell.Linked(cell.East()) {
                eastBoundary = " " // single space
            }
            top += body + eastBoundary

            southBoundary := "---"
            if cell.Linked(cell.South()) {
                southBoundary = "   " // three spaces
            }
            bot += southBoundary + corner
        }
        output = append(output, top + "\n")
        output = append(output, bot + "\n")
    }
    return output
}

func (g *Grid) backgroundColorFor(cell *c.Cell) color.RGBA {
    distance, ok := g.distances.Cell(cell)
    if !ok {
        return color.RGBA{0, 50, 0, 255}
    }

    _, maxDist := g.distances.Max()
    steps := float64(maxDist - (maxDist - distance))
    intensity := 1.0 - (steps / (float64(maxDist)/10 + steps))

    // dark := uint8(255 * intensity)
    bright := 20 + uint8(200 * intensity)

    return color.RGBA{bright, 0, bright, 255}
}

func (g *Grid) ToPng(filename string, cellSize int, background bool) {
    f, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    backgroundClr := color.RGBA{255, 255, 255, 255}
    wallsClr := color.RGBA{0, 0, 0, 255}
    wallsWidth := 6

    imgWidth := g.cols * cellSize + wallsWidth
    imgHeight := g.rows * cellSize + wallsWidth
    img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

    draw.Draw(img, img.Bounds(), &image.Uniform{backgroundClr}, image.ZP, draw.Src)

    drawMaze:
    for cell := range g.Cells() {
        cellRow, cellCol := cell.Row(), cell.Col()
        x0 := cellCol * cellSize
        y0 := cellRow * cellSize
        x1 := (cellCol + 1) * cellSize
        y1 := (cellRow + 1) * cellSize

        if background {
            cellBackgroundClr := g.backgroundColorFor(cell)
            drawRect(img, x0, y0, x1, y1, cellBackgroundClr)
            continue
        }

        if cell.North() == nil {
            drawRect(img, x0, y0, x1, y0, wallsClr)
        }
        if cell.West() == nil {
            drawRect(img, x0, y0, x0, y1, wallsClr)
        }
        if !cell.Linked(cell.East()) {
            drawRect(img, x1, y0, x1, y1, wallsClr)
        }
        if !cell.Linked(cell.South()) {
            drawRect(img, x0, y1, x1, y1, wallsClr)
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

func (g *Grid) prepareGrid() {
    grid := make([][]*c.Cell, g.rows)
    for row := range grid {
        column := make([]*c.Cell, g.cols)
        grid[row] = column

        for col := range column {
            grid[row][col] = c.NewCell(row, col)
        }
    }
    g.grid = grid
}

func (g *Grid) configureCells() {
    for cell := range g.Cells() {
        row, col := cell.Row(), cell.Col()
        if row-1 >= 0 {
            cell.SetNorth(g.grid[row-1][col])
        }
        if row+1 < g.rows {
            cell.SetSouth(g.grid[row+1][col])
        }
        if col-1 >= 0 {
            cell.SetWest(g.grid[row][col-1])
        }
        if col+1 < g.cols {
            cell.SetEast(g.grid[row][col+1])
        }
    }
}
