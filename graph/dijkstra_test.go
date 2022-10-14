package graph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {

	graph := make(map[string]map[string]int)

	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]int)
	graph["a"]["end"] = 1

	graph["b"] = make(map[string]int)
	graph["b"]["a"] = 3
	graph["b"]["end"] = 5

	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	const INT_MAX = (int(^uint(0) >> 1))
	costs["end"] = INT_MAX

	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["end"] = ""

	dijkstraSearch(graph, costs, &parents)
	want := map[string]string{"a": "b", "b": "start", "end": "a"}
	if !reflect.DeepEqual(parents, want) {
		fmt.Printf("dijkstra can't find shortest path want %#v, got %#v\n", want, parents)
	}

	reverse := make(map[string]string)
	var path string

	weights := 0
	for key := range parents {
		weights += graph[parents[key]][key]
		reverse[parents[key]] = key
	}
	key := "start"
	for {
		if v, ok := reverse[key]; ok {
			path += key + "=>" + v + " "
			key = v
		} else {
			break
		}
	}
	fmt.Printf("path: %#v\n", path)
	fmt.Printf("weights: %d\n", weights)

}
