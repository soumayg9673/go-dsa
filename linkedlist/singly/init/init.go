package main

type Node struct {
	data int
	next *Node
}

func main() {
	head := &Node{data: 10}

	// Implementation: 1
	Imp1(head)

	// Implementation: 2
	Imp2(head)
}

// Implementation: 1
func Imp1(head *Node) {
	node1 := &Node{data: 20}
	node2 := &Node{data: 30}

	head.next = node1
	node1.next = node2
}

// Implementation: 2
func Imp2(head *Node) {
	head.next = &Node{data: 20}
	head.next.next = &Node{data: 30}
}
