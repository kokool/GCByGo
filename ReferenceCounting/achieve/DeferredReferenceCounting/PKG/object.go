package PKG

import (
	"errors"
	"fmt"
)

type Object struct {
	//编号
	no string
	//计数器，书本说是无符号的
	refCnt int
	//域，限定空间，第一个是自己的地址，第二个是引用的地址。
	data []*Object
}

//存放根直接引用的对象的地址
var roots []*Object

//专门存放对象
var heap []*Object

//添加引用关系
func (obj *Object) AddRef(ref *Object) error {
	//正常情况
	for i, child := range obj.data {
		if child == nil {
			obj.data[i] = ref
			return nil
		}
	}
	//书本提到的特殊情况：根直接引用
	if obj.no == "roots" {
		roots = append(roots, ref)
	}
	return errors.New("data の size not enough")
}

func (obj *Object) getInterface() (string, int) {
	return obj.no, len(obj.data)
}

func (obj *Object) updatePtr(oldptr *Object, ptr *Object, list *FreeLinkedList, zct *ZCT) error {
	if ptr == nil {
		return errors.New("ptr is null")
	}
	//引用被增加
	ptr.incRefCnt()
	oldptr.decRefCnt(zct, list)
	for i, child := range obj.data {
		if child == oldptr {
			obj.data[i] = ptr
			return nil
		}
	}
	//without oldptr
	return errors.New("Without Reference")
}

func (obj *Object) incRefCnt() {
	if obj == nil {
		return
	}
	obj.refCnt++
}

func (obj *Object) decRefCnt(zct *ZCT, list *FreeLinkedList) {
	if obj == nil {
		return
	}
	obj.refCnt--
	if obj.refCnt == 0 {
		if zct.isFull() == true {
			zct.scanZCT(list)
		}
		zct.push(obj)
	}
}

func (obj *Object) delete(list *FreeLinkedList) {
	//for child in obj.ref
	//因为第一个位置放的是自己的地址
	// fmt.Println("delete", obj)
	//对 obj 的子对象的计数器进行减量操作，对计数器值变成 0 的对象执行 delete() 函数，最后回收 obj。
	for i := 1; i < len(obj.data); i++ {
		//你多分配了些域的空间，但是没有用，那么就没必要进行下面的操作
		// fmt.Println("delete-child", obj.data[i])
		if obj.data[i] != nil {
			obj.data[i].refCnt--
			if obj.data[i].refCnt == 0 {
				obj.data[i].delete(list)
			}
		}
	}
	obj.reclaim(list)

}

//对子对象的计数器进行减量操作与回收
func (obj *Object) reclaim(list *FreeLinkedList) {
	//代表引用关系清空，代表消除自己的地址
	for i, child := range obj.data {
		if child != nil {
			obj.data[i] = nil
		}
	}
	fmt.Printf("%s has been reclaimed\n", obj.no)
	//这里就加入空闲链表中
	list.insertAtEnd(obj)
	mergeChunk(list)
}
