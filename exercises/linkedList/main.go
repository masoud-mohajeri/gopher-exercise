package linkedList

func LinkedListDemo() {
	list := List{}
	list.push(1)
	list.push(2)
	list.push(3)
	list.pop()
	list.push(4)
	list.push(5)

	list.show() // -> 4 -> 3 -> 2 -> 1
}
