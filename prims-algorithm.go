package main

import (
	"fmt"
	"math/rand"
	// "math"
	"time"
)

type node struct {
	x         int
	y         int
	character string
	parent    *node
}

const wall_character = "█"
const maze_character = " "
const width = 51
const height = 51

var nodes []node = initGrid(width, height)
var walls []*node = make([]*node, 0)

func main() {
	rand.Seed(time.Now().Unix())
	startingCell := findRandomCell()

	explore(startingCell)

	for len(walls) != 0 {
		wall, index := findRandomWall()

		x, y := findNextCellCoords(wall, wall.parent)

		if isEdge(x, y) {
			removeWall(wall, index)

			continue
		}

		nextCell := locateNode(x, y)

		if isWall(nextCell) {
			markAsPassage(wall)
			explore(nextCell)
		}

		removeWall(wall, index)
	}

	draw(nodes)
}

func isWall(cell *node) bool {
	return cell.character == wall_character
}

func isEdge(x int, y int) bool {
	return x <= 0 || x >= width-1 || y <= 0 || y >= height-1
}

func findRandomWall() (*node, int) {
	index := rand.Intn(len(walls))

	return walls[index], index
}

func findRandomCell() *node {
	index := rand.Intn(len(nodes))

	return &nodes[index]
}

func markAsPassage(node *node) {
	node.character = maze_character
}

func findNextCellCoords(wall *node, cell *node) (int, int) {
	x := wall.x - cell.x + wall.x
	y := wall.y - cell.y + wall.y

	return x, y
}

func removeWall(wall *node, index int) {
	walls = append(walls[:index], walls[index+1:]...)
}

func explore(cell *node) {
	markAsPassage(cell)
	addWalls(*cell)
}

func addWalls(cell node) {
	if cell.x < width-1 {
		wall := locateNode(cell.x+1, cell.y)

		addWall(cell, wall)
	}

	if cell.x > 0 {
		wall := locateNode(cell.x-1, cell.y)

		addWall(cell, wall)
	}

	if cell.y < height-1 {
		wall := locateNode(cell.x, cell.y+1)

		addWall(cell, wall)
	}

	if cell.y > 0 {
		wall := locateNode(cell.x, cell.y-1)

		addWall(cell, wall)
	}
}

func addWall(cell node, wall *node) {
	wall.parent = &cell
	walls = append(walls, wall)
}

func locateNode(x int, y int) *node {
	return &nodes[x+y*width]
}

func initGrid(width int, height int) []node {
	grid := make([]node, 0, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid = append(grid, node{x, y, wall_character, nil})
		}
	}

	return grid
}

func draw(nodes []node) {
	for _, node := range nodes {
		if node.x == 0 && node.y != 0 {
			fmt.Print("\n")
		}

		fmt.Print(node.character)
	}

	fmt.Print("\n")
}
