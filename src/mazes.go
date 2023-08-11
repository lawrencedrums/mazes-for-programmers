package main

func main() {
    btGrid := NewGrid(20, 20)
    swGrid := NewGrid(20, 20)

    bt := NewBinaryTree()
    bt.on(btGrid)

    // fmt.Println(btGrid.ToString())
    btGridStart := btGrid.grid[10][10]
    btGrid.distances = btGridStart.Distances()

    btGrid.ToPng(true, 80, "mazes/btGrid.png")

    sw := NewSiderwinder()
    sw.on(swGrid)

    swGridStart := swGrid.grid[10][10]
    swGrid.distances = swGridStart.Distances()

    // find longest path
    // newStart, _ := distances.Max()
    // newDistances := newStart.Distances()
    // goal, _ := newDistances.Max()
    // swGrid.distances = newDistances.PathTo(goal)

    // swGrid.distances = distances.PathTo(swGrid.grid[swGrid.rows-1][10])

    // fmt.Println(swGrid.ToString())
    swGrid.ToPng(true, 80, "mazes/swGrid.png")
}
