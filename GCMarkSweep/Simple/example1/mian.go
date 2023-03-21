package main

type Object struct {
	value  interface{}
	marked bool
}

type GC struct {
	objects []*Object
}

func (gc *GC) AddObject(obj *Object) {
	gc.objects = append(gc.objects, obj)
}

func (gc *GC) Mark(obj *Object) {
	obj.marked = true
}

func (gc *GC) Sweep() {
	newObjects := []*Object{}
	for _, obj := range gc.objects {
		if obj.marked {
			obj.marked = false
			newObjects = append(newObjects, obj)
		} else {
			// 清除未被标记的对象
			obj = nil
		}
	}
	gc.objects = newObjects
}

/*
在这个示例中，Object是垃圾回收算法的基础数据结构，
每个对象都有一个值（value）、一个标记（marked）。

GC类是垃圾回收器的主要实现，它维护了对象的切片，并提供了三个操作：
添加对象（AddObject）、标记对象（Mark）和清除未标记对象（Sweep）。

当一个对象被创建时，
它会被添加到对象切片中（AddObject），
当它被引用时，它会被标记（Mark），
在垃圾回收过程中，对象切片会被遍历，未被标记的对象将被清除（Sweep）
*/
func main() {
	gc := &GC{}
	gc.AddObject(&Object{value: 1})
	gc.AddObject(&Object{value: 2})
	gc.AddObject(&Object{value: 3})
	gc.Mark(&Object{value: 1})
	gc.Sweep()
}
