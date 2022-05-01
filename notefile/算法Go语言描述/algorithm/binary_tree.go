package algorithm

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
      4
     /\
    2  6
   /|  |\
  1 3  5 8
 /       |\
0        7 9
*/
var tree = &TreeNode{
	Val: 4,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	},
	Right: &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	},
}

//深度优先遍历（Depth-First-Search）：先序、中序、后序
//广度优先遍历（Breadth-First-Search）：层序

// 先序遍历递归
func PreOrderTraversal(node *TreeNode) {
	if node != nil {
		fmt.Print(node.Val, " ")
		PreOrderTraversal(node.Left)
		PreOrderTraversal(node.Right)
	}
} // 4 2 1 0 3 6 5 8 7 9

// 中序遍历递归
func InOrderTraversal(node *TreeNode) {
	if node != nil {
		InOrderTraversal(node.Left)
		fmt.Print(node.Val, " ")
		InOrderTraversal(node.Right)
	}
} //0 1 2 3 4 5 6 7 8 9

// 后续遍历递归
func PostOrderTraversal(node *TreeNode) {
	if node != nil {
		PostOrderTraversal(node.Left)
		PostOrderTraversal(node.Right)
		fmt.Print(node.Val, " ")
	}
} // 0 1 3 2 5 7 9 8 6 4

// 先序遍历模拟栈代替递归
func (node *TreeNode) PreOrderTraversal() {
	if node == nil {
		return
	}
	var stack []*TreeNode
	p := node
	for p != nil || len(stack) > 0 {
		for p != nil { //到最左下叶
			fmt.Print(p.Val, " ")    //访问结点
			stack = append(stack, p) //结点入栈
			p = p.Left
		}
		p = stack[len(stack)-1]      //结点不再访问
		stack = stack[:len(stack)-1] //直接弹出
		p = p.Right
	}
} // 4 2 1 0 3 6 5 8 7 9

// 中序遍历模拟栈代替递归
func (node *TreeNode) InOrderTraversal() {
	if node == nil {
		return
	}
	var stack []*TreeNode
	p := node
	for p != nil || len(stack) > 0 {
		for p != nil { //到最左下叶
			stack = append(stack, p) //结点入栈
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(p.Val, " ") //弹出时访问
		p = p.Right
	}
} // 0 1 2 3 4 5 6 7 8 9

// 后序遍历模拟栈代替递归
func (node *TreeNode) PostOrderTraversal() {
	if node == nil {
		return
	}
	var stack []*TreeNode
	p := node
	var tmp *TreeNode //辅助结点
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]               //获取栈顶元素，先不出栈
		if p.Right != nil && p.Right != tmp { //右子树不为空且未访问
			p = p.Right
		} else { //右子树已经访问或为空，接下来出栈访问结点，第二次栈顶
			fmt.Print(p.Val, " ")
			stack = stack[:len(stack)-1] //出栈
			tmp = p                      //指向访问过的右子树结点
			p = nil                      //使得p为空继续访问栈顶
		}
	}
} // 0 1 3 2 5 7 9 8 6 4

// 层序遍历：每一层从左到右访问每一个节点，需要使用辅助的先进先出的队列。
func (node *TreeNode) LevelOrderTraversal() {
	if node == nil {
		return
	}
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:] //队头元素不断出队
		fmt.Print(p.Val, " ")
		if p.Left != nil {
			queue = append(queue, p.Left) //左子树非空，入队列
		}
		if p.Right != nil {
			queue = append(queue, p.Right) //右子树非空，入队列
		}
	}
} // 4 2 6 1 3 5 8 0 7 9

// 层序遍历分层打印
func (node *TreeNode) LevelTraversal() {
	if node == nil {
		return
	}
	level := 1
	queue := []*TreeNode{node}
	for size := 1; size > 0; size = len(queue) {
		fmt.Printf("\n<%v>:", level)
		for i := 0; i < size; i++ { //遍历本层
			p := queue[i]
			fmt.Print(" ", p.Val)
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
		queue = queue[size:] //整层出队
		level++
	}
} /*
   <1>: 4
   <2>: 2 6
   <3>: 1 3 5 8
   <4>: 0 7 9
*/

// 根据先序和中序遍历构建二叉树
func BuildTree(preorder []int, inorder []int) *TreeNode {
	size := len(preorder)
	if size == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]} //先序遍历最先访问根结点
	i := 0
	for i < size {
		if inorder[i] == root.Val {
			break
		}
		i++
	}
	root.Left = BuildTree(preorder[1:i+1], inorder[:i])
	root.Right = BuildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func BuildTreeByPreorder(preorder []int, inorder []int) *TreeNode {
	size := len(preorder)
	if size == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	inorderIndex := 0
	for i := 1; i < len(preorder); i++ {
		p := stack[len(stack)-1]
		node := &TreeNode{Val: preorder[i]}
		if p.Val != inorder[inorderIndex] {
			p.Left = node
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				p = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			p.Right = node
		}
		stack = append(stack, node)
	}
	return root
}

// 根据中序和后序遍历构建二叉树
func CreateTree(inorder []int, postorder []int) *TreeNode {
	size := len(inorder)
	if size == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[size-1]} //后序遍历最后访问根结点
	i := 0
	for i < size {
		if inorder[i] == root.Val {
			break
		}
		i++
	}
	root.Left = CreateTree(inorder[:i], postorder[:i])
	root.Right = CreateTree(inorder[i+1:], postorder[i:size-1])
	return root
}

func CreateTreeByPostorder(inorder []int, postorder []int) *TreeNode {
	size := len(inorder)
	if size == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[size-1]}
	stack := []*TreeNode{root}
	inorderIndex := size - 1
	for i := size - 2; i >= 0; i-- {
		p := stack[len(stack)-1]
		node := &TreeNode{Val: postorder[i]}
		if p.Val != inorder[inorderIndex] {
			p.Right = node
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				p = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex--
			}
			p.Left = node
		}
		stack = append(stack, node)
	}
	return root
}
