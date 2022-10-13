package graph

import (
	"fmt"

	"golang.org/x/exp/slices"
)

//	查询name是否有符合search条件的（非加权计算最短路径=>段数）
func breadthFirst(m map[string][]string, name string, search string) bool {
	// 保存已经访问过的节点
	var searched []string
	searchQueue := m[name]
	for i := 0; i < len(searchQueue); i++ {
		v := searchQueue[i]
		if !slices.Contains(searched, v) {
			if v == search {
				fmt.Printf("search %s", search)
				return true
			} else {
				searchQueue = append(searchQueue, m[v]...)
				searched = append(searched, v)
			}
		}
	}
	return false
}
