package graph

import (
	"fmt"
)

var (
	vertices = []string{"0", "1", "2", "3", "4"}

	edges = [][]string{{"0", "1"}, {"0", "4"}, {"1", "4"}, {"1", "3"}, {"1", "2"}, {"2", "3"}, {"3", "4"}}
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

// node represent the
type node struct {
	value string
	next  *node
}

// AdjacencyList represent the graph representation in the adjacency list format.
func AdjacencyList(vertex []string, edge [][]string) {
	var adjGraph = make([]*node, len(vertex))

	// verticesPoint represent the map of vertex and its position.
	var verticesPoint = make(map[string]int)

	// populate vertices point
	for i := range vertex {
		verticesPoint[vertex[i]] = i
	}

	for i := range edge {
		// edge[i][0]
		p1 := verticesPoint[edge[i][0]]
		p2 := verticesPoint[edge[i][1]]

		// p1 -> p2
		if adjGraph[p1] == nil {
			adjGraph[p1] = &node{value: edge[i][1]}
		} else {
			freeNode(adjGraph[p1], &node{value: edge[i][1]})
		}

		// if adjGraph[p1] == nil {
		// 	adjGraph[p1] = &node{value: edge[i][1]}
		// } else {
		// 	p1Node := adjGraph[p1]

		// 	for p1Node != nil && p1Node.next != nil {
		// 		p1Node = p1Node.next
		// 	}

		// 	p1Node.next = &node{value: edge[i][1]}
		// }

		// p2 -> p1
		if adjGraph[p2] == nil {
			adjGraph[p2] = &node{value: edge[i][0]}
		} else {
			freeNode(adjGraph[p2], &node{value: edge[i][0]})
		}
		// if adjGraph[p2] == nil {
		// 	adjGraph[p2] = &node{value: edge[i][0]}
		// } else {
		// 	p2Node := adjGraph[p2]
		// 	for p2Node != nil && p2Node.next != nil {
		// 		p2Node = p2Node.next
		// 	}
		// 	p2Node.next = &node{value: edge[i][0]}
		// }

	}
	displayAdacencyList(adjGraph)
}

// displayMatrix is the helper function to view 2d format.
func displayAdacencyList(arr []*node) {

	for i := range arr {
		if arr[i] == nil {
			continue
		}

		dispNode(arr[i])
		// for nd.next != nil {
		// 	fmt.Print(" ->", nd.value)
		// 	nd = nd.next
		// }

		fmt.Println()
	}
}

func freeNode(nd *node, nv *node) {

	if nd == nil {
		return
	}
	if nd.next == nil {
		nd.next = nv
		return
	}

	freeNode(nd.next, nv)
}

func dispNode(nd *node) {

	if nd == nil {
		return
	}

	fmt.Print(" ->", nd.value)

	dispNode(nd.next)
}
