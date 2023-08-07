package main

import "math/rand"

type Sidewinder struct {}

func NewSiderwinder() (sw *Sidewinder) {
    sw = &Sidewinder{}
    return
}

func (sw *Sidewinder) on(grid *Grid) {
    for row := range grid.grid {
        var run []*Cell

        for col := range grid.grid[0] {
            cell := grid.grid[row][col]
            run = append(run, cell)

            atEasternBoundary := (cell.East == nil)
            atNorthernBoundary := (cell.North == nil)
            shouldCloseOut := atEasternBoundary ||
                              (!atNorthernBoundary && rand.Intn(2) == 0)

            if shouldCloseOut {
                member := run[rand.Intn(len(run))]
                if member.North != nil {
                    member.Link(member.North, true)
                }
                run = nil
            } else {
                cell.Link(cell.East, true)
            }
        }
    }
}
