package grid

import "math/rand"

type Mask struct {
    rows, cols int
    bits [][]bool
}

func NewMask(rows, cols int) *Mask {
    bits := make([][]bool, rows)
    for row := range bits {
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
