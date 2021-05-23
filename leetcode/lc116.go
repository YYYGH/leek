package leetcode

/**
 * Definition for a Node.
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTreeNode(root.Left, root.Right)
	return root
}

func connectTreeNode(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTreeNode(node1.Left, node1.Right)
	connectTreeNode(node2.Left, node2.Right)
	connectTreeNode(node1.Right, node2.Left)
}

func ConnectV2(root *Node) *Node {
	for root == nil {
		return nil
	}

	for left := root; left != nil; left = left.Left {
		for node := left; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}
