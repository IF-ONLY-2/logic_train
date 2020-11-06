package list

type elem struct {
	Data interface{}
	Next *elem
}

type List struct {
	Head *elem
	len  int
}

// NewList create new list
func NewList() List {
	l := List{
		Head: &elem{Data: nil, Next: nil},
		len:  0,
	}
	return l
}

func (l *List) Push(data interface{}) {
	e := l.Head
	for e.Next != nil {
		e = e.Next
	}
	n := &elem{Data: data, Next: nil}
	e.Next = n
	l.len++
	return
}

func (l *List) InsertIndex(n int, data *elem) {
	if n >= l.len {
		panic("out of range")
	}
	e := l.Head
	for i := 0; i < n; i++ {
		e = e.Next
	}
	data.Next = e.Next
	e.Next = data
	l.len++
	return
}

func (l *List) InsertElem(pre *elem, data *elem) {
	data.Next = pre.Next
	pre.Next = data
	l.len++
	return
}

func (l *List) Reserve() {
	e := l.Head.Next
	if e == nil {
		return
	}
	n1 := e.Next
	e.Next = nil
	for n1 != nil {
		n2 := n1.Next
		n1.Next = e
		e = n1
		n1 = n2
	}
	l.Head.Next = e
	return
}
