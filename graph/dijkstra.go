package graph

import (
	"golang.org/x/exp/slices"
)

var processed []string

// 计算加权最短路径
// 狄克斯特拉算法只适用于有向无环图 DAG directed acyclic graph
// 不适用包含负权重的图
func dijkstraSearch(graph map[string]map[string]int, costs map[string]int, parent *map[string]string) {
	node := findLowestCostNode(costs)
	for {
		if node != "" {
			cost := costs[node]
			neighbors := graph[node]
			// 计算当前node到它所有邻居节点node的cost
			for key := range neighbors {
				newCost := cost + neighbors[key]
				if costs[key] > newCost {
					costs[key] = newCost
					(*parent)[key] = node
				}
			}
			processed = append(processed, node)
			node = findLowestCostNode(costs)
		} else {
			break
		}
	}
}

// 查找花费最少的节点
func findLowestCostNode(costs map[string]int) string {
	lowestCost := (int(^uint(0) >> 1))
	var lowestCostNode string
	for k, v := range costs {
		if v < lowestCost && !slices.Contains(processed, k) {
			lowestCost = v
			lowestCostNode = k
		}
	}
	return lowestCostNode
}
