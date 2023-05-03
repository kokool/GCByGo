package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Object struct {
	id       string //标号
	refCount int    //引用计数器
}

func (obj *Object) AddRef() {
	obj.refCount++
	fmt.Println(obj.id, obj.refCount)
}

func (obj *Object) Release() {
	obj.refCount--
	if obj.refCount <= 0 {
		fmt.Println(obj.id + "对象已经被释放")
		runtime.SetFinalizer(obj, nil)
	}
	fmt.Println(obj.id, obj.refCount)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	//创建两个对象
	obj1 := &Object{id: "obj1"}
	obj2 := &Object{id: "obj2"}

	obj1.AddRef()
	obj2.AddRef()
	obj1.Release()

	//这里有可能不会生效，因为go语言它本身就有自己的垃圾回收机制，
	//它不一定认为我们现在处理的这个对象值得回收，所以要多运行几次才有可能出现。
	runtime.SetFinalizer(obj1, func(obj *Object) {
		fmt.Println(obj1.id + "对象已经被回收")
	})

	//sync.WaitGroup等待两个goroutine完成
	go func() {
		obj1.AddRef()
		//通知waitGroup计数器减少
		wg.Done()
	}()

	go func() {
		obj2.AddRef()
		wg.Done()
	}()

	wg.Wait() //决定是否启动goroutine，即go的并发
	obj2.Release()
	//强制执行垃圾回收操作
	runtime.GC()
}
