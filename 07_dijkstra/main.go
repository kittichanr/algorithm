package main

import (
	"fmt"
	"math"
)

type graph = map[string]map[string]int

var g = graph{
	"start": map[string]int{"a": 6, "b": 2},
	"a":     map[string]int{"fin": 1},
	"b":     map[string]int{"a": 3, "fin": 5},
	"fin":   map[string]int{},
}

const endNode = "NONE"

func dijkstra(graph graph, start string) (path []string) {
	costs := initHashTableCosts(graph, start)
	parents := initHashTableParents(graph, start)

	if costs == nil || parents == nil {
		return nil
	}

	processed := map[string]bool{}
	node := findLowestCostNode(costs, processed)
	for node != endNode {
		cost := costs[node]
		neighbors := graph[node]

		for n := range neighbors {
			newCost := cost + neighbors[n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed[node] = true
		node = findLowestCostNode(costs, processed)
	}
	path = calcPathInParents(parents, start)
	return path
}

func initHashTableCosts(graph graph, start string) map[string]int {
	if _, ok := graph[start]; !ok {
		return nil
	}
	costs := g[start]
	costs[start] = 0
	for k := range g {
		if _, ok := costs[k]; !ok && k != start {
			costs[k] = math.MaxInt32
		}
	}
	return costs
}

func initHashTableParents(graph graph, start string) map[string]string {
	if _, ok := graph[start]; !ok {
		return nil
	}
	parents := map[string]string{}
	for k := range g[start] {
		parents[k] = start
	}
	if _, ok := parents["fin"]; !ok {
		parents["fin"] = endNode
	}
	return parents
}

func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	lowestCost := math.MaxInt32
	lowestCostNode := endNode
	for k, v := range costs {
		if _, ok := processed[k]; !ok && v < lowestCost {
			lowestCost = v
			lowestCostNode = k
		}
	}
	return lowestCostNode
}

func calcPathInParents(parents map[string]string, start string) []string {
	if parents["fin"] == endNode {
		return nil
	}
	currentPoint := "fin"
	path := []string{}
	for currentPoint != start {
		for k, v := range parents {
			if k == currentPoint {
				path = append(path, currentPoint)
				currentPoint = v
			}
		}
	}
	if len(path) == 1 {
		return []string{start, path[0]}
	}
	reversed := []string{start}
	for i := len(path) - 1; i >= 0; i-- {
		reversed = append(reversed, path[i])
	}
	return reversed
}

func main() {
	if path := dijkstra(g, "start"); path == nil {
		fmt.Println("no path.")
	} else {
		fmt.Println(path)
	}
}
