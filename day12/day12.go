package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	s := string(bytes)
	return strings.Split(s, "\n"), nil
}

type Edge struct {
	Node   *Node
	Weight int
}
type Node struct {
	X      uint
	Y      uint
	Height uint8
	Edges  []Edge
}

func (n *Node) String() string {
	var edgeString = "["
	for _, edge := range n.Edges {
		edgeString += fmt.Sprintf("Node(%d,%d)[%s], weight=%d", edge.Node.X, edge.Node.Y, string(edge.Node.Height), edge.Weight)
		edgeString += ","
	}
	edgeString += "]"

	return fmt.Sprintf("Node (%d,%d)[%s] with Edges %s", n.X, n.Y, string(n.Height), edgeString)
}

func (e *Edge) String() string {
	return fmt.Sprintf("Edge[%s][%d]", e.Node, e.Weight)
}

// returns the head
func readGraph(data []string) (tree [][]*Node, start *Node, destination *Node) {
	tree = make([][]*Node, len(data))

	start = nil
	destination = nil

	for i, line := range data {
		tree[i] = make([]*Node, len(line))
		for j, c := range line {
			tree[i][j] = &Node{
				X:      uint(j),
				Y:      uint(i),
				Height: uint8(c),
				Edges:  make([]Edge, 0),
			}
			if c == 'S' {
				start = tree[i][j]
			} else if c == 'E' {
				destination = tree[i][j]
			}
		}
	}

	// now we have to link everything up.
	for y := 0; y < len(tree); y++ {
		for x := 0; x < len(tree[0]); x++ {
			node := tree[y][x]
			// we have to check around this node for things that we can move to

			directions := []struct {
				x int
				y int
			}{
				{0, -1}, // down
				{0, 1},  // up
				{-1, 0}, //left
				{1, 0},  // right
			}

			for _, dir := range directions {
				if y+dir.y >= 0 && y+dir.y < len(tree) && x+dir.x >= 0 && x+dir.x < len(tree[0]) {
					weight := calculateWeight(node.Height, tree[y+dir.y][x+dir.x].Height)
					if canMakeEdge(weight) {
						node.Edges = append(node.Edges, Edge{
							Node:   tree[y+dir.y][x+dir.x],
							Weight: weight,
						})
					}
				}
			}
		}
	}

	return
}

func canMakeEdge(weight int) bool {
	// allow going up 1 or down any amount
	//return weight <= 1

	//return weight == 0 || weight == -1 || weight <= 0
	return weight >= -1
}

func calculateWeight(a, b uint8) int {
	//if (a == 'z' || a == 'y') && b == 'E' {
	//	return 1
	//} else if a == 'S' && (b == 'a' || b == 'b') {
	//	return 1
	//}

	if a == 'E' && (b == 'z' || b == 'y') {
		return -1
	}

	if a == 'S' || b == 'S' || a == 'E' || b == 'E' {
		return 100000
	}

	val := int(b) - int(a)
	return val
}

// using Djikstras algo,
// iterate over list of possible paths, take one then assess
func shortestPath(tree [][]*Node, start *Node, targetHeight uint8) int {
	visited := make(map[*Node]bool)
	distanceFromStart := make(map[*Node]int)

	for y := 0; y < len(tree); y++ {
		for x := 0; x < len(tree[0]); x++ {
			distanceFromStart[tree[y][x]] = -1
		}
	}
	distanceFromStart[start] = 0

	curNode := start
	for curNode != nil {
		fmt.Println("at node", curNode, "distance From start is ", distanceFromStart[curNode])
		for _, edge := range curNode.Edges {
			distanceForEdge := distanceFromStart[curNode] + 1
			if distanceFromStart[edge.Node] == -1 || distanceForEdge < distanceFromStart[edge.Node] {
				distanceFromStart[edge.Node] = distanceForEdge
			}
		}
		visited[curNode] = true
		curNode = getNextNodeToVisit(visited, distanceFromStart)
	}

	minA := -1
	for node, value := range distanceFromStart {
		if node.Height == 'a' && value != -1 {
			if minA == -1 || value < minA {
				minA = value
			}
		}
	}
	return minA
}

// returns the next node to visit
// returns nil if there are no more nodes to visit
// from the list of unvisited nodes, we take one with the smallest distance from start
func getNextNodeToVisit(visited map[*Node]bool, distances map[*Node]int) *Node {
	var minNode *Node = nil
	min := -1
	for node, dis := range distances {
		if visited[node] || dis == -1 {
			continue
		}
		if min == -1 || dis < min {
			min = dis
			minNode = node
		}
	}
	return minNode
}

func main() {
	data, err := readFile("day12/day12_input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	tree, start, dest := readGraph(data)

	fmt.Println("start", start)
	fmt.Println("dest", dest)

	for y := 0; y < len(tree); y++ {
		for x := 0; x < len(tree[0]); x++ {
			if len(tree[y][x].Edges) == 0 {
				fmt.Println("No edges for this node", tree[y][x])
			}
		}
	}

	////path := shortestPath(tree, start, dest)
	//path := shortestPath(tree, start, 'E')
	//fmt.Println("Part 1: Shortest Path is ", path)

	path := shortestPath(tree, dest, 'a')
	fmt.Println("Part 2: Shortest Path is ", path)
}
