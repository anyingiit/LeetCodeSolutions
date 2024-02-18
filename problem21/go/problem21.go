package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists 合并两个升序链表为一个新的升序链表并返回。
// 思路：
// - 如果两个链表中任何一个为空，则直接返回另一个非空链表，或者两者都为空时返回空。
// - 选择两个链表中较小头节点的链表作为主链表（head），另一个链表作为参照链表（anotherNode）。
// - 遍历参照链表，将其节点按顺序插入到主链表中，直到参照链表为空。
// - 使用两个指针cur和anotherNode，cur始终指向主链表当前操作的节点，anotherNode指向参照链表当前操作的节点。
// - 如果anotherNode指向的节点值小于等于cur的下一个节点值，则将anotherNode节点插入到cur和cur的下一个节点之间。
// - 如果cur的下一个节点为空，或者anotherNode的值大于cur的下一个节点的值，cur指针向前移动。
// - 循环上述过程，直到参照链表遍历完毕。
// 返回合并后的链表头节点。
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 处理两个链表中任何一个为空的情况
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// 初始化head和anotherNode指针，确定初始较小的节点作为合并链表的头
	var head, anotherNode *ListNode
	if list1.Val < list2.Val {
		head = list1
		anotherNode = list2
	} else {
		head = list2
		anotherNode = list1
	}

	// 使用cur指针遍历主链表
	cur := head
	for anotherNode != nil {
		// 如果cur指向的节点的下一个节点为空，直接将剩余的anotherNode链表接在cur后面
		if cur.Next == nil {
			cur.Next = anotherNode
			break
		}

		// 如果anotherNode的值小于等于cur下一个节点的值，执行插入操作
		if anotherNode.Val <= cur.Next.Val {
			// 保存cur下一个节点，将anotherNode插入到cur和cur下一个节点之间
			curNextTemp := cur.Next
			cur.Next = anotherNode

			// 更新anotherNode指针到下一个节点，同时将原anotherNode的下一个节点接在刚插入的节点后面
			anotherNodeNextTemp := anotherNode.Next
			anotherNode.Next = curNextTemp

			// 移动anotherNode和cur指针
			anotherNode = anotherNodeNextTemp
			cur = cur.Next
		} else {
			// 如果anotherNode的值大于cur下一个节点的值，移动cur指针到下一个节点
			cur = cur.Next
		}
	}

	// 返回合并后的链表头节点
	return head
}
