package week01

// https://leetcode-cn.com/problems/merge-two-sorted-lists/

type ListNode struct {
  	Val int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	p1 := l1
	p2 := l2

	// 新链表的扫描指针，给他一个保护节点
	newListCursor := &ListNode{}
	// 存储答案
	var ans *ListNode

	// 两个指针同时遍历，哪个小取哪个，如果碰到一个是nil的话，就直接取另一个
	for p1 != nil || p2 != nil{
		if ans == nil{
			ans = newListCursor
		}
		if p1 == nil{
			newListCursor.Next = p2
			newListCursor = newListCursor.Next
			p2 = p2.Next
			continue
		}
		if p2 == nil{
			newListCursor.Next = p1
			newListCursor = newListCursor.Next
			p1 = p1.Next
			continue
		}
		if p1.Val < p2.Val{
			newListCursor.Next = p1
			p1 = p1.Next
		}else{
			newListCursor.Next = p2
			p2 = p2.Next
		}
		newListCursor = newListCursor.Next
	}
	return ans.Next
}