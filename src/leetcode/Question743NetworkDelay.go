package main

import (
	"math"
)

/**
https://leetcode.com/problems/network-delay-time/
dijstra algorithm
*/
func networkDelayTimeV1(times [][]int, N int, K int) int {
	set := make(map[int]int)
	grid := make([][]int, len(times))
	copy(grid, times)
	set[K] = 0
	for N != len(set) && len(grid)+1 != len(set) {
		flag := true
		change := false
		for node, _ := range set {
			next := -1
			index := math.MaxInt32
			weight := math.MaxInt32
			for i := 0; i < len(grid); i++ {
				if node == grid[i][0] {
					if weight > grid[i][2] {
						next = grid[i][1]
						weight = grid[i][2]
						index = i
					}
				}
			}
			if next == -1 {
				flag = false
				continue
			}
			grid[index][2] = math.MaxInt32
			if v, ok := set[next]; ok {
				if v > set[node]+weight {
					set[next] = set[node] + weight
					change = true
				}
			} else {
				set[next] = set[node] + weight
				change = true
			}
		}
		if !flag && !change {
			break
		}
	}
	if N != len(set) {
		return -1
	}
	result := 0
	for _, v := range set {
		if v == math.MaxInt32 {
			return -1
		}
		if result < v {
			result = v
		}
	}
	return result
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func networkDelayTime(times [][]int, N int, K int) int {
	dist := make(map[int]int)
	graph := make(map[int][][]int)
	for i := 0; i < len(times); i++ {
		edge := []int{times[i][1], times[i][2]}
		var edges [][]int
		if v, ok := graph[times[i][0]]; ok {
			edges = v
		} else {
			edges = make([][]int, 0)
		}
		edges = append(edges, edge)
		graph[times[i][0]] = edges
	}
	for i := 1; i <= N; i++ {
		dist[i] = math.MaxInt32
	}
	dist[K] = 0
	seen := make([]bool, N+1)
	for true {
		candNode := -1
		candDist := math.MaxInt32
		for i := 1; i <= N; i++ {
			if !seen[i] && dist[i] < candDist {
				candDist = dist[i]
				candNode = i
			}
		}
		if candNode < 0 {
			break
		}
		seen[candNode] = true
		if edges, ok := graph[candNode]; ok {
			for j := 0; j < len(edges); j++ {
				info := edges[j]
				dist[info[0]] = min(dist[info[0]], dist[candNode]+info[1])
			}
		}
	}
	ans := 0
	for _, v := range dist {
		if v == math.MaxInt32 {
			return -1
		}
		ans = max(ans, v)
	}
	return ans
}

func main() {
	println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
	//println(networkDelayTime([][]int{{2,1,1},{1,2,3}},2,2))
	//println(networkDelayTime([][]int{{1,2,1},{2,3,2},{1,3,4}},3,1))
	//println(networkDelayTime([][]int{{1,2,1}},2,2))
}
