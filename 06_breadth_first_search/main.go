package main

import "fmt"

var graph = map[string][]string{
	"you":    {"alice", "bob", "claire"},
	"bob":    {"anuj", "peggy"},
	"alice":  {"you", "peggy", "alice"},
	"claire": {"jonny", "thom", "bob"},
	"anuj":   {},
	"peggy":  {"you"},
	"thom":   {},
	"jonny":  {"peggy", "peggy"},
	"jim":    {"thom"},
}

func personIsSeller(name string) bool {
	return len(name) > 0 && name[len(name)-1] == 'x'
}

// TIme Complexity O(v+e)
func bfs(graph map[string][]string, name string) string {
	searchQueue := graph[name]
	visited := map[string]bool{name: true}

	for len(searchQueue) > 0 {
		person := searchQueue[0]

		if _, exist := visited[person]; !exist {
			if personIsSeller(person) {
				return person
			} else {
				searchQueue = append(searchQueue, graph[person]...)
				visited[person] = true
			}
		}
		if len(searchQueue) > 1 {
			searchQueue = searchQueue[1:]
		} else {
			return ""
		}
	}
	return ""
}

func main() {
	r := bfs(graph, "you")
	// r = bfs(graph, "jonny")

	if r != "" {
		fmt.Printf("found seller: %v\n", r)
	} else {
		fmt.Println("There is no seller.")
	}
}
