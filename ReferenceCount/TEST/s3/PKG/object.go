package PKG

import "fmt"

type Object struct {
	//注意大小写，表示是否公开属性
	//编号
	No string
	//计数器，书本写的是无符号
	RefCnt int
	//域
	//大小自己指定，但是会受限于空闲链表freeLinkedList中分块的域大小
	Data []byte
	//指针引用，存对象的地址
	children []*Object
}

func (obj *Object) newObj(size int) {

}

//空闲链表表示空的话，没必要继续
func allocation_fail() {
	panic("allocation failed")
}

/*updatePtr*/
/*
update_ptr(ptr, obj){
 inc_ref_cnt(obj)
 dec_ref_cnt(*ptr)
 //更新指针 ptr（指向的A），使其指向对象 obj(被指向的C)，被指向的对象加计数器
 *ptr = obj//下面的代码相反，是令obj->ptr
}
*/
//操作
//1.令ptr(即C).refCnt++变成2，ptr.IncRefCnt()
//2.obj.child(即B).refCnt--至0，然后并入空闲链表
//3.接着让obj(即A).refCnt维持不变，obj->ptr
func (obj *Object) UpdatePtr(ptr *Object) {
	ptr.IncRefCnt()
	// fmt.Println(obj)
	// fmt.Println(ptr)
	obj.DecRefCnt()
	// fmt.Println(obj)
	// fmt.Println(ptr)
	obj.AddRef(ptr)
	// fmt.Println(obj)
	// fmt.Println(ptr)
}

// incRefCnt increments the reference count of an Object
func (obj *Object) IncRefCnt() {
	obj.RefCnt++
}

// decRefCnt decrements the reference count of an Object
// and frees it if the reference count drops to 0.
func (obj *Object) DecRefCnt() {
	obj.RefCnt--
	if obj.RefCnt == 0 {
		//释放引用该对象的其他对象
		for child := range obj.children {
			// child.DecRefCnt()
			fmt.Println(child)
		}
		//reclaim() 函数将 obj 连接到空闲链表
		//reclaim(obj)
		obj.destroy()
	}
}

//释放对象的资源，暂时替代reclaim(obj)
func (obj *Object) destroy() {
	obj.Data = nil
	obj.children = obj.children[:len(obj.children)-1]
}

//添加引用关系：obj->ptr
func (obj *Object) AddRef(ptr *Object) {
	obj.children = append(obj.children, ptr)
}

//删除引用关系(可能用不上)
// func (obj *Object) RemoveRef(ptr *Object) {
// 	delete(obj.children, ptr)
// }
