package node

import (
	"fmt"
	"strings"
)

// Node struct is used to build a Binary Tree/Node
type Node struct {
	Value       int
	left, right *Node
}

// Add adds a Node element to the Binary tree
func Add(n *Node, value int) *Node {
	if n == nil {
		aux := Node{Value: value}
		return &aux
	}

	if value == n.Value {
		return n
	}

	if value < n.Value {
		n.left = Add(n.left, value)
	} else {
		n.right = Add(n.right, value)
	}

	return n
}

// PrintAll prints the full Tree in all the orders
func PrintAll(n *Node) {
	// fmt.Println("Pre-order tree traversal:", TraversePreOrder(n))
	// fmt.Println("In-order tree traversal:", TraverseInOrder(n))
	// fmt.Println("Post-order tree traversal:", TraversePostOrder(n))

	// fmt.Println()
	// using functions as parameters:

	printer("Pre-order tree traversal:", n, TraversePreOrder)
	printer("In-order tree traversal:", n, TraverseInOrder)
	printer("Post-order tree traversal:", n, TraversePostOrder)
}

func printer(msg string, n *Node, f func(n *Node) string) {
	fmt.Println(msg, f(n))
}

// TraversePreOrder prints the Tree in order
func TraversePreOrder(n *Node) string {
	if n == nil {
		return ""
	}

	var ss []string
	ss = append(ss, fmt.Sprintf("%d ", n.Value))
	ss = append(ss, TraversePreOrder(n.left))
	ss = append(ss, TraversePreOrder(n.right))

	return strings.TrimRight(strings.Join(ss, ""), " ")
}

// TraverseInOrder prints the Tree in order
func TraverseInOrder(n *Node) string {
	if n == nil {
		return ""
	}

	var ss []string
	ss = append(ss, TraverseInOrder(n.left))
	ss = append(ss, fmt.Sprintf("%d ", n.Value))
	ss = append(ss, TraverseInOrder(n.right))

	return strings.TrimRight(strings.Join(ss, ""), " ")
}

// TraversePostOrder prints the Tree in order
func TraversePostOrder(n *Node) string {
	if n == nil {
		return ""
	}

	var ss []string
	ss = append(ss, TraversePostOrder(n.left))
	ss = append(ss, TraversePostOrder(n.right))
	ss = append(ss, fmt.Sprintf("%d ", n.Value))

	return strings.TrimRight(strings.Join(ss, ""), " ")
}
