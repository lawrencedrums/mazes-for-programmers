package util

import (
    "fmt"

    "mazes/generator"
    "mazes/grid"
)

func CompareDeadends(rows, cols int) {
    algos := map[string]func(grid.Grider){
        "binary_tree": generator.BinaryTree,
        "siderwinder": generator.Sidewinder,
        "aldous_broder": generator.AldousBroder,
        "wilsons": generator.Wilsons,
        "hunt_and_kill": generator.HuntAndKill,
        "recursive_backtracker": generator.RecursiveBacktracker,
    }

    tries := 100
    averages := make(map[string]int)
    for name, algo := range algos {
        fmt.Println("Running", name)

        totalDeadends := 0
        deadendCounts := make([]int, tries)
        for i := 0; i < tries; i++ {
            grid := grid.NewGrid(rows, cols)
            algo(grid)

            deadendCounts[i] = len(grid.Deadends())
            totalDeadends += len(grid.Deadends())
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
