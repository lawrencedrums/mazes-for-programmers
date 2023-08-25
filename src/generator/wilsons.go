package generator

import (
	"math/rand"
	"slices"

    "mazes/grid"
    c "mazes/grid/cell"
)

func Wilsons(grid *grid.Grid) {
    unvisited := make([]*c.Cell, grid.Size())
    for cell := range grid.Cells() {
        unvisited = append(unvisited, cell)
    }

	firstIdx := rand.Intn(len(unvisited))
	unvisited = slices.Delete(unvisited, firstIdx, firstIdx+1)

	for len(unvisited) > 0 {
		cell := unvisited[rand.Intn(len(unvisited))]

		var path []*c.Cell
		path = append(path, cell)

		for slices.Contains(unvisited, cell) {
			cell = cell.RandomNeighbor()

			position := -1
			for i, cellInPath := range path {
				if cellInPath == cell {
					position = i
					break
				}
			}

			if position > -1 { // if formed a loop
				path = slices.Delete(path, position+1, len(path))
			} else {
				path = append(path, cell)
			}
		}

		for i := 0; i < len(path)-1; i++ {
			path[i].Link(path[i+1], true)
			unvisited = remove(unvisited, path[i])
		}
	}
}

func remove(path []*c.Cell, target *c.Cell) []*c.Cell{
	for i, cell := range path {
		if cell == target {
			return slices.Delete(path, i, i+1)
		}
	}
    return path
}
