// using the source here .ttps://gist.github.com/N02870941/33d20ab51abf7f68474f3348a6147c70
package bst

import (
	"fmt"
	"strconv"
)

//------------------------------------------------------------------------------

type Node struct {
	data  int
	left  *Node
	right *Node
}

//------------------------------------------------------------------------------

type BST struct {
	size int
	root *Node
}

//------------------------------------------------------------------------------

func (tree *BST) insert(data int) {

	tree.root = tree.add(data, tree.root)
}

//------------------------------------------------------------------------------

func (tree *BST) add(data int, node *Node) *Node {

	if node == nil {

		tree.size++

		return &Node{data: data}
	}

	if data < node.data {

		node.left = tree.add(data, node.left)

	} else if data > node.data {

		node.right = tree.add(data, node.right)
	}

	return node
}

//------------------------------------------------------------------------------

func (tree *BST) remove(data int) {

	tree.root = tree.delete(data, tree.root)
}

//------------------------------------------------------------------------------

func (tree *BST) delete(data int, node *Node) *Node {

	if node == nil {

		return nil
	}

	if data < node.data {

		node.left = tree.delete(data, node.left)

	} else if data > node.data {

		node.right = tree.delete(data, node.right)

	} else {

		if node.left != nil && node.right != nil {

			min := tree.min(node.right)
			node.data = min.data
			node.right = tree.delete(min.data, node.right)

		} else {

			if node.left == nil && node.right == nil {

				node = nil

			} else if node.left == nil {

				node = node.right

			} else {

				node = node.left
			}

			tree.size--
		}
	}

	return node
}

//------------------------------------------------------------------------------

func (tree *BST) min(node *Node) *Node {

	for node.left != nil {

		node = node.left
	}

	return node
}

//------------------------------------------------------------------------------

func (tree BST) inorder() {

	inorder := tree.in(tree.root, "inorder: ")

	fmt.Println(inorder)
}

//------------------------------------------------------------------------------

func (tree BST) in(node *Node, str string) string {

	if node != nil {
		str = tree.in(node.left, str)
		str = str + strconv.Itoa(node.data) + " "
		str = tree.in(node.right, str)
	}

	return str
}

//------------------------------------------------------------------------------

func (tree BST) preorder() {

	preorder := tree.pre(tree.root, "preorder: ")

	fmt.Println(preorder)
}

//------------------------------------------------------------------------------

func (tree BST) pre(node *Node, str string) string {

	if node != nil {
		str = str + strconv.Itoa(node.data) + " "
		str = tree.pre(node.left, str)
		str = tree.pre(node.right, str)
	}

	return str
}

//------------------------------------------------------------------------------

func (tree BST) postorder() {

	postorder := tree.post(tree.root, "postorder: ")

	fmt.Println(postorder)
}

//------------------------------------------------------------------------------

func (tree BST) post(node *Node, str string) string {

	if node != nil {
		str = tree.post(node.left, str)
		str = tree.post(node.right, str)
		str = str + strconv.Itoa(node.data) + " "
	}

	return str
}

//------------------------------------------------------------------------------

func (tree BST) display() {
	fmt.Println("size", tree.size)

	tree.preorder()
	tree.inorder()
	tree.postorder()
	tree.treeString()
}

//------------------------------------------------------------------------------

func (tree BST) treeString() {

	str := tree.toTreeString("", true, "", tree.root)

	fmt.Println(str)
}

//------------------------------------------------------------------------------

func (tree BST) toTreeString(prefix string, top bool, str string, node *Node) string {

	left := new(Node)
	right := new(Node)

	if node == nil {

		return ""
	}

	left = node.left
	right = node.right

	if right != nil {

		temp := tree.path(top, ""+prefix, "│   ", "    ")
		str = tree.toTreeString(temp, false, str, right)
	}

	str = tree.path(top, str+prefix, "└──", "┌──")
	str = str + " " + strconv.Itoa(node.data) + "\n"

	if left != nil {

		temp := tree.path(top, ""+prefix, "    ", "│   ")
		str = tree.toTreeString(temp, true, str, left)
	}

	return str
}

//------------------------------------------------------------------------------

func (tree BST) path(condition bool, str string, choice1 string, choice2 string) string {

	if condition {

		str += choice1

	} else {

		str += choice2
	}

	return str
}
