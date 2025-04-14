package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrNonSquareMatrix = errors.New("input graph must be a square adjacency matrix")
	ErrInvalidGraph    = errors.New("invalid graph structure")
)

type Result struct {
	Path        []int
	Duration    time.Duration
	VertexCount int
	EdgeCount   int
}

func FindHamiltonianPath(graph [][]int) (*Result, error) {
	startTime := time.Now()
	n := len(graph)

	if err := validateGraph(graph); err != nil {
		return nil, err
	}

	path := make([]int, n)
	for i := range path {
		path[i] = -1
	}
	path[0] = 0

	result := &Result{
		VertexCount: n,
		EdgeCount:   countEdges(graph),
		Duration:    time.Since(startTime),
	}

	if backtrack(graph, path, 1) {
		result.Path = make([]int, len(path))
		copy(result.Path, path)
	}

	return result, nil
}

func validateGraph(graph [][]int) error {
	n := len(graph)
	if n == 0 {
		return ErrInvalidGraph
	}

	for _, row := range graph {
		if len(row) != n {
			return ErrNonSquareMatrix
		}
		for _, val := range row {
			if val != 0 && val != 1 {
				return ErrInvalidGraph
			}
		}
	}
	return nil
}

func countEdges(graph [][]int) int {
	n, count := len(graph), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func backtrack(graph [][]int, path []int, pos int) bool {
	if pos == len(graph) {
		return true
	}

	for v := 1; v < len(graph); v++ {
		if isValidNext(v, graph, path, pos) {
			path[pos] = v
			if backtrack(graph, path, pos+1) {
				return true
			}
			path[pos] = -1
		}
	}
	return false
}

func isValidNext(vertex int, graph [][]int, path []int, pos int) bool {
	if graph[path[pos-1]][vertex] == 0 {
		return false
	}

	for i := 0; i < pos; i++ {
		if path[i] == vertex {
			return false
		}
	}
	return true
}

func main() {

graph := [][]int{
	{0,1,0,0,1,0,1,0,0,0,1,0,0,0,0,1,0,0,1,0,0,0,0,0,0,1,0,0,0,0},
	{1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,0,1,0,0,0},
	{0,1,0,1,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,1,0,0,0,0,0,1,0,0},
	{0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0,1,0},
	{1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0,1},
	{0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0},
	{1,0,0,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0},
	{0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0},
	{0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0},
	{0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0},
	{1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0},
	{0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0},
	{0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1},
	{0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0},
	{0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0,1},
	{1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0,0},
	{0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0,0},
	{0,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0,0},
	{1,0,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1,0},
	{0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,1},
	{0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0},
	{0,0,1,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0},
	{0,1,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0},
	{0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0},
	{0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0},
	{1,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,1},
	{0,1,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0},
	{0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0},
	{0,0,0,1,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1},
	{0,0,0,0,1,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0},
}

	result, err := FindHamiltonianPath(graph)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Path) == 0 {
		fmt.Println("No Hamiltonian Path found.")
	} else {
		fmt.Printf("Hamiltonian Path found: %v\n", result.Path)
		fmt.Printf("Path discovery took: %v\n", result.Duration)
		fmt.Printf("Graph has %d vertices and %d edges\n", result.VertexCount, result.EdgeCount)
	}
}
