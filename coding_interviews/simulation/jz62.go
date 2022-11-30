package simulation

func LastRemaining_Solution(n int, m int) int {
	if m == 0 {
		return n
	}

	// write code here
	head := &Node{
		value: 0,
	}

	tmp := head
	for i := 1; i < n; i++ {
		t := &Node{
			value: i,
		}

		tmp.next = t
		tmp = tmp.next
	}
	tmp.next = head

	i := 1
	f := head
	p := head.next

	for p != f {
		if i == m-1 {
			f.next = p.next
			p = p.next

			i = 0
			continue
		}

		p = p.next
		f = f.next
		i++
	}
	return p.value
}

// 0 1 2 3
// 2

type Node struct {
	value int
	next  *Node
}
