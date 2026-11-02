package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(*tree.Tree)
	walk = func(node *tree.Tree) {
		if node == nil {
			return
		}
		walk(node.Left)
		ch <- node.Value
		walk(node.Right)
	}
	walk(t)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println("Same(tree.New(1), tree.New(1)): ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2)): ", Same(tree.New(1), tree.New(2)))
}
