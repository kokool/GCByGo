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

func (obj *Object) UpdatePtr(oldptr *Object, ptr *Object) {
	if ptr == nil {
		return
	}
	if oldptr == ptr {
		return
	}
	ptr.IncRefCnt()
	DecRefCnt(&oldptr)
	obj.children = []*Object{ptr}
	// obj.IncRefCnt()
}

func (obj *Object) IncRefCnt() {
	if obj == nil {
		return
	}
	obj.RefCnt++
}

// func (obj *Object) DecRefCnt() {
// 	if obj == nil {
// 		return
// 	}
// 	obj.RefCnt--
// 	if obj.RefCnt == 0 {
// 		for _, child := range obj.children {
// 			child.DecRefCnt()
// 		}
// 		obj.destroy()
// 	}
// }

func DecRefCnt(obj **Object) {
	if obj == nil {
		return
	}
	(*obj).RefCnt--
	if (*obj).RefCnt == 0 {
		for _, child := range (*obj).children {
			DecRefCnt(&child)
		}
		// (*obj).RefCnt++
		(*obj).destroy()
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
	// if obj.children == nil {
	// 	obj.Data = nil
	// }
	fmt.Println(obj)
	// fmt.Printf("%s has been destroyed\n", obj.No)
}

func main() {
	root := &Object{No: "root", RefCnt: 1, Data: make([]byte, 2)}
	a := &Object{No: "A", RefCnt: 0, Data: make([]byte, 2)}
	b := &Object{No: "B", RefCnt: 0, Data: make([]byte, 2)}
	c := &Object{No: "C", RefCnt: 0, Data: make([]byte, 2)}
	// d := &Object{No: "D", RefCnt: 0, Data: make([]byte, 2)}
	root.AddRef(a)
	root.AddRef(c)
	a.AddRef(b)
	fmt.Println("初始阶段")
	fmt.Println(root)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// fmt.Println(d)
	a.UpdatePtr(b, c)
	fmt.Println("修改阶段")
	fmt.Println(root)
	fmt.Printf("%p", a)
	fmt.Println(a)
	fmt.Printf("%p", b)
	fmt.Println(b)
	fmt.Printf("%p", c)
	fmt.Println(c)
	// fmt.Printf("%p", d)
	// fmt.Println(d)
}

//output:
/*

初始阶段
&{root 0 [0 0] [0xc000114050 0xc0001140f0]}
&{A 1 [0 0] [0xc0001140a0]}
&{B 1 [0 0] []}
&{C 1 [0 0] []}
B has been destroyed
A has been destroyed
修改阶段
&{root 1 [0 0] [0xc000114050 0xc0001140f0]}
0xc000114050&{A 0 [] [0xc0001140f0]}
//我真不知道这个存在引用关系的值为什么会递归成这样。

0xc0001140a0&{B 0 [] []}
0xc0001140f0&{C 2 [0 0] []}
*/

//真正的问题：因为理解错误了书本的update_ptr()函数
