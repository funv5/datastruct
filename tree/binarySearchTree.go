package tree

// BinarySearchTree ...
type BinarySearchTree struct {
	root *treeNode
}

type treeNode struct {
	key   interface{}
	left  *treeNode
	right *treeNode
}

// Insert ...
func (b *BinarySearchTree) Insert(key interface{}) {
	newNode := treeNode{}
	newNode.key = key

	if b.root == nil {
		b.root = &newNode
	} else {
		insertNode(b.root, &newNode)
	}
}

func insertNode(node, newNode *treeNode) {
	newkey, _ := newNode.key.(int)
	key, _ := node.key.(int)
	if newkey < key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// Search ...
func (b *BinarySearchTree) Search(key interface{}) bool {
	return false
}

type callbackFunc func(key interface{})

// InOrderTraverse ...
func (b *BinarySearchTree) InOrderTraverse(callback callbackFunc) {
	inOrderTraverse(b.root, callback)
}

func inOrderTraverse(node *treeNode, callback callbackFunc) {
	if node != nil {
		inOrderTraverse(node.left, callback)
		callback(node.key)
		inOrderTraverse(node.right, callback)
	}
}

// PreOrderTraverse ...
func (b *BinarySearchTree) PreOrderTraverse() {

}

// PostOrderTraverse ...
func (b *BinarySearchTree) PostOrderTraverse() {

}

// Min ...
func (b *BinarySearchTree) Min() interface{} {
	return nil
}

// Max ...
func (b *BinarySearchTree) Max() interface{} {
	return nil
}

// Remove ...
func (b *BinarySearchTree) Remove() interface{} {
	return nil
}
