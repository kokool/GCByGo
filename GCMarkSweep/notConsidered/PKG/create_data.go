package PKG

import "fmt"

func Init_data() {
	//初始化堆中的所有对象，对于多出来的对象则进行默认操作表示成非活动对象
	h := Object{marked: false, node_type: "Null", children: []int{-1}, size: 0, next_freeIndex: -100}
	for i := range heap {
		heap[i] = h
	}

	var key_type = newNodeType("Key")
	var data_type = newNodeType("Data")

	//对象指向的对象（活动对象）
	heap[1] = Object{children: []int{11}, node_type: data_type, marked: false, size: 2, next_freeIndex: -100}
	heap[3] = Object{children: []int{5, 4}, node_type: key_type, marked: false, size: 2, next_freeIndex: -100}
	heap[4] = Object{children: []int{44}, node_type: data_type, marked: false, size: 2, next_freeIndex: -100}
	heap[5] = Object{children: []int{55}, node_type: data_type, marked: false, size: 2, next_freeIndex: -100}
	//对象指向的对象（非活动对象）
	heap[0] = Object{children: []int{20}, node_type: data_type, marked: false, size: 2, next_freeIndex: -100}
	heap[2] = Object{children: []int{1}, node_type: key_type, marked: false, size: 2, next_freeIndex: -100}
	heap[6] = Object{children: []int{66}, node_type: data_type, marked: false, size: 2, next_freeIndex: -100}
	//roots指向的对象
	roots = []int{1, 3}
}

func Print_data() {
	for i := range heap {
		fmt.Printf("--- object %d ---\n", i)
		fmt.Println(heap[i])
	}
}
