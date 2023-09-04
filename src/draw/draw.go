package draw

import (
    "image"
    "image/color"
    "image/draw"
    "math"
)

func DrawCircle(img draw.Image, x0, y0, radius int, clr color.Color) {
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
func DrawLine(img draw.Image, x0, y0, x1, y1 int, clr color.Color) {
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

func DrawRect(img draw.Image, x0, y0, x1, y1 int, clr color.Color) {
    rect := image.Rect(x0, y0, x1, y1)
    draw.Draw(img, rect, &image.Uniform{clr}, image.ZP, draw.Src)
}
