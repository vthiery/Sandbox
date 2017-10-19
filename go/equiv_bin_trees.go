package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	InternalWalk(t, ch)
	close(ch)
}

func InternalWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		InternalWalk(t.Left, ch)
		ch <- t.Value
		InternalWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Print(v)
	}
}

