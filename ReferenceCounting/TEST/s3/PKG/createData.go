package PKG

import "fmt"

func NewObject(no string, size int, list *FreeLinkedList) *Object {
	obj := pickupChunk(no, size, list)
	// fmt.Println("obj", obj)
	if obj == nil {
		allocation_fail()
	} else {
		obj.RefCnt = 1
		return obj
	}
	return nil
}

func pickupChunk(no string, size int, list *FreeLinkedList) *Object {
	current := list.Head
	for current != nil {
		var object interface{}
		temp := current.Data
		object = temp
		// fmt.Printf("%T\n", object)
		if ms, ok := object.(*Object); ok {
			//allocate_size可分配的
			oldNo, allocate_size := ms.GetInterface()
			// fmt.Println(oldNo, allocate_size)
			if oldNo != "head" {
				if allocate_size == size {
					list.DeleteNode(object)
					// list.PrintLinkedList()
					return &Object{No: oldNo, Data: make([]byte, size)}
				} else if allocate_size > size {
					list.DeleteNode(object)
					remainingChunk := &Object{
						No:   oldNo,
						Data: make([]byte, allocate_size-size),
					}
					list.InsertAtEnd(remainingChunk)
					// list.PrintLinkedList()
					return &Object{No: no, Data: make([]byte, size)}
				} else {
					allocation_fail()
				}
			}
		}
		current = current.Next
	}
	return nil
}

//书本图3.2的例子
//Num 数量,Size 总需要的空间
var Num int = 4
var Size int = 8
var heap []*Object

func Example1(Num, Size int, list *FreeLinkedList) {
	avgSize := Size / Num
	root := &Object{No: "root", RefCnt: 1, Data: make([]byte, 0)}
	for c, ch := 'A', 'A'+rune(Num); c < ch; c++ {
		heap = append(heap, NewObject(string(c), avgSize, list))
	}
	fmt.Println(root)
	fmt.Println(heap)
}

//BASE_SPACE：初始可分配空间大小，以一个链表节点的形式出场
func InitData(BASE_SPACE int) {
	head := &Node{Data: &Object{No: "head"}}
	node0 := &Object{No: "No.0", Data: make([]byte, 10)}
	list := &FreeLinkedList{Head: head}
	list.InsertAtEnd(node0)
	list.PrintLinkedList()
	Example1(Num, Size, list)
	// NewObject("A", 2, list)
}
