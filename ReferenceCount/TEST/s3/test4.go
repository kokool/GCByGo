//测试失败原因分析：应该就是递归的条件判断问题
//明明节点A还在被rott引用，但是却refCnt=0
package main

import "fmt"

type Object struct {
	No       string
	RefCnt   int
	Data     []byte
	children []*Object
}

func (obj *Object) UpdatePtr(ptr *Object) {
	if ptr == nil {
		return
	}
	ptr.IncRefCnt()
	obj.DecRefCnt()
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
		obj.destroy()
	}
	//obj.RefCnt++
}

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
	root := &Object{No: "root", RefCnt: 0, Data: make([]byte, 2)}
	a := &Object{No: "A", RefCnt: 0, Data: make([]byte, 2)}
	b := &Object{No: "B", RefCnt: 0, Data: make([]byte, 2)}
	c := &Object{No: "C", RefCnt: 0, Data: make([]byte, 2)}
	root.AddRef(a)
	root.AddRef(c)
	a.AddRef(b)
	fmt.Println("初始阶段")
	fmt.Println(root)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	a.UpdatePtr(c)
	fmt.Println("修改阶段")
	fmt.Println(root)
	fmt.Printf("%p", a)
	fmt.Println(a)
	fmt.Printf("%p", b)
	fmt.Println(b)
	fmt.Printf("%p", c)
	fmt.Println(c)
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

//我的解决思路：在减量操作回溯refCnt,之后统一查看那些没有被引用的对象
//将它们统一处理掉即可。
