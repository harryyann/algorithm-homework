package week03

type ListNode struct {
    Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	// 从数组中取出两个进行合并，把合并完的新链表插到数组中，直到数组剩下最后一个链表
	for len(lists) != 1{
		l1 := lists[0]
		l2 := lists[1]
		lists = lists[2:]
		newList := mergeTwoLists(l1, l2)
		lists = append(lists, newList)
	}
	return lists[0]
}

// 递归法合并两个链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode)*ListNode{
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}

	var head *ListNode

	if l1.Val <= l2.Val{
		head = l1
		head.Next = mergeTwoLists(l1.Next, l2)
	}else{
		head = l2
		head.Next = mergeTwoLists(l1, l2.Next)
	}
	return head
}
