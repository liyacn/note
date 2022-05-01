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
	Value  int      // 值
	Count  int      // 值出现次数
	BF     int8     // 平衡因子（Balance-Factor：右子树高度-左子树高度）
	Parent *AvlNode // 双亲结点
	Left   *AvlNode // 左子树
	Right  *AvlNode // 右子树
}
type AvlTree struct {
	Root *AvlNode
}

/*
Insert 按照BST的规则将新结点插入到AVL树中。
插入结点是parent结点的左叶子则p.bf-1，插入结点是parent结点的右叶子则p.bf+1。
如果parent的bf为0，说明插入前是±1，未改变平衡性，插入成功。
如果parent的bf为±1，说明插入前平衡因子为0，以parent为根的树高度增加，需要继续向上更新。
如果parent的bf为±2，则需要对p结点做旋转处理。
*/
func (tree *AvlTree) Insert(v int) {
	if tree.Root == nil {
		tree.Root = &AvlNode{Value: v, Count: 1}
		return
	}
	var c *AvlNode
	p := tree.Root
	for {
		if v == p.Value {
			p.Count++
			return
		}
		if v < p.Value {
			if p.Left == nil {
				c = &AvlNode{Value: v, Count: 1, Parent: p}
				p.Left = c
				break
			}
			p = p.Left
		} else {
			if p.Right == nil {
				c = &AvlNode{Value: v, Count: 1, Parent: p}
				p.Right = c
				break
			}
			p = p.Right
		}
	}
	for p != nil {
		if c.Value < p.Value {
			p.BF--
		} else {
			p.BF++
		}
		if p.BF == 0 {
			return
		}
		if p.BF == -2 {
			if c.BF < 0 { // 左左失衡，右单旋
				tree.rotateR(p)
			} else { // 左右双旋
				tree.rotateLR(p)
			}
			return
		} else if p.BF == 2 {
			if c.BF > 0 { // 右右失衡，左单旋
				tree.rotateL(p)
			} else { // 右左双旋
				tree.rotateRL(p)
			}
			return
		}
		c, p = p, p.Parent // 1和-1继续向上翻找
	}
}

/*
Delete 删除规则与BST类似。
删除结点是parent结点的左叶子则p.bf+1，删除结点是parent结点的右叶子则p.bf-1。
如果parent的bf为±1，说明删除前为0，未改变子树高度，无需调整。
如果parent的bf为0，说明删除前是±1，以parent为根的树高度减少，需要继续向上更新。
如果parent的bf为±2，则需要对p结点做旋转处理。

			!
			3
		   / \
		  2   6
		 /   / \
	    1   4   7
		     \   \
		      5   8

删除左子树为空的结点c(4)，s = c.right(5)，直接改变p(6).bf，从p(6)向上调整。
删除右子树为空的结点c(2)，s = c.left(1)，直接改变p(3).bf，从p(3)向上调整。
删除两个子树均不为空的结点c(6)，右子树最小结点为s(7)，此时s==c.right，s.bf=c.bf-1，从s(7)向上调整。
删除两个子树均不为空的结点c(3)，右子树最小结点为s(4)，此时t(s.parent(6))!=c，t.bf自增1，从t(6)向上调整。
*/
func (tree *AvlTree) Delete(v int) {
	for c := tree.Root; c != nil; {
		if v == c.Value {
			if c.Count > 1 {
				c.Count--
			} else {
				p := c.Parent
				n := p         // 向上调整的起点
				var s *AvlNode // 删除c后的新子树的根结点
				mbf := false   // p.bf直接改变
				if c.Left == nil {
					s = c.Right
					mbf = true
				} else if c.Right == nil {
					s = c.Left
					mbf = true
				} else {
					s = c.Right // 右子树最小结点
					for s.Left != nil {
						s = s.Left
					}
					s.Left = c.Left // 新的根结点拼接待删除结点的左子树
					c.Left.Parent = s
					if t := s.Parent; t == c { // 也即 s == c.Right
						s.BF = c.BF - 1
						n = s
					} else {
						s.BF = c.BF
						t.Left = s.Right // parent结点拼接取出元素的子树
						if s.Right != nil {
							s.Right.Parent = t
						}
						s.Right = c.Right // 新根结点拼接待删除结点的右子树
						c.Right.Parent = s
						t.BF++
						n = t
					}
				}
				if s != nil {
					s.Parent = p
				}
				if p == nil {
					tree.Root = s
				} else {
					if v < p.Value { // 是其左结点
						p.Left = s
						if mbf {
							p.BF++
						}
					} else { // 是其右结点
						p.Right = s
						if mbf {
							p.BF--
						}
					}
				}
				for n != nil {
					if n.BF == -1 || n.BF == 1 {
						return
					}
					if n.BF == -2 {
						if n.Left.BF <= 0 { // 左左失衡，右单旋
							tree.rotateR(n)
						} else { // 左右双旋
							tree.rotateLR(n)
						}
						n = n.Parent
					} else if n.BF == 2 {
						if n.Right.BF >= 0 { // 右右失衡，左单旋
							tree.rotateL(n)
						} else { // 右左双旋
							tree.rotateRL(n)
						}
						n = n.Parent
					} else {
						r := n.Parent
						if r != nil {
							if n.Value < r.Value {
								r.BF++
							} else {
								r.BF--
							}
						}
						n = r
					}
				}
			}
			return
		}
		if v < c.Value {
			c = c.Left
		} else {
			c = c.Right
		}
	}
}

func (tree *AvlTree) rotateR(p *AvlNode) {
	r := p.Parent
	c := p.Left
	p.Left = c.Right
	if c.Right != nil {
		c.Right.Parent = p
	}
	c.Right = p
	p.Parent = c
	c.Parent = r
	if r == nil {
		tree.Root = c
	} else {
		if c.Value < r.Value {
			r.Left = c
		} else {
			r.Right = c
		}
	}
	if p.BF == -2 && c.BF == 0 {
		p.BF, c.BF = -1, 1
	} else if p.BF == -2 && c.BF == -1 {
		p.BF, c.BF = 0, 0
	} else if p.BF == -2 && c.BF == -2 {
		p.BF, c.BF = 1, 0
	} else if p.BF == -1 && c.BF == 1 {
		p.BF, c.BF = 0, 2
	} else if p.BF == -1 && c.BF == 0 {
		p.BF, c.BF = 0, 1
	} else if p.BF == -1 && c.BF == -1 {
		p.BF, c.BF = 1, 1
	}
}
func (tree *AvlTree) rotateL(p *AvlNode) {
	r := p.Parent
	c := p.Right
	p.Right = c.Left
	if c.Left != nil {
		c.Left.Parent = p
	}
	c.Left = p
	p.Parent = c
	c.Parent = r
	if r == nil {
		tree.Root = c
	} else {
		if c.Value < r.Value {
			r.Left = c
		} else {
			r.Right = c
		}
	}
	if p.BF == 2 && c.BF == 0 {
		p.BF, c.BF = 1, -1
	} else if p.BF == 2 && c.BF == 1 {
		p.BF, c.BF = 0, 0
	} else if p.BF == 2 && c.BF == 2 {
		p.BF, c.BF = -1, 0
	} else if p.BF == 1 && c.BF == -1 {
		p.BF, c.BF = 0, -2
	} else if p.BF == 1 && c.BF == 0 {
		p.BF, c.BF = 0, -1
	} else if p.BF == 1 && c.BF == 1 {
		p.BF, c.BF = -1, -1
	}
}

func (tree *AvlTree) rotateLR(p *AvlNode) {
	tree.rotateL(p.Left)
	tree.rotateR(p)
}
func (tree *AvlTree) rotateRL(p *AvlNode) {
	tree.rotateR(p.Right)
	tree.rotateL(p)
}

/*
Height 获取任意结点高度。nil结点高度定为0，叶子结点高度定为1。
根据平衡因子可以锁定最高子树的方向，大于0向右，小于0向左，等于0两边任意。
for循环根据最高子树方向搜索到底即可获取树的高度，无需使用递归。
*/
func (node *AvlNode) Height() int {
	h := 0
	for p := node; p != nil; {
		h++
		if p.BF > 0 {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	return h
}
