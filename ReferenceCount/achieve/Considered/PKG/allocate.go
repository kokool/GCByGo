package PKG

func allocation_fail() {
	panic("allocation failed")
}

func newObject(no string, size int, list *FreeLinkedList) *Object {
	obj := pickupChunk(no, size, list)
	if obj == nil {
		allocation_fail()
	} else {
		//注意我这里跟书本的伪代码不一样，因为我认为一开始新建的默认为0，而只要被引用了才改变refCnt
		obj.refCnt = 0
		return obj
	}
	return nil
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
					return &Object{no: no, data: make([]byte, size)}
				} else if allocate_size > size {
					list.deleteNode(object)
					remainingChunk := &Object{
						no:   oldNo,
						data: make([]byte, allocate_size-size),
					}
					list.insertAtEnd(remainingChunk)
					return &Object{no: no, data: make([]byte, size)}
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
	newNode := &Object{no: "No.0", data: make([]byte, totalSize)}
	list.insertAtEnd(newNode)
	list.printLinkedList()
}
