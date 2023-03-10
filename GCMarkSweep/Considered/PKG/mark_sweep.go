package PKG

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
				index := obj.children[i]
				mark(&heap[index])
			}
		}
	}
}

func Sweep_phase() {
	//默认-1表示一开始没有分块
	chunk_index = -1

	for id := range heap {
		if heap[id].marked == true {
			heap[id].marked = false
		} else {
			// fmt.Println(free_list)
			// fmt.Println("a", chunk_index, heap[id].next_freeIndex, id)
			become_chunk(&heap[id])
			heap[id].next_freeIndex = chunk_index
			chunk_index = id
			// fmt.Println("a", chunk_index, heap[id].next_freeIndex, id)
			free_list = append(free_list, chunk_index)
			// fmt.Println(free_list)
		}
	}
	coalescing()
}

func become_chunk(obj *Object) {
	obj.children = []int{-1}
}
