package PKG

import "fmt"

func allocation_fail() {
	panic("allocation failed")
}

func newObject(no string, size int, list *FreeLinkedList, zct *ZCT) *Object {
	obj := pickupChunk(no, size, list)
	if obj == nil {
		//对哦，如果你都没空间了，那么obj自然就返回nil了
		// fmt.Println("newObject", obj)
		zct.scanZCT(list)
		obj = pickupChunk(no, size, list)
		if obj == nil {
			allocation_fail()
		}
	}
	//无论如何域的第一个位置必须是自己的地址
	obj.data[0] = obj
	obj.refCnt = 1
	return obj
}

func pickupChunk(no string, size int, list *FreeLinkedList) *Object {
	current := list.head
	for current != nil {
		var object interface{}
		temp := current.data
		object = temp
		if ms, ok := object.(*Object); ok {
			oldNo, allocate_size := ms.getInterface()
			if oldNo != "head" {
				if allocate_size == size {
					list.deleteNode(object)
					return &Object{no: no, data: make([]*Object, size)}
				} else if allocate_size > size {
					list.deleteNode(object)
					remainingChunk := &Object{
						no:   oldNo,
						data: make([]*Object, allocate_size-size),
					}
					list.insertAtEnd(remainingChunk)
					return &Object{no: no, data: make([]*Object, size)}
				} else {
					allocation_fail()
				}
			}
		}
		current = current.next
	}
	return nil
}

func mergeChunk(list *FreeLinkedList) {
	current := list.head
	var totalSize int = 0
	for current != nil {
		var object interface{}
		temp := current.data
		object = temp
		if ms, ok := object.(*Object); ok {
			//allocate_size可分配的
			oldNo, size := ms.getInterface()
			if oldNo != "head" {
				list.deleteNode(object)
				totalSize += size
			}
		}
		current = current.next
	}
	if totalSize != 0 {
		newNode := &Object{no: "No.0", data: make([]*Object, totalSize)}
		list.insertAtEnd(newNode)
	}
	fmt.Printf("此时的空闲链表的情况：")
	list.printLinkedList()
}
