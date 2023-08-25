package generator

import (
    "math/rand"

    "mazes/grid"
    "mazes/grid/cell"
)

func Sidewinder(grid grid.Grider) {
    for row := range grid.Rows() {
        var run []*cell.Cell

        for _, cell := range row {
            run = append(run, cell)

            atEasternBoundary := (cell.East() == nil)
            atNorthernBoundary := (cell.North() == nil)
            shouldCloseOut := atEasternBoundary ||
                (!atNorthernBoundary && rand.Intn(2) == 0)

            if shouldCloseOut {
                member := run[rand.Intn(len(run))]
                if member.North() != nil {
                    member.Link(member.North(), true)
                }
                run = nil
            } else {
                cell.Link(cell.East(), true)
            }
        }
    }
}
