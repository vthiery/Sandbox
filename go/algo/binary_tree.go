package main

import (
	"errors"
	"fmt"
	"log"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Node definition
type Node struct {
	Value   string
	Data    string
	Left    *Node
	Right   *Node
	balance int
}

func (node *Node) insert(value, data string) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		node.Data = data
		return false
	}
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value, Data: data}
			if node.Right == nil {
				node.balance = -1
			} else {
				node.balance = 0
			}
		} else {
			if node.Left.insert(value, data) {
				if node.Left.balance < -1 || node.Left.balance > 1 {
					node.rebalance(node.Left)
				} else {
					node.balance--
				}
			}
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value, Data: data}
			if node.Left == nil {
				node.balance = 1
			} else {
				node.balance = 0
			}
		} else {
			if node.Right.insert(value, data) {
				if node.Right.balance < -1 || node.Right.balance > 1 {
					node.rebalance(node.Right)
				} else {
					node.balance++
				}
			}
		}
	}
	if node.balance != 0 {
		return true
	}
	return false
}

func (node *Node) rotateLeft(otherNode *Node) {
	r := otherNode.Right
	otherNode.Right = r.Left
	r.Left = otherNode
	if otherNode == node.Left {
		node.Left = r
	} else {
		node.Right = r
	}
	otherNode.balance = 0
	r.balance = 0
}

func (node *Node) rotateRight(otherNode *Node) {
	l := otherNode.Left
	otherNode.Left = l.Right
	l.Right = otherNode
	if otherNode == node.Left {
		node.Left = l
	} else {
		node.Right = l
	}
	otherNode.balance = 0
	l.balance = 0
}

func (node *Node) rotateRightLeft(otherNode *Node) {
	otherNode.Right.Left.balance = 1
	otherNode.rotateRight(otherNode.Right)
	otherNode.Right.balance = 1
	node.rotateLeft(otherNode)
}

func (node *Node) rotateLeftRight(otherNode *Node) {
	otherNode.Left.Right.balance = -1
	otherNode.rotateLeft(otherNode.Left)
	otherNode.Left.balance = -1
	node.rotateRight(otherNode)
}

func (node *Node) rebalance(otherNode *Node) {
	switch {
	case otherNode.balance == -2 && otherNode.Left.balance == -1:
		node.rotateRight(otherNode)
	case otherNode.balance == 2 && otherNode.Right.balance == 1:
		node.rotateLeft(otherNode)
	case otherNode.balance == -2 && otherNode.Left.balance == 1:
		node.rotateLeftRight(otherNode)
	case otherNode.balance == 2 && otherNode.Right.balance == -1:
		node.rotateRightLeft(otherNode)
	}
}

func (node *Node) find(s string) (string, bool) {
	if node == nil {
		return "", false
	}
	if s == node.Value {
		return node.Data, true
	}
	if s < node.Value {
		return node.Left.find(s)
	}
	return node.Right.find(s)
}

func (node *Node) findMax(parent *Node) (*Node, *Node) {
	if node.Right == nil {
		return node, parent
	}
	return node.Right.findMax(node)
}

func (node *Node) replaceNode(parent, replacement *Node) error {
	if node == nil {
		return errors.New("Cannot replace on a nil node")
	}
	if node == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

func (node *Node) delete(s string, parent *Node) error {
	if node == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}
	if s < node.Value {
		return node.Left.delete(s, node)
	} else if s > node.Value {
		return node.Right.delete(s, node)
	}
	// we found the node to be deleted
	// if the node has no children, remove it
	if node.Left == nil && node.Right == nil {
		node.replaceNode(parent, nil)
		return nil
	}
	// if the node has one child, replace the node with its child
	if node.Left == nil {
		node.replaceNode(parent, node.Right)
		return nil
	} else if node.Right == nil {
		node.replaceNode(parent, node.Left)
		return nil
	}
	// the node has two children
	repl, replParent := node.Left.findMax(node)
	node.Value = repl.Value
	node.Data = repl.Data

	return repl.delete(repl.Value, replParent)
}

// Tree with a root node
type Tree struct {
	Root *Node
}

// Insert a value in a tree
func (t *Tree) Insert(value, data string) {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return
	}
	t.Root.insert(value, data)
	if t.Root.balance < -1 || t.Root.balance > 1 {
		t.rebalance()
	}
}

func (t *Tree) rebalance() {
	tmpParent := &Node{Left: t.Root, Value: "tmpParent"}
	tmpParent.rebalance(t.Root)
	t.Root = tmpParent.Left
}

// Find in a tree
func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.find(s)
}

// Delete in a tree
func (t *Tree) Delete(s string) error {
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}
	tmpParent := &Node{Right: t.Root}
	err := t.Root.delete(s, tmpParent)
	if err != nil {
		return err
	}
	if tmpParent.Right == nil {
		t.Root = nil
	}
	return nil
}

// Traverse a tree
func (t *Tree) Traverse(node *Node, f func(*Node)) {
	if node == nil {
		return
	}
	t.Traverse(node.Left, f)
	f(node)
	t.Traverse(node.Right, f)
}

func main() {
	values := []string{"a", "b", "c", "e", "d"}
	data := []string{"alice", "bernard", "carlos", "elen", "daphne"}

	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		tree.Insert(values[i], data[i])
	}

	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting "+s+": ", err)
	}
	fmt.Print("After deleting '" + s + "': ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()
}
