//实现分块合并的目标
//下列情况，模拟在已经有多个对象进入到空闲链表中。
package main

import (
	L "GCByGo/ReferenceCount/TEST/s3/PKG"
	"fmt"
)

type No struct {
	count int
	data  []byte
}

func (obj *No) GetInterface() (int, int) {
	return obj.count, len(obj.data)
}

func main() {
	node1 := &No{count: 10, data: make([]byte, 2)}
	node2 := &No{count: 10, data: make([]byte, 2)}
	node3 := &No{count: 10, data: make([]byte, 2)}
	node4 := &No{count: 10, data: make([]byte, 2)}
	head1 := &L.Node{Data: &No{count: 1}}
	list := &L.FreeLinkedList{Head: head1}
	list.InsertAtEnd(node1)
	list.InsertAtEnd(node2)
	list.InsertAtEnd(node3)
	list.InsertAtEnd(node4)
	list.PrintLinkedList()
	current := list.Head
	var totalSize int = 0
	for current != nil {
		var object interface{}
		temp := current.Data
		object = temp
		if ms, ok := object.(*No); ok {
			//allocate_size可分配的
			oldNo, size := ms.GetInterface()
			if oldNo != 1 {
				list.DeleteNode(object)
				totalSize += size
			}
		}
		current = current.Next
	}
	fmt.Println(totalSize)
	newNode := &No{count: 10, data: make([]byte, totalSize)}
	list.InsertAtEnd(newNode)
	list.PrintLinkedList()
}
