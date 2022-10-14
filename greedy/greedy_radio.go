package greedy

import (
	mapset "github.com/deckarep/golang-set/v2"
)

func greedyRadio(statesNeeded mapset.Set[string], stations map[string]mapset.Set[string]) []string {
	var finalStations []string

	for {
		if statesNeeded.Cardinality() > 0 {
			var bestStation string
			statesCovered := mapset.NewSet[string]()
			for station, states := range stations {
				covered := statesNeeded.Intersect(states)
				if covered.Cardinality() > statesCovered.Cardinality() {
					bestStation = station
					statesCovered = covered
				}
			}
			statesNeeded = statesNeeded.Difference(statesCovered)
			finalStations = append(finalStations, bestStation)
		} else {
			break
		}

	}
	return finalStations
}
