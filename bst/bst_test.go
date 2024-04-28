package bst

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	tree := &BST{}

	elements := [7]int{50, 30, 20, 40, 70, 60, 80}

	for _, element := range elements {

		tree.insert(element)
	}
	fmt.Println("Elements:", elements)

	tree.display()
}
