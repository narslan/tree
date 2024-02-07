package tree

import "testing"

func TestTree(t *testing.T) {

	nt := newTree("A", 1)
	ntb := nt.addTree("B")
	ntb.addTree("H")
	ntb.addTree("J")

	ntc := nt.addTree("C")
	ntc.addTree("D")
	nte := ntc.addTree("E")
	ntc.addTree("F")

	nte.addTree("G")

	t.Log(nt)

}
