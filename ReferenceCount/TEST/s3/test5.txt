package main

import (
	"fmt"
)

type Object struct {
	refCount int
	children []*Object
}

func NewObject() *Object {
	obj := &Object{refCount: 1}
	return obj
}

func (obj *Object) AddChild(child *Object) {
	obj.children = append(obj.children, child)
}

func (obj *Object) UpdatePtr(ptr **Object, newObj *Object) {
	obj.IncRefCount(newObj)
	obj.DecRefCount(*ptr)
	*ptr = newObj
}

func (obj *Object) IncRefCount(child *Object) {
	child.refCount++
}

func (obj *Object) DecRefCount(child *Object) {
	child.refCount--
	if child.refCount == 0 {
		for _, c := range child.children {
			obj.DecRefCount(c)
		}
		obj.Reclaim(child)
	}
}

func (obj *Object) Reclaim(child *Object) {
	fmt.Println("Object reclaimed")
}

func main() {
	// 创建三个对象
	obj1 := NewObject()
	obj2 := NewObject()
	obj3 := NewObject()
	fmt.Println("root", obj1)
	fmt.Println("a", obj2)
	fmt.Println("c", obj3)
	// 将obj1的一个子对象指向obj2
	obj1.AddChild(obj2)
	fmt.Println("root", obj1)
	fmt.Println("a", obj2)
	fmt.Println("c", obj3)
	// 将obj1的一个子对象指向obj3
	obj1.AddChild(obj3)
	fmt.Println("root", obj1)
	fmt.Println("a", obj2)
	fmt.Println("c", obj3)
	// 更新obj2的引用，指向obj3
	obj2.UpdatePtr(&obj2, obj3)
	fmt.Println("root", obj1)
	fmt.Println("a", obj2)
	fmt.Println("c", obj3)
}
