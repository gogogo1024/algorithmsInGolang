package graph

import (
	"testing"
)

func TestBreadthFirst(t *testing.T) {

	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	found := breadthFirst(graph, "you", "thom")
	wanted := true
	if found != wanted {
		t.Errorf("breadthFirst on '%#v'; want '%v', got '%v';", graph, wanted, found)
	}
}
