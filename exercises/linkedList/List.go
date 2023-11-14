package linkedList

import "fmt"

type List struct {
	head *Node
}

func (l *List) push(value int8) {
	node := Node{value: value, next: nil}

	node.next = l.head

	l.head = &node

}

func (l *List) pop() {
	if l.head.next != nil {
		l.head = l.head.next
	} else {
		l.head = nil
	}

}

func (l *List) show() {
	p := l.head

	for p != nil {
		value := p.value
		if p != l.head {
			fmt.Printf("-> %v ", value)
		} else {
			fmt.Printf("%v ", value)
		}
		p = p.next
	}
}
