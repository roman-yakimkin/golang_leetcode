package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	var lastElem *ListNode
	var newVal = func(x int, y int) int {
		if x == -1000 {
			return y
		}
		if y == -1000 {
			return x
		}
		if x < y {
			return x
		}
		return y
	}

	for list1 != nil || list2 != nil {
		val1, val2 := -1000, -1000
		if list1 != nil {
			val1 = list1.Val
		}
		if list2 != nil {
			val2 = list2.Val
		}
		nVal := newVal(val1, val2)

		if val1 >= -100 || val2 >= -100 {
			newElem := ListNode{
				Val: nVal,
			}
			if result == nil {
				result = &newElem
			}

			if lastElem == nil {
				lastElem = &newElem
			} else {
				lastElem.Next = &newElem
				lastElem = lastElem.Next
			}

			if nVal == val1 {
				list1 = list1.Next
			} else if nVal == val2 {
				list2 = list2.Next
			}
		}
	}

	return result
}
