package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	from   string
	to     string
	weight int
}

type DSU struct {
	parent map[string]string
	rank   map[string]int
}

func NewDSU(vertices []string) *DSU {
	dsu := &DSU{
		parent: make(map[string]string),
		rank:   make(map[string]int),
	}

	for _, vertex := range vertices {
		dsu.parent[vertex] = vertex
		dsu.rank[vertex] = 0
	}

	return dsu
}

func (d *DSU) Find(x string) string {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b string) bool {
	rootA := d.Find(a)
	rootB := d.Find(b)

	if rootA == rootB {
		return false
	}

	if d.rank[rootA] < d.rank[rootB] {
		d.parent[rootA] = rootB
	} else if d.rank[rootA] > d.rank[rootB] {
		d.parent[rootB] = rootA
	} else {
		d.parent[rootB] = rootA
		d.rank[rootA]++
	}

	return true
}

func kruskal(vertices []string, edges []Edge) ([]Edge, int) {

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	dsu := NewDSU(vertices)
	mst := []Edge{}
	totalWeight := 0

	fmt.Println("Edges sorted by weight:")
	for _, edge := range edges {
		fmt.Printf("%s-%s = %d\n", edge.from, edge.to, edge.weight)
	}

	fmt.Println("\nKruskal's steps:")

	for _, edge := range edges {
		if dsu.Union(edge.from, edge.to) {
			mst = append(mst, edge)
			totalWeight += edge.weight

			fmt.Printf("TAKE   %s-%s = %d\n",
				edge.from, edge.to, edge.weight)
		} else {
			fmt.Printf("SKIP   %s-%s = %d (cycle)\n",
				edge.from, edge.to, edge.weight)
		}

		// For 8 vertices, an MST has 7 edges.
		if len(mst) == len(vertices)-1 {
			break
		}
	}

	return mst, totalWeight
}

func main() {
	vertices := []string{
		"A", "B", "C", "D",
		"E", "F", "G", "Z",
	}

	edges := []Edge{
		{from: "A", to: "B", weight: 2},
		{from: "A", to: "F", weight: 1},
		{from: "B", to: "D", weight: 2},
		{from: "D", to: "F", weight: 3},
		{from: "B", to: "C", weight: 2},
		{from: "C", to: "E", weight: 3},
		{from: "E", to: "G", weight: 7},
		{from: "F", to: "G", weight: 5},
		{from: "D", to: "E", weight: 4},
		{from: "B", to: "E", weight: 4},
		{from: "C", to: "Z", weight: 1},
		{from: "G", to: "Z", weight: 6},
	}

	fmt.Println("====================================")
	fmt.Println("         KRUSKAL'S ALGORITHM")
	fmt.Println("====================================")
	fmt.Println("Finding the Minimum Spanning Tree (MST)\n")

	mst, totalWeight := kruskal(vertices, edges)

	fmt.Println("\n====================================")
	fmt.Println("RESULT - MINIMUM SPANNING TREE")
	fmt.Println("====================================")

	for _, edge := range mst {
		fmt.Printf("%s -- %s  weight = %d\n",
			edge.from, edge.to, edge.weight)
	}

	fmt.Printf("\nNumber of MST edges: %d\n", len(mst))
	fmt.Printf("Total MST weight: %d\n", totalWeight)
	fmt.Println("====================================")
}
