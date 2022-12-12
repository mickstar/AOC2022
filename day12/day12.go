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
	Weight uint
}
type Node struct {
	X      uint
	Y      uint
	Height uint8
	Edges  []Edge
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

			if x-1 >= 0 {
				weight := calculateWeight(node.Height, tree[y][x-1].Height)
				if weight == 1 {
					node.Edges = append(node.Edges, Edge{
						Node:   tree[y][x-1],
						Weight: uint(weight),
					})
				}
			}
			if x+1 < len(tree[0]) {
				weight := calculateWeight(node.Height, tree[y][x+1].Height)
				if weight == 1 {
					node.Edges = append(node.Edges, Edge{
						Node:   tree[y][x+1],
						Weight: uint(weight),
					})
				}
			}
			if y-1 >= 0 {
				weight := calculateWeight(node.Height, tree[y-1][x].Height)
				if weight == 1 {
					node.Edges = append(node.Edges, Edge{
						Node:   tree[y-1][x],
						Weight: uint(weight),
					})
				}
			}
			if y+1 < len(tree) {
				weight := calculateWeight(node.Height, tree[y+1][x].Height)
				if weight == 1 {
					node.Edges = append(node.Edges, Edge{
						Node:   tree[y+1][x],
						Weight: uint(weight),
					})
				}
			}
		}
	}

	return
}

func calculateWeight(a, b uint8) int {
	if a == 'z' && b == 'E' {
		return 1
	} else if a == 'S' && b == 'a' {
		return 1
	} else if a == 'S' || b == 'S' || a == 'E' || b == 'E' {
		return -1
	}

	val := int(b) - int(a)
	if val == 0 {
		return 1
	}
	return val
}

// using Djikstras algo,
// iterate over list of possible paths, take one then assess
func shortestPath(tree [][]*Node, start *Node, dest *Node) int {
	visited := make(map[*Node]bool)
	distanceFromStart := make(map[*Node]int)

	for y := 0; y < len(tree); y++ {
		for x := 0; x < len(tree[0]); x++ {
			distanceFromStart[tree[y][x]] = -1
		}
	}
	distanceFromStart[start] = 0

	curNode := start
	i := 0
	for curNode != nil {
		i++
		fmt.Println("at node", curNode, "distance From start is ", distanceFromStart[curNode])
		for _, edge := range curNode.Edges {
			distanceForEdge := distanceFromStart[curNode] + int(edge.Weight)
			if distanceFromStart[edge.Node] == -1 || distanceFromStart[edge.Node] > distanceForEdge {
				distanceFromStart[edge.Node] = distanceForEdge
			}

			if edge.Node == dest {
				return distanceFromStart[edge.Node]
			}
		}
		visited[curNode] = true
		curNode = getNextNodeToVisit(visited, distanceFromStart)
		if curNode == nil {
			fmt.Println("finished with ", len(visited), "visited and distances", distanceFromStart)
		}
	}

	return distanceFromStart[dest]
}

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

	path := shortestPath(tree, start, dest)
	fmt.Println("Shortest Path is ", path)
}
