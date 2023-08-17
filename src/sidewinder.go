package main

import "math/rand"

type Sidewinder struct {}

func NewSiderwinder() *Sidewinder {
    return &Sidewinder{}
}

func (sw *Sidewinder) on(grid *Grid) {
    for row := range grid.grid {
        var run []*Cell

        for col := range grid.grid[0] {
            cell := grid.grid[row][col]
            run = append(run, cell)

            atEasternBoundary := (cell.east == nil)
            atNorthernBoundary := (cell.north == nil)
            shouldCloseOut := atEasternBoundary ||
                (!atNorthernBoundary && rand.Intn(2) == 0)

            if shouldCloseOut {
                member := run[rand.Intn(len(run))]
                if member.north != nil {
                    member.Link(member.north, true)
                }
                run = nil
            } else {
                cell.Link(cell.east, true)
            }
        }
    }
}
