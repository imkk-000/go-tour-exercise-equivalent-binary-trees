package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel channelInt.
func Walk(t *tree.Tree, channelInt chan int) {
	if t.Left != nil {
		Walk(t.Left, channelInt)
	}
	if t.Right != nil {
		Walk(t.Right, channelInt)
	}
	channelInt <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	resultTree1 := walkInTree(t1)
	resultTree2 := walkInTree(t2)
	resultTreeLength := len(resultTree1)
	if resultTreeLength != len(resultTree2) {
		return false
	}
	for index := 0; index < resultTreeLength; index++ {
		if resultTree1[index] != resultTree2[index] {
			return false
		}
	}
	return true
}

func walkInTree(tree *tree.Tree) []int {
	result := []int{}
	channelInt := make(chan int)
	go func() {
		Walk(tree, channelInt)
		close(channelInt)
	}()
	for {
		v, ok := <-channelInt
		if !ok {
			sort.Ints(result)
			break
		}
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
