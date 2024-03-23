package leetcode

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func printGraph(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("Node: %d, Neighbors: ", node.Val)
	for _, neighbor := range node.Neighbors {
		fmt.Printf("%d ", neighbor.Val)
	}
	fmt.Println()
	for _, neighbor := range node.Neighbors {
		printGraph(neighbor)
	}
}

func dfsClone(node *Node, copies map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if val, ok := copies[node.Val]; ok {
		return val
	}
	copy := &Node{Val: node.Val}
	copies[copy.Val] = copy

	clonedNeighbers := make([]*Node, len(node.Neighbors))
	for i, neighbor := range node.Neighbors {
		clonedNeighbers[i] = dfsClone(neighbor, copies)
	}

	copy.Neighbors = clonedNeighbers
	return copy
}

func bfsClone(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, node)
	clonedNode := &Node{Val: node.Val}

	copies := make(map[int]*Node)
	copies[node.Val] = clonedNode

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		for _, neighbor := range currentNode.Neighbors {
			if _, ok := copies[neighbor.Val]; !ok {
				copies[neighbor.Val] = &Node{neighbor.Val, nil}
				queue = append(queue, neighbor)
			}
			copies[currentNode.Val].Neighbors = append(copies[currentNode.Val].Neighbors, copies[neighbor.Val])
		}
	}
	return clonedNode
}
func CloneGraph() {

	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = append(node1.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node3)

	node4.Neighbors = append(node4.Neighbors, node3)
	fmt.Println("initial value:")
	sample := *node1
	printGraph(node1)

	copies := make(map[int]*Node, 100)
	newNode := dfsClone(&sample, copies)
	fmt.Println("dfs cloned value:")
	printGraph(newNode)

	bfs := bfsClone(&sample)
	fmt.Println("bfs cloned value:")
	printGraph(bfs)

}
