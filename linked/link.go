package linked

type Object interface{}

type Node struct {
	data Object
	next *Node
}

type List struct {
	size int   // 车辆数量
	head *Node // 车头
	tail *Node // 车尾
}

func (list *List) Init() {
	(*list).size = 0   // 此时链表是空的
	(*list).head = nil // 没有车头
	(*list).tail = nil // 没有车尾
}

func (list *List) Put(data Object) bool {
	return list.append(
		&Node{
			data: data,
			next: nil,
		}, )
}

// 获取第一个元素, 并从链表中删除
// ---------- ----------
func (list *List) Push() Object {
	n := list.get(0)
	if n == nil {
		return nil
	} else {
		//
		list.remove(0, n)
		return n.data
	}
}

func (list *List) append(node *Node) bool {
	if node == nil {
		return false
	}

	(*node).next = nil
	// 将新元素放入单链表中
	if (*list).size == 0 {
		(*list).head = node
	} else {
		oldTail := (*list).tail
		(*oldTail).next = node
	}

	// 调整尾部位置，及链表元素数量
	(*list).tail = node // node成为新的尾部
	(*list).size++      // 元素数量增加
	return true
}

// 获取某个位置的元素
// ---------- ----------
func (list *List) get(i int) *Node {
	if i >= (*list).size {
		return nil
	}
	item := (*list).head
	for j := 0; j < i; j++ { // 从head数i个
		item = (*item).next
	}
	return item
}

// 删除元素
// ---------- ----------
func (list *List) remove(i int, node *Node) bool {
	if i >= (*list).size {
		return false
	}

	if i == 0 { // 删除头部
		node = (*list).head
		(*list).head = (*node).next
		if (*list).size == 1 { // 如果只有一个元素，那尾部也要调整
			(*list).tail = nil
		}
	} else {
		preItem := (*list).head
		for j := 1; j < i; j++ {
			preItem = (*preItem).next
		}

		node = (*preItem).next
		(*preItem).next = (*node).next

		if i == ((*list).size - 1) { // 若删除的尾部，尾部指针需要调整
			(*list).tail = preItem
		}
	}
	(*list).size--
	return true
}
