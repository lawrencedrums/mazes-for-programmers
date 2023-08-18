package main

import "math/rand"

type AldousBroder struct {}

func NewAldousBroder() *AldousBroder {
    return &AldousBroder{}
}

func (ab *AldousBroder) on(grid *Grid) {
    cell := grid.RandomCell()
    unvisited := grid.Size() - 1

    for unvisited > 0 {
        neighbors := cell.Neighbors()
        randomNeighbor := neighbors[rand.Intn(len(neighbors))]

        if len(randomNeighbor.Links()) == 0 {
            cell.Link(randomNeighbor, true)
            unvisited -= 1
        }

        cell = randomNeighbor
    }
}
