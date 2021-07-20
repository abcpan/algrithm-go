package algorithm


func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var stack = make([]int, 0)
	var p1 = head
	// 取出所有元素
	var p2 = head
	for p1 != nil {
		stack = append(stack, p1.Val)
		p1 = p1.Next
	}
	for p2 != nil {
		tailIndex := len(stack) -1
		last := stack[tailIndex]
		stack = stack[0: tailIndex]
		if last != p2.Val {
			return false
		}
		p2 = p2.Next
	}
	return true
}