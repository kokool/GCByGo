package PKG

import "fmt"

//1.定义链表节点（对象/分块）结构体：
// data
// 每个链表节点应该包括存储数据的变量：object{size，refCnt}
// 指向下一个节点的指针
type Node struct {
	//注意大写，公共，不然导包时，用不了
	Data interface{}
	Next *Node
}

// 2.定义链表结构体：链表应该包含一个头节点，以及一些基本的操作方法，如插入、删除和搜索。
type FreeLinkedList struct {
	Head *Node
}

//在链表头部插入节点（鸡肋的操作）
// func (list *FreeLinkedList) InsertAtBeginning(data interface{}) {
// 	newNode := &Node{Data: data}
// 	if list.Head == nil {
// 		list.Head = newNode
// 	} else {
// 		newNode.Next = list.Head
// 		list.Head = newNode
// 	}
// }

//在链表尾部插入节点，注意格式要变成*PKG.Object,而不是*PKG.Node
func (list *FreeLinkedList) InsertAtEnd(data interface{}) {
	// fmt.Printf("i %T\n", data)
	newNode := &Node{Data: data}
	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

//删除链表中的某个节点
func (list *FreeLinkedList) DeleteNode(data interface{}) {
	// fmt.Println("head", list.Head.Data, data)
	if list.Head == nil {
		return
	}
	if list.Head.Data == data {
		list.Head = list.Head.Next
		return
	}
	prev := list.Head
	current := list.Head.Next
	// fmt.Println("pointer", prev, current)
	for current != nil {
		// fmt.Println("now", current.Data, data)
		// fmt.Printf("d %T\n", current.Data)
		// fmt.Printf("d %T\n", data)
		if current.Data == data {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}
}

//在链表中查找节点
func (list *FreeLinkedList) FindNode(target interface{}) *Node {
	current := list.Head
	for current != nil {
		if current.Data == target {
			return current
		}
		current = current.Next
	}
	return nil
}

//将链表节点定义成string方法，不需要直接调用
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Data)
}

//输出链表信息
func (list *FreeLinkedList) PrintLinkedList() {
	current := list.Head
	for current != nil {
		fmt.Printf("%v->", current.Data)
		current = current.Next
		if current == nil {
			fmt.Print("nil\n")
		}
	}
}
