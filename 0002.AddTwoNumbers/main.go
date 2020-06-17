package main

import "fmt"

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	result := addTwoNumbers(
		&ListNode{5, nil},
		&ListNode{5, nil})
	//&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
	//&ListNode{5, &ListNode{6, &ListNode{4, nil}}})

	for i := result; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := 0, 0
	node1, node2 := l1, l2
	flag1, flag2 := true, true
	var result, pre *ListNode
	sum, jin, mod := 0, 0, 0
	for {

		if node1 != nil {
			num1 = node1.Val
			node1 = node1.Next
		} else {
			flag1 = false
			num1 = 0
		}

		if node2 != nil {
			num2 = node2.Val
			node2 = node2.Next
		} else {
			flag2 = false
			num2 = 0
		}
		sum = num1 + num2 + jin

		if flag1 || flag2 || sum > 0 {
			jin = sum / 10
			mod = sum % 10
			node := &ListNode{mod, nil}
			if result == nil {
				result = node
				pre = node
			} else {
				pre.Next = node
				pre = node
			}
		} else {
			break
		}
	}
	return result
}
