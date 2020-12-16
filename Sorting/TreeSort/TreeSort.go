package TreeSort

type node struct {
	val int
	left *node
	right *node
}

type btree struct {
	root *node
}

func newNode(val int) *node{
	return &node{val,nil,nil}
}

func insert(root *node,val int) *node{
	if root == nil{
		return newNode(val)
	}
	if val < root.val{
		root.left = insert(root.left,val)
	}else{
		root.right = insert(root.right,val)
	}
	return root
}

func inorderCopy(n *node,array []int,index *int){
	if n!=nil{
		inorderCopy(n.left, array, index)
		array[*index] = n.val
		*index++
		inorderCopy(n.right, array, index)
	}
}

func TreeSort(arr []int,tree *btree) []int{
	for _, element := range arr {
		tree.root = insert(tree.root, element)
	}
	index := 0
	inorderCopy(tree.root, arr, &index)
	return arr
}