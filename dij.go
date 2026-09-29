package main

import (
	"fmt"
	"math"
)

type Edge struct {
	to     string
	weight int
}

func buildGraph() map[string][]Edge {
	graph := make(map[string][]Edge)

	addEdge := func(from, to string, weight int) {
		graph[from] = append(graph[from], Edge{to, weight})
		graph[to] = append(graph[to], Edge{from, weight})
	}

	addEdge("A", "B", 2)
	addEdge("A", "F", 1)
	addEdge("B", "D", 2)
	addEdge("D", "F", 3)
	addEdge("B", "C", 2)
	addEdge("C", "E", 3)
	addEdge("E", "G", 7)
	addEdge("F", "G", 5)
	addEdge("D", "E", 4)
	addEdge("B", "E", 4)
	addEdge("C", "Z", 1)
	addEdge("G", "Z", 6)

	return graph
}

func dijkstra(graph map[string][]Edge, start string, end string) ([]string, int) {
	dist := make(map[string]int)
	previous := make(map[string]string)
	visited := make(map[string]bool)

	for vertex := range graph {
		dist[vertex] = math.MaxInt
	}

	dist[start] = 0

	for {
		current := ""
		bestDistance := math.MaxInt

		for vertex, distance := range dist {
			if !visited[vertex] && distance < bestDistance {
				bestDistance = distance
				current = vertex
			}
		}

		if current == "" {
			break
		}

		visited[current] = true

		fmt.Printf("Visit %s -> distance = %d\n", current, dist[current])

		for _, edge := range graph[current] {
			newDistance := dist[current] + edge.weight

			if newDistance < dist[edge.to] {
				dist[edge.to] = newDistance
				previous[edge.to] = current

				fmt.Printf("  Update %s: %d via %s\n",
					edge.to, newDistance, current)
			}
		}

		if current == end {
			break
		}
	}

	path := []string{}
	current := end

	if dist[end] == math.MaxInt {
		return path, -1
	}

	for current != "" {
		path = append([]string{current}, path...)
		if current == start {
			break
		}
		current = previous[current]
	}

	return path, dist[end]
}

func main() {
	graph := buildGraph()

	start := "A"
	end := "Z"

	fmt.Println("===================================")
	fmt.Println("       DIJKSTRA'S ALGORITHM")
	fmt.Println("====================================")
	fmt.Printf("Find the shortest route from %s to %s\n\n", start, end)

	path, distance := dijkstra(graph, start, end)

	fmt.Println("\n====================================")
	fmt.Println("RESULT")
	fmt.Println("====================================")
	fmt.Printf("Shortest path: %v\n", path)
	fmt.Printf("Total cost: %d\n", distance)
	fmt.Println("====================================")
}
