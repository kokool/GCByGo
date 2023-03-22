package main

import (
	"fmt"
)

type Object struct {
	id       string //标号
	marked   bool   //删除标记
	data     int    //数据或者被引用对象
	refCount int    //引用计数器
}

//ZCT（Zero Count Table）
var deferredDeleteList []*Object

//引用列表
var refs []*Object

func NewObject(d int) *Object {
	obj := &Object{data: d, refCount: 1}
	return obj
}

func (obj *Object) AddRef() {
	obj.refCount++
	fmt.Println(obj.id, obj.refCount)
}

func (obj *Object) Release() {
	obj.refCount--
	if obj.refCount == 0 {
		fmt.Println(obj.id + "为0，被放入到ZCT中")
		deferredDeleteList = append(deferredDeleteList, obj)
	}
}

func mark(obj *Object) {
	if obj == nil || obj.marked {
		return
	}
	obj.marked = true
	mark(findRef())
}

func findRef(refs []*Object, ref *Object) int {
	// for i, r := range *refs {
	// 	if r == ref {
	// 		return i
	// 	}
	// }
	// return -1
	return nil
}
func main() {

}
