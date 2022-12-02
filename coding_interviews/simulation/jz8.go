package simulation

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode.Right != nil {
		tmp := pNode.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}

		return tmp
	}

	if pNode.Next != nil {
		for pNode.Next != nil {
			if pNode.Next.Left == pNode {
				return pNode.Next
			}

			pNode = pNode.Next
		}
	}

	return nil
}
