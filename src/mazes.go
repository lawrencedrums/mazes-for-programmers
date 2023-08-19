package main

import "fmt"

type algo func(*Grid)

func main() {
	rows, cols := 25, 25
	grid := NewGrid(rows, cols)

	fmt.Println()

	algo := NewHuntAndKill()
	algo.on(grid)

	// start := grid.grid[0][0]
	// distances := start.Distances()
	// grid.distances = distances.PathTo(grid.grid[rows-1][cols-1])

	fmt.Println(grid.ToString())
    coloredMazes(algo.on, rows, cols)
}

func coloredMazes(algo func(*Grid), rows, cols int) {
	for i := 0; i < 5; i++ {
		grid := NewGrid(rows, cols)
		algo(grid)

		middle := grid.grid[rows/2][cols/2]
		grid.distances = middle.Distances()

		filename := fmt.Sprintf("mazes/%02d.png", i+1)
		grid.ToPng(true, 80, filename)
	}
}
