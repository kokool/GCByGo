package PKG

import "fmt"

type Object struct {
	no       string
	refCnt   int
	data     []byte
	children []*Object
}

func (obj *Object) getInterface() (string, int) {
	return obj.no, len(obj.data)
}

func (obj *Object) updatePtr(oldptr *Object, ptr *Object, list *FreeLinkedList) {
	if ptr == nil {
		return
	}
	ptr.incRefCnt()
	oldptr.decRefCnt(list)
	obj.children = []*Object{ptr}
}

func (obj *Object) incRefCnt() {
	if obj == nil {
		return
	}
	obj.refCnt++
}

func (obj *Object) decRefCnt(list *FreeLinkedList) {
	if obj == nil {
		return
	}
	obj.refCnt--
	if obj.refCnt == 0 {
		for _, child := range obj.children {
			child.decRefCnt(list)
		}
		obj.reclaim(list)
	}
}

func (obj *Object) AddRef(ptr *Object) {
	if ptr == nil {
		return
	}
	obj.children = append(obj.children, ptr)
	ptr.incRefCnt()
}

func (obj *Object) reclaim(list *FreeLinkedList) {
	obj.children = nil
	fmt.Printf("%s has been reclaimed\n", obj.no)
	//这里就加入空闲链表中
	list.insertAtEnd(obj)
}
