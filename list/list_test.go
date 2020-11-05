package list

import "testing"

func TestLinkReserve(t *testing.T) {
	ll := NewList()

	ll.Push(1)
	ll.Push(3)
	ll.Push(5)

	data := &elem{Data: 2}
	ll.InsertIndex(1, data)
	ll.InsertElem(data, &elem{Data: 22})
	for e := ll.Head; e != nil; e = e.Next {
		t.Log(e)
	}
}

func Test(t *testing.T) {

}
