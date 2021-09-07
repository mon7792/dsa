package graph

import "fmt"

var (
	vertices = []string{"0", "1", "2", "3", "4"}

	edges = [][]string{{"0", "1"}, {"0", "4"}, {"1", "4"}, {"1", "3"}, {"1", "2"}, {"2", "3"}, {"3", "4"}}
)

// graph representation by adjacency matrix
func AdjacencyMatrix() {

	var matrix = make([][]int, len(vertices))
	fmt.Println(matrix)
}
