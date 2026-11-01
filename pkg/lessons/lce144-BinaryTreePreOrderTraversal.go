package lessons

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	var result []int = []int{}
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	if root.Left != nil {
		result = append(result, PreorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, PreorderTraversal(root.Right)...)
	}
	return result
}

func PreorderTraversalStack(root *TreeNode) []int {
	var result []int = []int{}
	if root == nil {
		return result
	}
	stack := NewStack()
	stack.Push(root)
	for stack.Size() > 0 {
		node := stack.Pop().(*TreeNode)
		result = append(result, node.Val)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
	return result
}
