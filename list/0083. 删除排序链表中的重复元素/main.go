package main

import "fmt"

func main() {
	fmt.Println(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}))

	fmt.Println(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}))

	fmt.Println(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var result string
	c := l
	for c != nil {
		result = fmt.Sprintf("%s->%d", result, c.Val)
		c = c.Next
	}
	return result
}

/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func deleteDuplicates1(head *ListNode) *ListNode {
	nodeSet := make(map[int]struct{})
	var pre *ListNode
	for n := head; n != nil; n = n.Next {
		if _, ok := nodeSet[n.Val]; ok {
			pre.Next = n.Next
		} else {
			nodeSet[n.Val] = struct{}{}
			pre = n
		}
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	var pre *ListNode
	for n := head; n != nil; n = n.Next {
		if pre == nil || pre.Val != n.Val {
			pre = n
			continue
		}
		if pre.Val == n.Val {
			pre.Next = n.Next
		}
	}
	return head
}
