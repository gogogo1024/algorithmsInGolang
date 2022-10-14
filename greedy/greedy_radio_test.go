package greedy

import (
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

func TestGreedyRadio(t *testing.T) {
	statesNeeded := mapset.NewSet[string]()
	statesNeeded.Add("mt")
	statesNeeded.Add("wa")
	statesNeeded.Add("or")
	statesNeeded.Add("id")
	statesNeeded.Add("nv")
	statesNeeded.Add("ut")
	statesNeeded.Add("ca")
	statesNeeded.Add("az")

	stations := make(map[string]mapset.Set[string])
	tmp1 := mapset.NewSet[string]()
	tmp1.Add("id")
	tmp1.Add("nv")
	tmp1.Add("ut")
	stations["one"] = tmp1

	tmp2 := mapset.NewSet[string]()

	tmp2.Add("wa")
	tmp2.Add("id")
	tmp2.Add("mt")
	stations["two"] = tmp2

	tmp3 := mapset.NewSet[string]()
	tmp3.Add("or")
	tmp3.Add("nv")
	tmp3.Add("ca")
	stations["three"] = tmp3

	tmp4 := mapset.NewSet[string]()
	tmp4.Add("nv")
	tmp4.Add("ut")
	stations["four"] = tmp4

	tmp5 := mapset.NewSet[string]()
	tmp5.Add("ca")
	tmp5.Add("az")
	stations["five"] = tmp5

	finalStations := greedyRadio(statesNeeded, stations)
	fmt.Printf("finalStations : %s", finalStations)

}
