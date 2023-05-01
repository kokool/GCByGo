//测试失败原因分析：
//理解错误书本中的update_ptr()
//函数用于更新指针 ptr，使其指向对象 obj，同时进行计数器值的增减。
//一开始其实没看懂下面这句话：
//1. 对指针 ptr 新引用的对象（obj）的计数器进行增量操作
//2. 对指针 ptr 之前引用的对象（*ptr）的计数器进行减量操作
//因此写成 func (obj *Object) UpdatePtr(oldptr *Object, ptr *Object)
//表示 obj->oldptr => obj->ptr
package main

import "fmt"

type Object struct {
	No       string
	RefCnt   int
	Data     []byte
	children []*Object
}

//问题1：千万不要理解错误书本的update_ptr()
func (obj *Object) UpdatePtr(oldptr *Object, ptr *Object) {
	if ptr == nil {
		return
	}
	//先增ptr是为了避免出现同一个对象的问题
	ptr.IncRefCnt()
	oldptr.DecRefCnt()
	//问题2：一个节点能够引用多个其他节点，如果可以，这里就需要修改！
	obj.children = []*Object{ptr}
}

func (obj *Object) IncRefCnt() {
	if obj == nil {
		return
	}
	obj.RefCnt++
}

func (obj *Object) DecRefCnt() {
	if obj == nil {
		return
	}
	obj.RefCnt--
	if obj.RefCnt == 0 {
		for _, child := range obj.children {
			child.DecRefCnt()
		}
		//暂时替代reclaim
		obj.destroy()
	}
}

//增加引用，谁被引用谁加计数器
func (obj *Object) AddRef(ptr *Object) {
	if ptr == nil {
		return
	}
	obj.children = append(obj.children, ptr)
	//引用计数法规定，谁被指向，谁的计数器就加值
	ptr.IncRefCnt()
}

func (obj *Object) destroy() {
	obj.children = nil
	obj.Data = nil
	fmt.Printf("%s has been destroyed\n", obj.No)
}

func main() {
	root := &Object{No: "root", RefCnt: 1, Data: make([]byte, 2)}
	a := &Object{No: "A", RefCnt: 0, Data: make([]byte, 2)}
	b := &Object{No: "B", RefCnt: 0, Data: make([]byte, 2)}
	c := &Object{No: "C", RefCnt: 0, Data: make([]byte, 2)}
	root.AddRef(a)
	root.AddRef(c)
	a.AddRef(b)
	fmt.Println("创建书本图3.2(a)中update_prt()函数执行时的情况")
	fmt.Println(root)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	a.UpdatePtr(b, c)
	fmt.Println("最终结果显示正确，执行图3.2(b)的结果")
	fmt.Println(root)
	fmt.Printf("%p", a)
	fmt.Println(a)
	fmt.Printf("%p", b)
	fmt.Println(b)
	fmt.Printf("%p", c)
	fmt.Println(c)
}
