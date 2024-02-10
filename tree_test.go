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

func TestTreeTraverse(t *testing.T) {

	nt := New("A")
	ntb := nt.AddTree("B")
	ntb.AddTree("H")
	ntb.AddTree("J")

	ntc := nt.AddTree("C")
	ntc.AddTree("D")
	nte := ntc.AddTree("E")
	ntc.AddTree("F")

	nte.AddTree("G")

	s := []string{"A", "B", "H", "J", "C", "D", "E", "G", "F"}
	w := nt.Traverse()

	if len(w) != len(s) {
		t.Fatalf("expected %d elements of tree but got %d", len(s), len(w))
	}
	for i := range s {
		if s[i] != w[i] {
			t.Fatalf("expected %s got %s", s[i], w[i])
		}

	}

}
