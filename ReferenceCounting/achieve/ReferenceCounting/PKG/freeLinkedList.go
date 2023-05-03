package PKG

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type FreeLinkedList struct {
	head *Node
}

func (list *FreeLinkedList) insertAtEnd(data interface{}) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *FreeLinkedList) deleteNode(data interface{}) {
	if list.head == nil {
		return
	}
	if list.head.data == data {
		list.head = list.head.next
		return
	}
	prev := list.head
	current := list.head.next
	for current != nil {
		if current.data == data {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.data)
}

func (list *FreeLinkedList) printLinkedList() {
	current := list.head
	for current != nil {
		fmt.Printf("%v->", current.data)
		current = current.next
		if current == nil {
			fmt.Print("nil\n")
		}
	}
}
