package treestruct

type BstNode struct {
	Value int      // 值
	Count int      // 值出现次数
	Left  *BstNode // 左子树
	Right *BstNode // 右子树
}
type BinarySearchTree struct {
	Root *BstNode // 根结点
}

func (tree *BinarySearchTree) Search(v int) *BstNode {
	for p := tree.Root; p != nil; {
		if p.Value == v {
			return p
		}
		if v < p.Value {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return nil
}

func (tree *BinarySearchTree) Min() *BstNode {
	if tree.Root == nil {
		return nil
	}
	res := tree.Root
	for res.Left != nil {
		res = res.Left
	}
	return res
}

func (tree *BinarySearchTree) Max() *BstNode {
	if tree.Root == nil {
		return nil
	}
	res := tree.Root
	for res.Right != nil {
		res = res.Right
	}
	return res
}

func (tree *BinarySearchTree) Insert(v int) {
	if tree.Root == nil {
		tree.Root = &BstNode{Value: v, Count: 1}
		return
	}
	p := tree.Root
	for {
		if v == p.Value {
			p.Count++
			return
		}
		if v < p.Value {
			if p.Left == nil {
				p.Left = &BstNode{Value: v, Count: 1}
				return
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &BstNode{Value: v, Count: 1}
				return
			}
			p = p.Right
		}
	}
}

/*
Delete 值所在结点的不同情况：
未找到不作任何处理，无子树直接删除。
有两个子树，右子树的最小值（或左子树最大值）所在结点替换被删的结点。
有一个子树，该子树直接替换被删的结点（优于按两个子树的逻辑处理）。

	    p?                 p?
	    |                 ||
	  3(c)               4(s)
	  /  \              // \\
	2     8            2    8
		 / \               / \
	   6(t) 9   ==>      6(t) 9
	   / \               // \
	 4(s) 7             5    7
	 \
	  5

3(c)右子树的最小元素是4(s)，将s移动到3(c)的位置，s.Left拼上原c.Left。
当t!=c即s!=c.Right时，还需将t.Left拼上原s.Right(含nil)，再将s.Right拼上原c.Right。
最后当p！=nil时将p和s连接上，p==nil时s直接设为根结点。
*/
func (tree *BinarySearchTree) Delete(v int) {
	var p *BstNode
	for c := tree.Root; c != nil; {
		if v == c.Value {
			if c.Count > 1 {
				c.Count--
			} else {
				var s *BstNode     // 删除c后的新子树的根结点
				if c.Left == nil { // 左子树为空，右子树替换
					s = c.Right // 包含了两个子树均为空的case
				} else if c.Right == nil { // 右子树为空，左子树替换
					s = c.Left
				} else { // 左右子树均不为空，右子树的最小元素替换待删除结点（Left,Right反过来则是左子树的最大元素替换待删除结点）
					t := c      // 右子树最小结点的双亲结点
					s = c.Right // 右子树最小结点
					for s.Left != nil {
						t = s
						s = s.Left
					}
					s.Left = c.Left // 新的根结点拼接待删除结点的左子树
					if t != c {     // 也即 s != c.Right
						t.Left = s.Right  // parent结点拼接取出元素的子树
						s.Right = c.Right // 新根结点拼接待删除结点的右子树
					}
				}
				if p == nil {
					tree.Root = s
				} else {
					if v < p.Value { // 是其左结点
						p.Left = s
					} else { // 是其右结点
						p.Right = s
					}
				}
			}
			return
		}
		p = c
		if v < c.Value {
			c = c.Left
		} else {
			c = c.Right
		}
	}
}

// InOrderTraversal 中序遍历返回有序数组
func (tree *BinarySearchTree) InOrderTraversal() []int {
	var res []int
	tree.Root.inOrderTraversal(&res)
	return res
}
func (node *BstNode) inOrderTraversal(res *[]int) {
	if node != nil {
		node.Left.inOrderTraversal(res)
		for i := 0; i < node.Count; i++ { // 重复值
			*res = append(*res, node.Value)
		}
		node.Right.inOrderTraversal(res)
	}
}
