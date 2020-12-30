package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

func (l linkedList) Display() {
	for i := 0; i < l.length-1; i++ {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	fmt.Printf("%v \n", l.head.data)
	fmt.Println("length of list: ", l.length)
}
func (l *linkedList) PushBack(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}
func addTwoNode(l1 *node, l2 *node, carry int) (*node, int) {
	cur1, cur2 := l1, l2
	var head *node
	num1, num2 := 0, 0
	num1 = cur1.data
	num2 = cur2.data
	newNode := new(node)
	newNode.data = (num1 + num2 + carry) % 10
	head = newNode
	carry = (num1 + num2 + carry) / 10
	return head, carry
}
func addTwoList(l1 *linkedList, l2 *linkedList) linkedList {
	cur1 := l1.head
	cur2 := l2.head
	list := linkedList{}
	carry := 0
	for i := 0; i < l1.length; i++ {
		newNode := new(node)
		newNode, carry = addTwoNode(cur1, cur2, carry)
		fmt.Println(carry)
		list.PushBack(newNode)
		l1.head = l1.head.next
		l2.head = l2.head.next
		cur1 = l1.head
		cur2 = l2.head
	}
	return list
}
func main() {
	list1 := linkedList{}
	node1_1 := &node{data: 4}
	node1_2 := &node{data: 5}
	node1_3 := &node{data: 6}
	list1.PushBack(node1_1)
	list1.PushBack(node1_2)
	list1.PushBack(node1_3)
	list2 := linkedList{}
	node2_1 := &node{data: 4}
	node2_2 := &node{data: 7}
	node2_3 := &node{data: 6}
	list2.PushBack(node2_1)
	list2.PushBack(node2_2)
	list2.PushBack(node2_3)
	list1.Display()
	list2.Display()
	list3 := addTwoList(&list1, &list2)
	list3.Display()
}
