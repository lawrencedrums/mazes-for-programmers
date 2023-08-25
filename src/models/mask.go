package models

import (
	"bufio"
	"image"
	"image/color"
	"math/rand"
	"os"
)

type Mask struct {
    rows, cols int
    bits [][]bool
}

func NewMask(rows, cols int) *Mask {
    bits := make([][]bool, rows)
    for row := range bits {
        bits[row] = make([]bool, cols)

        for col := range bits[0] {
            bits[row][col] = true
        }
    }
    return &Mask{rows, cols, bits}
}

func (m *Mask) Count() (count int) {
    for row := range m.bits {
        for col := range m.bits[0] {
            if m.bits[row][col] {
                count++
            }
        }
    }
    return
}

func (m *Mask) RandomLocation() (row, col int) {
    for {
        row = rand.Intn(m.rows)
        col = rand.Intn(m.cols)
        if location := m.bits[row][col]; location {
            return
        }
    }
}

func (m *Mask) SetBit(row, col int, switchOn bool) {
    m.bits[row][col] = switchOn
}


func MaskFromTxt(filename string) *Mask {
    lines := readlines(filename)

    rows := len(lines)
    cols := len(lines[0])
    mask := NewMask(rows, cols)

    for row := range mask.bits {
        for col := range mask.bits[0] {
            if string(lines[row][col]) == "x" {
                mask.bits[row][col] = false
            } else {
                mask.bits[row][col] = true
            }
        }
    }
    return mask
}

func readlines(filename string) (lines []string) {
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func MaskFromPng(filename string) *Mask {
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }

    img, _, err := image.Decode(f)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    bounds := img.Bounds()
    mask := NewMask(bounds.Max.X, bounds.Max.Y)

    for row := 0; row < mask.rows; row++ {
        for col := 0; col < mask.cols; col++ {
            black := color.RGBA{0, 0, 0, 255}
            if img.At(row, col) == black {
                mask.bits[row][col] = false
            } else {
                mask.bits[row][col] = true
            }
        }
    }
    return mask
}
