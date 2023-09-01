package models

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
    "os"
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
            drawLine(img, ax, ay, cx, cy, wallsClr)
        }
        if !cell.Linked(cell.East()) {
            drawLine(img, cx, cy, dx, dy, wallsClr)
        }
    }
    drawCircle(img, center, center, pg.rows*cellSize, wallsClr)

    if err = png.Encode(f, img); err != nil {
        fmt.Print("Failed to encode: %v", err)
    }
}

func drawCircle(img draw.Image, x0, y0, radius int, clr color.Color) {
    // midpoint circle algorithm
    x, y, dx, dy := radius-1, 0, 1, 1
    err := dx - (radius * 2)

    for x > y {
        img.Set(x0+x, y0+y, clr)
        img.Set(x0+y, y0+x, clr)
        img.Set(x0-y, y0+x, clr)
        img.Set(x0-x, y0+y, clr)
        img.Set(x0-x, y0-y, clr)
        img.Set(x0-y, y0-x, clr)
        img.Set(x0+y, y0-x, clr)
        img.Set(x0+x, y0-y, clr)

        if err <= 0 {
            y++
            err += dy
            dy += 2
        }
        if err > 0 {
            x--
            dx += 2
            err += dx - (radius * 2)
        }
    }
}

func drawLine(img draw.Image, x0, y0, x1, y1 int, clr color.Color) {
    var xMin, yMin, xMax, yMax int

    if x0 < x1 {
        xMin, xMax = x0, x1
    } else {
        xMin, xMax = x1, x0
    }

    if y0 < y1 {
        yMin, yMax = y0, y1
    } else {
        yMin, yMax = y1, y0
    }

    slope := float64(y1 - y0) / float64(x1 - x0)
    yIntercept := float64(y0) - slope * float64(x0)

    // div by 0, draw nothing
    if math.IsNaN(slope) {
        return
    }

    // vertical line
    if math.IsInf(slope, 0) {
        for y := yMin; y <= yMax; y++ {
            img.Set(xMin, y, clr)
        }
        return
    }

    // normal slope
    if math.Abs(slope) <= 1 {
        for x := xMin; x <= xMax; x++ {
            y := int(float64(x) * slope + yIntercept)
            img.Set(x, y, clr)
        }
    } else {
        for y := yMin; y <= yMax; y++ {
            x := int((float64(y) - yIntercept) / slope)
            img.Set(x, y, clr)
        }
    }
}
