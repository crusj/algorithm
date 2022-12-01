package simulation

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

func Clone(head *RandomListNode) *RandomListNode {
	// write your code here
	copied := make(map[int]map[int]bool)
	nodes := make(map[int]*RandomListNode)

	var fn func(from, old, new *RandomListNode)
	fn = func(from, old, new *RandomListNode) {
		new.Label = old.Label

		if copied[old.Label] == nil {
			copied[old.Label] = make(map[int]bool)
		}
		copied[old.Label][from.Label] = true

		// 递归copy next
		if old.Next != nil && !copied[old.Next.Label][old.Label] {
			if nodes[old.Next.Label] != nil {
				new.Next = nodes[old.Next.Label]
			} else {
				nodes[old.Next.Label] = &RandomListNode{}
				new.Next = nodes[old.Next.Label]
				fn(old, old.Next, new.Next)
			}
		}

		// 递归copy random
		if old.Random != nil && !copied[old.Random.Label][old.Label] {
			if nodes[old.Random.Label] != nil {
				new.Random = nodes[old.Random.Label]
			} else {
				nodes[old.Random.Label] = &RandomListNode{}
				new.Random = nodes[old.Random.Label]
				fn(old, old.Random, new.Random)
			}
		}
	}

	newHead := &RandomListNode{}
	fn(&RandomListNode{}, head, newHead)

	return newHead
}
