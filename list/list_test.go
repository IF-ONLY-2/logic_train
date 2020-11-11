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
	for e := ll.Head.Next; e != nil; e = e.Next {
		t.Log(e)
	}

	ll.Reserve()
	for e := ll.Head.Next; e != nil; e = e.Next {
		t.Log(e)
	}
}

func TestArr(t *testing.T) {
	arr := []int{1, 2}

	t.Log(arr[:1])
	t.Log(arr[1:])

}
