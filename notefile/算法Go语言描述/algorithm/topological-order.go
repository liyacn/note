package algorithm

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

时间复杂度O(m+n)，空间复杂度O(n)，m为边数，n为顶点数。
*/

func TopologicalOrder(n int, edges [][2]int) []int { // 限定顶点为[0,n-1]
	deg := make([]int, n)         //入度表，顶点非紧凑时改用map
	adjacency := make([][]int, n) //邻接表，顶点非紧凑时改用map
	for _, e := range edges {     //e[0]和e[1]分别代表from和to
		adjacency[e[0]] = append(adjacency[e[0]], e[1])
		deg[e[1]]++
	}
	queue := make([]int, 0, n) //辅助队列，需要最小字典顺序时改用优先队列
	for i, d := range deg {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	res := make([]int, 0, n)
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		res = append(res, pos)
		for _, v := range adjacency[pos] {
			deg[v]--
			if deg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(res) == n {
		return res
	}
	return nil
}
