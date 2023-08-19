package main

import "fmt"

func main() {
	rows, cols := 10, 10
	grid := NewGrid(rows, cols)

	fmt.Println()

	ws := NewWilsons()
	ws.on(grid)

	// start := grid.grid[0][0]
	// distances := start.Distances()
	// grid.distances = distances.PathTo(grid.grid[rows-1][cols-1])

	fmt.Println(grid.ToString())
}

func coloredMazes(rows, cols int) {
	ab := NewAldousBroder()

	for i := 0; i < 5; i++ {
		grid := NewGrid(rows, cols)
		ab.on(grid)

		middle := grid.grid[rows/2][cols/2]
		grid.distances = middle.Distances()

		filename := fmt.Sprintf("mazes/ab_%02d.png", i+1)
		grid.ToPng(true, 80, filename)
	}
}
