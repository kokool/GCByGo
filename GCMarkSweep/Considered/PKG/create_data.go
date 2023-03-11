package PKG

import "fmt"

func mutator() {

}

//BASE_SPACE初步规定好的堆的空间
//初步规定好的对象数量
func Init_data(BASE_NUM, BASE_SPACE int) {
	for i := 0; i < BASE_NUM; i++ {
		no := NewObject("Null", nil, 0)
		heap = append(heap, *no)
	}

	//对象指向的对象（活动对象）
	heap[0] = *NewObject("Data", []int{11}, 2)
	// heap[3] = *newObject("Key", []int{5, 4}, 2)
	heap[3] = *NewObject("Key", []int{4}, 2)
	heap[4] = *NewObject("Data", []int{11}, 2)
	//对象指向的对象（非活动对象）
	heap[5] = *NewObject("Data", []int{11}, 2)
	heap[1] = *NewObject("Data", []int{11}, 2)
	heap[2] = *NewObject("Data", []int{11}, 2)
	heap[6] = *NewObject("Data", []int{11}, 2)
	//判断堆的初始化的空间够不够？
	if BASE_SPACE < init_space {
		panic("堆不够空间分配对象")
	} else if BASE_SPACE > init_space { //最终剩下的空间用来新增一个对象
		heap = append(heap, *NewObject("Data", nil, BASE_SPACE-init_space))
	}

	//roots指向的对象.
	// roots = []int{1, 3}
	roots = append(roots, 0)
	roots = append(roots, 3)
	// fmt.Println(heap)
	// fmt.Printf("类型是%T\n", heap)

}

func Print_data() {
	for i := range heap {
		fmt.Printf("--- object %d ---\n", i)
		fmt.Println(heap[i])
	}
}
