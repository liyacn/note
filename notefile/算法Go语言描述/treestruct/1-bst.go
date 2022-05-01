package treestruct

type BstNode struct {
	value int      // 值
	left  *BstNode // 左子树
	right *BstNode // 右子树
}
type BinarySearchTree struct {
	root *BstNode // 根结点
}

func (tree *BinarySearchTree) Search(v int) *BstNode {
	for p := tree.root; p != nil; {
		if v == p.value {
			return p
		}
		if v < p.value {
			p = p.left
		} else {
			p = p.right
		}
	}
	return nil
}

func (tree *BinarySearchTree) Min() *BstNode {
	if tree.root == nil {
		return nil
	}
	res := tree.root
	for res.left != nil {
		res = res.left
	}
	return res
}

func (tree *BinarySearchTree) Max() *BstNode {
	if tree.root == nil {
		return nil
	}
	res := tree.root
	for res.right != nil {
		res = res.right
	}
	return res
}

func (tree *BinarySearchTree) Insert(v int) bool {
	if tree.root == nil {
		tree.root = &BstNode{value: v}
		return true
	}
	p := tree.root
	for {
		if v == p.value {
			return false
		}
		if v < p.value {
			if p.left == nil {
				p.left = &BstNode{value: v}
				break
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = &BstNode{value: v}
				break
			}
			p = p.right
		}
	}
	return true
}

/*
删除值所在结点的不同情况：
未找到不作任何处理，无子树直接删除。
有一个子树，该子树直接替换被删的结点。
有两个子树，右子树的最小值结点（或左子树最大结点）替换被删的结点。
    p?                 p?
    |                 ||
   3(o)              4(n)
  /  \              // \\
 2    8            2    8
     / \               / \
   6(t) 9   ==>      6(t) 9
  / \               // \
4(n) 7             5    7
\
 5
3(o)右子树的最小结点是4(n)，将4(n)移动到3(o)的位置，n.left拼上原o.left。
当t!=o即n!=o.right时，还需将t.left拼上原n.right(含nil)，再将n.right拼上原o.right。
最后当p!=nil时将p和n连接上，p==nil时n直接设为根结点。
*/

func (tree *BinarySearchTree) Delete(v int) bool {
	o := tree.root //待删除old结点
	var p *BstNode //o.parent
	for o != nil {
		if v == o.value { //找到待删除结点
			break
		}
		p = o
		if v < o.value {
			o = o.left
		} else {
			o = o.right
		}
	}
	if o == nil { //未找到
		return false
	}
	var n *BstNode //替换old的new结点
	if o.left == nil {
		n = o.right //左子树为空，右子树替换(含nil)
	} else if o.right == nil {
		n = o.left //右子树为空，左子树替换
	} else { //左右子树均不为空，右子树的最小结点(或左子树的最大结点)替换待删除结点
		t := o      //n.parent
		n = o.right //右子树的最小结点
		for n.left != nil {
			t = n
			n = n.left
		}
		n.left = o.left //new结点拼上old结点的左子树
		if t != o {     //即n!=o.right
			t.left = n.right
			n.right = o.right
		}
	}
	if p == nil {
		tree.root = n
	} else {
		if v < p.value {
			p.left = n //是其左子树
		} else {
			p.right = n //是其右子树
		}
	}
	return true
}

// 中序遍历返回有序数组
func (tree *BinarySearchTree) InOrderTraversal() []int {
	var res []int
	tree.root.inOrderTraversal(&res)
	return res
}
func (node *BstNode) inOrderTraversal(res *[]int) {
	if node != nil {
		node.left.inOrderTraversal(res)
		*res = append(*res, node.value)
		node.right.inOrderTraversal(res)
	}
}
