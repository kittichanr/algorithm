package main

import "fmt"

var statesNeeded = []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

var stations = map[string][]string{
	"kone":   {"id", "nv", "ut"},
	"ktwo":   {"wa", "id", "mt"},
	"kthree": {"or", "nv", "ca"},
	"kfour":  {"nv", "ut"},
	"kfive":  {"ca", "az"},
}

func findStations(statesNeeded []string, stations map[string][]string) []string {
	result := []string{}
	included := map[string]bool{}
	for len(statesNeeded) > 0 {
		bestStation := ""
		statesCovered := []string{}
		for name, states := range stations {
			covered := intersect(statesNeeded, states)
			if len(covered) > len(statesCovered) {
				bestStation = name
				statesCovered = covered
			}
		}
		statesNeeded = subtract(statesNeeded, statesCovered)
		if _, ok := included[bestStation]; !ok && bestStation != "" {
			result = append(result, bestStation)
			included[bestStation] = true
		}
	}
	return result
}

func intersect(a []string, b []string) []string {
	r := []string{}
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				r = append(r, aa)
			}
		}
	}
	return r
}

func subtract(a []string, b []string) []string {
	r := []string{}
	for _, aa := range a {
		exist := false
		for _, bb := range b {
			if aa == bb {
				exist = true
				break
			}
		}
		if !exist {
			r = append(r, aa)
		}
	}
	return r
}

func main() {
	// set covering
	s := findStations(statesNeeded, stations)
	fmt.Println(s)
}
