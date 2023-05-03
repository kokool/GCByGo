package PKG

import (
	"errors"
)

type ZCT struct {
	saveObject []*Object
}

func newZCT(size int) *ZCT {
	return &ZCT{
		saveObject: make([]*Object, size),
	}
}

func (zct *ZCT) isFull() bool {
	for _, zobj := range zct.saveObject {
		if zobj == nil {
			return false
		}
	}
	return true
}

func (zct *ZCT) push(obj *Object) error {
	for i, zobj := range zct.saveObject {
		if zobj == nil {
			zct.saveObject[i] = obj
			return nil
		}
	}
	return errors.New("ZCT is full")
}

//删除操作
func (zct *ZCT) remove(obj *Object) error {
	for i, zobj := range zct.saveObject {
		if zobj == obj {
			zct.saveObject[i] = nil
			return nil
		}
	}
	return errors.New("Object not found in ZCT")
}

//目的：就是为了删除计数器为0的对象，早删除才能有空间分配
func (zct *ZCT) scanZCT(list *FreeLinkedList) {
	//在ZCT清干净前，先让heap中涉及到的obj都给清掉
	heap = heap[:zct.removeHeapElWithZCT(heap)]

	//这里这么操作是因为根直接引用的对象有可能是0，对于非直接引用的不用管
	for _, r := range roots {
		r.refCnt++
	}
	for _, zobj := range zct.saveObject {
		// 如果存在计数器值为 0 的对象(注意ZCT因为限定了空间，所以回存在空对象)，则将此对象从 $zct中删除
		//还有一定要将zobj != nil写在前面，不然报错
		if zobj != nil && zobj.refCnt == 0 {
			zct.remove(zobj)
			//对子对象的计数器进行减量操作与回收
			zobj.delete(list)
		}
	}

	for _, r := range roots {
		r.refCnt--
	}
}

//清理掉heap中已经放入到ZCT的对象，注意这是对值引用的修改
func (zct *ZCT) removeHeapElWithZCT(heap []*Object) int {
	for i := 0; i < len(zct.saveObject); i++ {
		for j := 0; j < len(heap); j++ {
			if zct.saveObject[i] == heap[j] {
				heap = append(heap[:j], heap[j+1:]...)
				j--
			}
		}
	}
	return len(heap)
}
