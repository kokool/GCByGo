package main

import (
	"fmt"

	L "GCByGo/ReferenceCount/TEST/s2/PKG"
)

func main() {
	node1 := &L.Node{Data: 10}
	node2 := &L.Node{Data: 20}
	node3 := &L.Node{Data: 30}
	node4 := &L.Node{Data: 40}

	list := &L.LinkedList{Head: node1}
	list.InsertAtBeginning(node2)
	list.InsertAtEnd(node3)
	list.InsertAtEnd(node4)

	list.PrintLinkedList() //20->10->30->40->nil
	list.DeleteNode(node2)
	fmt.Println(list.FindNode(node3))  //in func String： Data:30
	fmt.Println(*list.FindNode(node3)) //in func String：{Data Next}
	fmt.Println(list.FindNode(node2))  //nil
	list.PrintLinkedList()             //10->30->40->nil
}
