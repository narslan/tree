package cuesheet

import (
	"fmt"
	"strings"
)

// NewTree creates an emty tree.
func NewTree(v any) *tree {

	return newTree(v, 0)
}

func newTree(v any, depth int) *tree {

	return &tree{value: v, depth: depth, degree: 0, index: "1", forest: make([]*tree, 0)}
}

// tree represents the root node of a tree
// It contains the subtrees.
type tree struct {
	value  any     //TODO(@narslan): use generics
	depth  int     //depth measures the distance between a node and the root
	degree int     //degree describes how many tree the forest have.
	index  string  //index is the node description for displaying tree in dewey notation.
	forest []*tree //forest is a set of trees.
}

func (t *tree) addTree(v any) *tree {
	nt := newTree(v, t.depth+1)
	t.degree = t.degree + 1
	nt.index = fmt.Sprintf("%s.%d", t.index, len(t.forest)+1)
	t.forest = append(t.forest, nt)
	return nt
}

func (t *tree) String() string {
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

// func (t *tree) Dewey() string {
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
