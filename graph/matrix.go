package graph

import (
	"fmt"
)

var (
	vertices = []string{"A", "B", "C", "D", "E"}

	edges = [][]string{{"A", "B"}, {"A", "E"}, {"B", "E"}, {"B", "D"}, {"B", "C"}, {"C", "D"}, {"D", "E"}}
)

// graph representation by adjacency matrix
func AdjacencyMatrix(vertex []string, edge [][]string) {

	var (
		//  matrix is the 2D adjacency array.
		matrix = make([][]int, len(vertex), len(vertex))

		// verticesPoint represent the map of vertex and its position.
		verticesPoint = make(map[string]string)

		// edgePoint represent the edges in the map format.
		edgePoint = make(map[string]bool)
	)

	//  2d array initialization.
	for i := range vertex {
		arr := make([]int, len(vertex))
		matrix[i] = arr
	}

	// populate vertices point
	for i := range vertex {
		verticesPoint[vertex[i]] = fmt.Sprintf("%d", i)
	}

	// poulate edge points
	for i := range edge {
		p1 := verticesPoint[edge[i][0]]
		p2 := verticesPoint[edge[i][1]]
		// birectional edge
		edgePoint[p1+p2] = true
		edgePoint[p2+p1] = true
	}

	// loop through matrix
	for i := range matrix {
		for j := range matrix[i] {
			if edgePoint[fmt.Sprintf("%d%d", i, j)] {
				matrix[i][j] = 1
			}
		}
	}

	// display matrix
	displayMatrix(matrix)

}

// displayMatrix is the helper function to view 2d format.
func displayMatrix(arr [][]int) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
