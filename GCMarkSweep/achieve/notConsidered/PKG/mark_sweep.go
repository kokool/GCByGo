package PKG

import "fmt"

func Mark_phase() {
	for r := range roots {
		var heap_index = roots[r]
		mark(&heap[heap_index])
	}
}

func mark(obj *Object) {
	if obj.marked == false {
		obj.marked = true
		if obj.node_type == "Key" {
			for i := range obj.children {
				//共有三种写法实现字符串转整数
				//strconv.Atoi(obj.children[i])
				//fmt.Sprintf("%d",obj.children[i])
				// index,_:=strconv.ParseInt(obj.children[i],10,64)
				index := obj.children[i]
				fmt.Println(index)
				mark(&heap[index])
			}
		}
	}
}

func Sweep_phase() {
	//默认-1表示指向null
	free_list = -1
	for id := range heap {
		if heap[id].marked == true {
			heap[id].marked = false
		} else {
			move_pointer(&heap[id])
			heap[id].next_freeIndex = free_list
			free_list = id
		}
	}
}

//当这是我们要清除标记的情况：
//	字符串就要设为空字符串切片
//	整数数组则写充满-1的整数切片，因为我们默认-1
//	当然我们其实还有很多表示方法，看自己喜欢，
func move_pointer(obj *Object) {
	// obj.children = []string{""}
	obj.children = []int{-1}
}
