package list

type TreeNode struct {
	left         *TreeNode   // left sub node
	right        *TreeNode   // right sub node
	subNodeCount int         //sub node count
	key          int         // sort value
	value        interface{} // node value
}

func (treeNode *TreeNode) Put(data interface{}, key int) {
	if treeNode.value == nil {
		treeNode.value = data
		treeNode.key = key
		return
	}

	if treeNode.key == key {
		treeNode.value = data
		return
	}

	if treeNode.key > key {
		treeNode.left.Put(data, key)
	}

	if treeNode.key < key {
		treeNode.right.Put(data, key)
	}

	treeNode.subNodeCount = treeNode.left.size() + treeNode.right.size()

}

func (treeNode *TreeNode) size() int {
	if treeNode.value == nil {
		return 0
	} else {
		return treeNode.subNodeCount
	}
}

func (treeNode *TreeNode) Get(key int) *TreeNode {
	if treeNode.value == nil {
		return nil
	}
	n := treeNode
	for n != nil {
		if n.key == key {
			break
		}
		if n.key < key {
			n = treeNode.right
		}
		if n.key > key {
			n = treeNode.left
		}
	}
	return n
}

func (treeNode *TreeNode) GetMaxValue() *TreeNode {
	return nil
}

func (treeNode *TreeNode) GetMinValue() *TreeNode {
	return nil
}
