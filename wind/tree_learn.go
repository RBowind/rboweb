package wind

type tree interface {
	insert() tree
	search() tree
}

//BinaryTree 二叉树
type BinaryTree struct {
	Left  *BinaryTree
	Value int
	Right *BinaryTree
}

func (bt *BinaryTree) insert() tree {
	if bt == nil {
		return nil
	}
	return bt
}

func (bt *BinaryTree) search() tree {
	return bt
}

//Btree B树
type Btree struct {
}

//AVLTree 平衡二叉树
type AVLTree struct{}

//RedBlackTree 红黑树
type RedBlackTree struct{}
