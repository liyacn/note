package algorithm

/*
在无向图中，如果任意两个顶点之间都存在路径相连，则称这些顶点属于同一个连通分量。
整个无向图可能由一个或多个互不连通的子图组成，每个这样的子图就是一个连通分量。

深度优先搜索（DFS）或广度优先搜索（BFS）：
	初始化所有节点为未访问状态。
	对于每一个节点，如果它未被访问：开始一次DFS/BFS，
	标记所有能访问到的节点为已访问，并认为它们属于同一个连通分量。
	连通分量数+1。
时间、空间复杂度均为O(n+m)，n为顶点数，m为边数。

并查集（Union-Find）：
	初始时，每个节点是自己的根节点。
	遍历每一条边，将边的两个顶点进行合并（union）。
	最终，所有互相连通的顶点会属于同一个集合，通过查找（find）操作可以判断两个顶点是否连通。
	统计有多少个不同的根节点，就有多少个连通分量。
时间复杂度O(n+m・α(n))（α为阿克曼函数，接近O(1)），空间复杂度O(n)。
*/

func CountComponentsBFS(n int, edges [][2]int) int {
	adj := make([][]int, n) // 邻接表，(index:node)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n) //(index:node)
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			visited[i] = true
			queue := []int{i}
			for len(queue) > 0 {
				node := queue[0]
				queue = queue[1:]
				for _, neighbor := range adj[node] {
					if !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
		}
	}
	return count
}
func CountComponentsDFS(n int, edges [][2]int) int {
	adj := make([][]int, n) // 邻接表，(index:node)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n) // (index:node)
	count := 0
	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i)
		}
	}
	return count
}
func CountComponentsUnionFind(n int, edges [][2]int) int {
	parent := make([]int, n) // (index:node)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // 路径压缩
		}
		return parent[x]
	}
	for _, e := range edges {
		rootX := find(e[0])
		rootY := find(e[1])
		if rootX != rootY {
			parent[rootX] = rootY
		}
	}
	count := 0 // 统计根节点数量（根节点满足parent[i] == i）
	for i := 0; i < n; i++ {
		if find(i) == i {
			count++
		}
	}
	return count
}

/*
判断有向图是否存在环，可以使用深度优先搜索（DFS），
其核心思路是在遍历过程中检测是否遇到了回边（指向已访问但尚未回溯的节点）。
维护三个状态标记：未访问、正在访问（递归栈中）、已访问
当访问到一个节点时：
    标记为 "正在访问"
    递归访问其所有邻接节点
    若邻接节点处于 "正在访问" 状态，说明发现环
    遍历完成后标记为 "已访问"
*/

func HasCircle(n int, edges [][2]int) bool {
	adj := make([][]int, n) // 邻接表，(index:node)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	visited := make([]int8, n) // 0-未访问，1-正在访问，2-已访问 (index:node)
	var dfs func(node int) bool
	dfs = func(node int) bool {
		if visited[node] == 1 {
			return true // 发现环
		}
		if visited[node] == 2 {
			return false // 已访问过，无环
		}
		visited[node] = 1 // 标记为正在访问
		for _, neighbor := range adj[node] {
			if dfs(neighbor) {
				return true
			}
		}
		visited[node] = 2 // 标记为已访问
		return false
	}
	for i := 0; i < n; i++ { // 检查所有未访问节点
		if visited[i] == 0 && dfs(i) {
			return true
		}
	}
	return false
}

/*
拓扑排序（Topological-Order）是一个有向无环图（Directed-Acyclic-Graph）的所有顶点的线性序列。
入度：指向该顶点的边的数量。出度：该顶点指向的其他点边的数量。

拓扑排序最经典的算法是Kahn算法，也叫入度表算法：
1. 按照一定的顺序进行构造有向图，记录后个节点的入度；
2. 从图中选择一个入度为0的顶点，输出该顶点；
3. 从图中删除该顶点及所有与该顶点相连的边
4. 重复上述两步，直至所有顶点输出。
5. 或者当前图中不存在入度为0的顶点为止。此时可说明图中有环。
6. 因此，也可以通过拓扑排序来判断一个图是否有环。

时间复杂度O(n+m)，空间复杂度O(n)，n为顶点数，m为边数。
*/

func TopologicalOrder(n int, edges [][2]int) []int {
	adj := make([][]int, n) // 邻接表(index:node)
	deg := make([]int, n)   // 入度表(index:node)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		deg[e[1]]++
	}
	queue := make([]int, 0, n) //辅助队列，需要最小字典顺序时改用优先队列
	for i, d := range deg {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	result := make([]int, 0, n)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		for _, neighbor := range adj[node] {
			deg[neighbor]--
			if deg[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if len(result) == n {
		return result
	}
	return nil
} // 如仅需判断是否有环，可将result切片追加改为count计数变量自增，count==n代表无环。
