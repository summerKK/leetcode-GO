package __3

type Node struct {
	val  interface{}
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Val() interface{} {
	return n.val
}

type Linked struct {
	current *Node
	n       int
}

func (l *Linked) Init() {
	l.current = &Node{}
}

// 增加新元素
func (l *Linked) Add(item interface{}) {
	n := &Node{
		val: item,
	}
	n.next = l.current
	l.current = n
	l.n++
}

//  先进后出
func (l *Linked) Del() interface{} {
	if l.Size() <= 0 {
		return nil
	}
	n := l.current
	l.current = l.current.next
	l.n--

	return n.Val()
}

func (l *Linked) Size() int {
	return l.n
}

func (l *Linked) Loop() <-chan interface{} {
	c := make(chan interface{})
	go func() {
		for n := l.current; n.next != nil; n = n.next {
			c <- n.val
		}
		close(c)
	}()

	return c
}

func (l *Linked) Peek() interface{} {
	return l.current.val
}
