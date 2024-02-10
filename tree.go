package tree

import (
	"fmt"
	"strings"
)

// NewTree creates an empty Tree.
func New(v any) *Tree {

	return newTree(v, 0)
}

func newTree(v any, depth int) *Tree {

	return &Tree{value: v, depth: depth, degree: 0, index: "1", forest: make([]*Tree, 0)}
}

// Tree represents the root node of a Tree
// It contains the subtrees.
type Tree struct {
	value  any     //TODO(@narslan): use generics
	depth  int     //depth measures the distance between a node and the root
	degree int     //degree describes how many Tree the forest have.
	index  string  //index is the node description for displaying Tree in dewey notation.
	forest []*Tree //forest is a set of trees.
}

func (t *Tree) AddTree(v any) *Tree {
	nt := newTree(v, t.depth+1)
	t.degree = t.degree + 1
	nt.index = fmt.Sprintf("%s.%d", t.index, len(t.forest)+1)
	t.forest = append(t.forest, nt)
	return nt
}

func (t *Tree) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("index: %s depth: %d degree:%d value: %v\n", t.index, t.depth, t.degree, t.value))

	if len(t.forest) == 0 {
		return b.String()
	} else {
		for _, v := range t.forest {
			b.WriteString(v.String())
		}
	}
	return b.String()
}

func (t *Tree) Traverse() []any {
	a := make([]any, 0)
	a = append(a, t.value)
	if len(t.forest) == 0 {
		return a
	} else {
		for _, v := range t.forest {
			v.Traverse()
		}
	}
	return a
}

// func (t *Tree) Dewey() string {
// 	var b strings.Builder
// 	b.WriteString(fmt.Sprintf("%d %v\n", t.depth, t.value))

// 	if len(t.forest) == 0 {
// 		return b.String()
// 	} else {
// 		for i, v := range t.forest {
// 			b.WriteString(fmt.Sprintf("%d.%s", i+1, v.Dewey()))
// 		}
// 	}
// 	return b.String()
// }
