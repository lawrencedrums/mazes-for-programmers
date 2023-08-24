package main

import "fmt"

func main() {
	rows, cols := 20, 20
	grid := NewGrid(rows, cols)

	algo := NewRecursiveBacktracker()
	algo.on(grid)

	// start := grid.grid[0][0]
	// distances := start.Distances()
	// grid.distances = distances.PathTo(grid.grid[rows-1][cols-1])

	// fmt.Println(grid.ToString())

    // coloredMazes(algo.on, rows, cols)
    compareDeadends(rows, cols)

    fmt.Println("done")
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

func compareDeadends(rows, cols int) {
    algos := map[string]func(*Grid){
        "binary_tree": NewBinaryTree().on,
        "siderwinder": NewSiderwinder().on,
        "aldous_broder": NewAldousBroder().on,
        "wilsons": NewWilsons().on,
        "hunt_and_kill": NewHuntAndKill().on,
        "recursive_backtracker": NewRecursiveBacktracker().on,
    }

    tries := 100
    averages := make(map[string]int)
    for name, algo := range algos {
        fmt.Println("Running", name)

        totalDeadends := 0
        deadendCounts := make([]int, tries)
        for i := 0; i < tries; i++ {
            grid := NewGrid(rows, cols)
            algo(grid)

            deadendCounts[i] = len(grid.DeadEnds())
            totalDeadends += len(grid.DeadEnds())
        }
        averages[name] = totalDeadends / len(deadendCounts)
    }

    var sortedAlgos []string
    var sortedDeadends []int
    for len(averages) > 0 {
        var minAlgo string
        minDeadend := rows * cols
        for k, v := range averages {
            if v < minDeadend {
                minAlgo = k
                minDeadend = v
            }
        }
        sortedAlgos = append(sortedAlgos, minAlgo)
        sortedDeadends = append(sortedDeadends, minDeadend)
        delete(averages, minAlgo)
    }

    fmt.Println("Average deadends per maze:")
    for i, algo := range sortedAlgos {
        deadends := sortedDeadends[i]
        percentage := float64(deadends) * 100.0 / float64(rows * cols)
        fmt.Printf("%25s: %3d/%d %.2f%%\n", algo, deadends, rows*cols, percentage)
    }
}
