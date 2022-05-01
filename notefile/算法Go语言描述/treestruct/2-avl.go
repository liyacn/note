package treestruct

/*
AVL树是是高度平衡而非完美平衡的，如下是一棵AVL树：
       7
      / \
     5   8
	/ \   \
   4   6   9
 /
3
AVL树满足BST的性质，Search，Min，Max，InOrderTraversal等方法逻辑完全相同，Insert和Delete操作后需要调整平衡性。
*/

type AvlNode struct {
	value  int      // 值
	bf     int8     // 平衡因子（Balance-Factor：右子树高度-左子树高度）
	parent *AvlNode // 上层结点
	left   *AvlNode // 左子树
	right  *AvlNode // 右子树
}
type AvlTree struct {
	root *AvlNode
}

/*
按照BST的规则将新结点插入到AVL树中。
插入结点是parent结点的左叶子则p.bf-1，插入结点是parent结点的右叶子则p.bf+1。
如果parent的bf为0，说明插入前是±1，未改变平衡性，插入成功。
如果parent的bf为±1，说明插入前平衡因子为0，以parent为根的树高度增加，需要继续向上寻找。
如果parent的bf为±2，则找到最小不平衡树p，需要对p结点做旋转处理。
*/

func (tree *AvlTree) Insert(v int) bool {
	if tree.root == nil {
		tree.root = &AvlNode{value: v}
		return true
	}
	p := tree.root
	for {
		if v == p.value {
			return false
		}
		if v < p.value {
			if p.left == nil {
				p.left = &AvlNode{value: v, parent: p}
				break
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = &AvlNode{value: v, parent: p}
				break
			}
			p = p.right
		}
	}
	var cbf int8
	for p != nil {
		if v < p.value {
			p.bf--
		} else {
			p.bf++
		}
		if p.bf == 1 || p.bf == -1 {
			cbf = p.bf
			p = p.parent
			continue
		}
		if p.bf == -2 {
			if cbf < 0 { //LL型
				tree.rotateR(p)
			} else { //LR型
				tree.rotateLR(p)
			}
		} else if p.bf == 2 {
			if cbf > 0 { //RR型
				tree.rotateL(p)
			} else { //RL型
				tree.rotateRL(p)
			}
		}
		break
	}
	return true
}

/*
删除规则与BST类似。
删除结点是parent结点的左叶子则p.bf+1，删除结点是parent结点的右叶子则p.bf-1。
如果parent的bf为±1，说明删除前为0，未改变子树高度，无需调整。
如果parent的bf为0，说明删除前是±1，以parent为根的树高度减少，需要继续向上寻找。
如果parent的bf为±2，则需要对p结点做旋转处理，再继续向上寻找。
    !
    3
   / \
  2   6
 /   / \
1   4   7
     \   \
      5   8
1: n=nil(o.child), p(2).bf++
8: n=nil(o.child), p(7).bf--
2: n(1)=o(2).left, p(3).bf++
7: n(8)=o(7).right, p(6).bf--
6: n(7)=o(6).right, n(7).bf=o(6).bf-1
3: n(4)!=o(3).right, n(4).bf=o(3).bf, n.parent(6).bf++
*/

func (tree *AvlTree) Delete(v int) bool {
	o := tree.root //待删除old结点
	for o != nil {
		if v == o.value { //找到待删除结点
			break
		}
		if v < o.value {
			o = o.left
		} else {
			o = o.right
		}
	}
	if o == nil { //未找到
		return false
	}
	var n *AvlNode //替换old的new结点
	p := o.parent  //向上调整的起点
	if o.left == nil {
		n = o.right //左子树为空，右子树替换(含nil)
	} else if o.right == nil {
		n = o.left //右子树为空，左子树替换
	} else { //或使用对称逻辑
		n = o.right //右子树的最小结点
		for n.left != nil {
			n = n.left
		}
		n.left = o.left //new结点拼上old结点的左子树
		o.left.parent = n
		if n.parent == o { //即n==o.right
			n.bf = o.bf - 1
			p = n
		} else {
			n.bf = o.bf
			n.parent.bf++
			p = n.parent
			n.parent.left = n.right
			if n.right != nil {
				n.right.parent = n.parent
			}
			n.right = o.right
			o.right.parent = n
		}
	}
	if n != nil {
		n.parent = o.parent
	}
	if o.parent == nil {
		tree.root = n
	} else {
		if v < o.parent.value {
			o.parent.left = n //是其左子树
			if p == o.parent {
				p.bf++
			}
		} else {
			o.parent.right = n //是其右子树
			if p == o.parent {
				p.bf--
			}
		}
	}
	for p != nil {
		if p.bf == -1 || p.bf == 1 {
			break
		}
		if p.bf == -2 {
			if p.left.bf <= 0 { //LL型
				tree.rotateR(p)
			} else { //LR型
				tree.rotateLR(p)
			}
			p = p.parent
		} else if p.bf == 2 {
			if p.right.bf >= 0 { //RR型
				tree.rotateL(p)
			} else { //RL型
				tree.rotateRL(p)
			}
			p = p.parent
		} else {
			r := p.parent
			if r != nil {
				if p.value < r.value {
					r.bf++
				} else {
					r.bf--
				}
			}
			p = r
		}
	}
	return true
}

func (tree *AvlTree) rotateR(p *AvlNode) {
	r := p.parent
	c := p.left
	p.left = c.right
	if c.right != nil {
		c.right.parent = p
	}
	c.right = p
	p.parent = c
	c.parent = r
	if r == nil {
		tree.root = c
	} else {
		if c.value < r.value {
			r.left = c
		} else {
			r.right = c
		}
	}
	if p.bf == -2 && c.bf == 0 {
		p.bf, c.bf = -1, 1
	} else if p.bf == -2 && c.bf == -1 {
		p.bf, c.bf = 0, 0
	} else if p.bf == -2 && c.bf == -2 {
		p.bf, c.bf = 1, 0
	} else if p.bf == -1 && c.bf == 1 {
		p.bf, c.bf = 0, 2
	} else if p.bf == -1 && c.bf == 0 {
		p.bf, c.bf = 0, 1
	} else if p.bf == -1 && c.bf == -1 {
		p.bf, c.bf = 1, 1
	}
}
func (tree *AvlTree) rotateL(p *AvlNode) {
	r := p.parent
	c := p.right
	p.right = c.left
	if c.left != nil {
		c.left.parent = p
	}
	c.left = p
	p.parent = c
	c.parent = r
	if r == nil {
		tree.root = c
	} else {
		if c.value < r.value {
			r.left = c
		} else {
			r.right = c
		}
	}
	if p.bf == 2 && c.bf == 0 {
		p.bf, c.bf = 1, -1
	} else if p.bf == 2 && c.bf == 1 {
		p.bf, c.bf = 0, 0
	} else if p.bf == 2 && c.bf == 2 {
		p.bf, c.bf = -1, 0
	} else if p.bf == 1 && c.bf == -1 {
		p.bf, c.bf = 0, -2
	} else if p.bf == 1 && c.bf == 0 {
		p.bf, c.bf = 0, -1
	} else if p.bf == 1 && c.bf == 1 {
		p.bf, c.bf = -1, -1
	}
}
func (tree *AvlTree) rotateLR(p *AvlNode) {
	tree.rotateL(p.left)
	tree.rotateR(p)
}
func (tree *AvlTree) rotateRL(p *AvlNode) {
	tree.rotateR(p.right)
	tree.rotateL(p)
}
