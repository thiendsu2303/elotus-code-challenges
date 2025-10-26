package main

import "fmt"

var ans, cnt []int
var graph [][]int

func sumOfDistancesInTree(n int, edges [][]int) []int {
	ans = make([]int, n)
	cnt = make([]int, n)
	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph = append(graph, []int{})
	}

	for _, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = make([]int, 0)
		}
		if graph[edge[1]] == nil {
			graph[edge[1]] = make([]int, 0)
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	dfs(0, -1)
	dfs2(0, -1)

	return ans
}

func dfs(node, parent int) {
	for _, i := range graph[node] {
		if i == parent {
			continue
		}
		dfs(i, node)
		cnt[node] += cnt[i]
		ans[node] += ans[i] + cnt[i]
	}
	cnt[node]++
}

func dfs2(node, parent int) {
	for _, i := range graph[node] {
		if i == parent {
			continue
		}
		ans[i] = ans[node] - cnt[i] + (len(cnt) - cnt[i])
		dfs2(i, node)
	}
}

func main() {
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	fmt.Println(sumOfDistancesInTree(n, edges))
}
