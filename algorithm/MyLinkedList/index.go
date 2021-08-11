package MyLinkedList


type LinkNode struct {
	Value int
	Next *LinkNode
}

type MyLinkedList struct {
	size int
	head *LinkNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{0,nil}
}


func (this *MyLinkedList) Get(index int) int {
	if index >= this.size || index < 0 || this.head == nil {
		return -1
	}
	cur := this.head
	for i := 0; i < index;i++ {
		cur = cur.Next
	}
	return cur.Value
}


func (this *MyLinkedList) AddAtHead(val int)  {
	if this.size == 0 {
		this.head = &LinkNode{val, nil}
	}else {
		next := this.head
		cur := &LinkNode{val, next}
		this.head = cur
	}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int)  {
	// 循环到尾节点
	if this.size == 0 {
		this.head = &LinkNode{val, nil}
	}else {
		cur := this.head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = &LinkNode{val, nil}
	}
	this.size++

}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.size || index < 0 {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}
	prev := this.head
	for i := 0; i < index - 1;i++ {
		prev = prev.Next
	}
	next := prev.Next
	prev.Next = &LinkNode{val, next}
	this.size++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index >= this.size || index < 0 {
		return
	}
	// 如果是头节点
	if index == 0 {
		next := this.head.Next
		this.head = next
	}else {
		prev := this.head
		for i := 0; i < index - 1;i++ {
			prev = prev.Next
		}
		cur := prev.Next
		next := cur.Next
		prev.Next = next
	}
	this.size--
}
