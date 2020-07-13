package main

import "fmt"

func main() {
	fmt.Println(mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}))
}

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	all := &ListNode{}
	var result *ListNode
	for {
		if l1 != nil && l2 != nil {
			v1 := l1.Val
			v2 := l2.Val
			if v1 < v2 {
				all.Next = l1
				if result == nil {
					result = l1
				}
				all = all.Next
				l1 = l1.Next
			} else {
				all.Next = l2
				if result == nil {
					result = l2
				}
				all = all.Next
				l2 = l2.Next
			}
		} else if l1 == nil {
			all.Next = l2
			if result == nil {
				result = l2
			}
			break
		} else if l2 == nil {
			all.Next = l1
			if result == nil {
				result = l1
			}
			break
		}

	}
	return result
}
