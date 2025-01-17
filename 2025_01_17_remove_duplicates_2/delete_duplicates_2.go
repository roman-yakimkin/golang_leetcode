package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	dupSet := make(map[int]struct{})

	for {
		if p == nil || p.Next == nil {
			break
		}
		if p.Val == p.Next.Val {
			dupSet[p.Val] = struct{}{}
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	p = head
	var result *ListNode
	var lastElem *ListNode
	for p != nil {
		if _, ok := dupSet[p.Val]; !ok {
			newElem := ListNode{
				Val: p.Val,
			}
			if result == nil {
				result = &newElem
			}
			if lastElem != nil {
				lastElem.Next = &newElem
			}
			lastElem = &newElem
		}
		p = p.Next
	}

	return result
}
