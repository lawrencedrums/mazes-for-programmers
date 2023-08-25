package grid

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"strings"

    c "mazes/grid/cell"
)

type Grider interface {
    DeadEnds() []*c.Cell
    RandomCell() *c.Cell
    Size() int
    ToPng(filename string, cellSize int)
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
    cells := make([][]*c.Cell, rows)
    for i := range cells {
        cells[i] = make([]*c.Cell, cols)
    }

    grid := &Grid{
        rows: rows,
        cols: cols,
        grid: cells,
        distances: &c.Distances{},
    }
    grid.prepareGrid()
    grid.configureCells()
    return grid
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

func (g *Grid) DeadEnds() (list []*c.Cell) {
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

            body := fmt.Sprintf(" %s ", g.contentsOf(cell)) // distance between 2 spaces
            corner := "+"

            eastBoundary := "│"
            if cell.Linked(cell.East) {
                eastBoundary = " " // single space
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

func (g *Grid) ToPng(filename string, cellSize int) {
    f, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    imgWidth := g.cols * cellSize
    imgHeight := g.rows * cellSize
    img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

    background := true
    backgroundClr := color.RGBA{255, 255, 255, 255}
    wallsClr := color.RGBA{0, 0, 0, 255}

    draw.Draw(img, img.Bounds(), &image.Uniform{backgroundClr}, image.ZP, draw.Src)

    drawMaze:
    for row := range g.grid {
        for col := range g.grid[0] {
            cell := g.grid[row][col]

            x0 := cell.Col * cellSize
            y0 := cell.Row * cellSize
            x1 := (cell.Col + 1) * cellSize
            y1 := (cell.Row + 1) * cellSize

            if background {
                cellBackgroundClr := g.backgroundColorFor(cell)
                drawRect(img, x0, y0, x1, y1, cellBackgroundClr)
                continue
            }

            if cell.North == nil {
                drawRect(img, x0, y0, x1, y0, wallsClr)
            }
            if cell.West == nil {
                drawRect(img, x0, y0, x0, y1, wallsClr)
            }
            if !cell.Linked(cell.East) {
                drawRect(img, x1, y0, x1, y1, wallsClr)
            }
            if !cell.Linked(cell.South) {
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

func (g *Grid) prepareGrid() {
    for row := range g.grid {
        for col := range g.grid[0] {
            g.grid[row][col] = c.NewCell(row, col)
        }
    }
}

func (g *Grid) configureCells() {
    for row := range g.grid {
        for col := range g.grid[0] {
            cell := g.grid[row][col]
            if row-1 >= 0 {
                cell.North = g.grid[row-1][col]
            }
            if row+1 < g.rows {
                cell.South = g.grid[row+1][col]
            }
            if col-1 >= 0 {
                cell.West = g.grid[row][col-1]
            }
            if col+1 < g.cols {
                cell.East = g.grid[row][col+1]
            }
        }
    }
}
