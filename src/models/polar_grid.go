package models

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
    "os"

    d "mazes/draw"
)

type PolarGrid struct {
    *Grid
}

func NewPolarGrid(rows, cols int) *PolarGrid {
    g := &PolarGrid{
        &Grid{
            rows: rows,
            cols: cols,
        },
    }
    g.prepareGrid()
    g.configureCells()
    return g
}

func (pg *PolarGrid) ToPng(filename string, cellSize int, background bool) {
    var (
        theta float64
        inner_radius float64
        outer_radius float64
        theta_ccw float64
        theta_cw float64

        ax, ay, cx, cy, dx, dy int
    )

    f, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    margin := 20
    imgSize := 2 * pg.rows * cellSize + margin
    center := imgSize / 2

    backgroundClr := color.RGBA{255, 255, 255, 255}
    wallsClr := color.RGBA{0, 0, 0, 255}

    img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

    draw.Draw(img, img.Bounds(), &image.Uniform{backgroundClr}, image.ZP, draw.Src)

    for cell := range pg.Cells() {
        theta        = 2.0 * math.Pi / float64(len(pg.grid[cell.Row()]))
        inner_radius = float64(cell.Row() * cellSize)
        outer_radius = float64((cell.Row() + 1) * cellSize)
        theta_ccw    = float64(cell.Col()) * theta
        theta_cw     = float64(cell.Col() + 1) * theta

        ax = center + int(inner_radius * math.Cos(theta_ccw))
        ay = center + int(inner_radius * math.Sin(theta_ccw))
        // bx = center + int(outer_radius * math.Cos(theta_ccw))
        // by = center + int(outer_radius * math.Sin(theta_ccw))
        cx = center + int(inner_radius * math.Cos(theta_cw))
        cy = center + int(inner_radius * math.Sin(theta_cw))
        dx = center + int(outer_radius * math.Cos(theta_cw))
        dy = center + int(outer_radius * math.Sin(theta_cw))

        if !cell.Linked(cell.North()) {
            d.DrawLine(img, ax, ay, cx, cy, wallsClr)
        }
        if !cell.Linked(cell.East()) {
            d.DrawLine(img, cx, cy, dx, dy, wallsClr)
        }
    }
    d.DrawCircle(img, center, center, pg.rows*cellSize, wallsClr)

    if err = png.Encode(f, img); err != nil {
        fmt.Print("Failed to encode: %v", err)
    }
}
