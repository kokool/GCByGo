package PKG

/*分配操作*/
func NewObject(node_type string, children []int, size int) *Object {
	//free_list不为空时就表示已经初始化了数据，空间分配就那么多，就只能用这些还可以用的空间
	if free_list != nil {
		chunk := pickup_chunk(node_type, children, size, free_list)
		if chunk != nil {
			return chunk
		} else {
			allocation_fail()
		}
	} else {
		object := newObject(node_type, children, size)
		return object
	}
	return nil
}

func newObject(node_type string, children []int, size int) *Object {
	if node_type == "Null" {
		if children == nil {
			children = []int{-100}
		}
	} else {
		init_space += size
		if size <= 0 {
			panic("分配不正确")
		}
		if children == nil {
			children = []int{11}
		}
	}
	return &Object{node_type: newNodeType(node_type), children: children, size: size, next_freeIndex: -100, marked: false}
}

func pickup_chunk(node_type string, children []int, needSize int, freeList []int) *Object {
	// fmt.Println(needSize, freeList)
	var heap_copy []Object
	for i := range freeList {
		index := freeList[i]
		//刚刚好就直接返回这个对象
		if heap[index].size == needSize {
			return &heap[index]
			//如果空间较大就要分块。
		} else if heap[index].size > needSize {
			heap[index].size -= needSize
			obj := *newObject(node_type, children, needSize)
			//注意：不能利用切片append插入，否则引用传递，我指的是append(heap_copy,heap[:index],obj)这样的写法
			for j := 0; j < index; j++ {
				heap_copy = append(heap_copy, heap[j])
			}
			heap_copy = append(heap_copy, obj)
			for k := index; k < len(heap); k++ {
				heap_copy = append(heap_copy, heap[k])
			}
			change_Index_list(heap_copy)
			return &heap[index]
		}
	}

	return nil
}
func allocation_fail() {
	//暂时不知道怎么搞,先return
	panic("all not-active object have not enough size")
}
