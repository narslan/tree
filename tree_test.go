package tree

import "testing"

func TestTree(t *testing.T) {

	nt := New("A")
	ntb := nt.AddTree("B")
	ntb.AddTree("H")
	ntb.AddTree("J")

	ntc := nt.AddTree("C")
	ntc.AddTree("D")
	nte := ntc.AddTree("E")
	ntc.AddTree("F")

	nte.AddTree("G")

	t.Log(nt)

}
