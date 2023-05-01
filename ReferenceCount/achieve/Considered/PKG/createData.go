package PKG

import "fmt"

//书本图3.2的例子
//Num 数量,Size 总需要的空间
var Num int = 3
var Size int = 6
var heap []*Object

func Example1(Num, Size int, list *FreeLinkedList) {
	avgSize := Size / Num
	root := &Object{no: "root", refCnt: 1, data: make([]byte, 0)}
	for c, ch := 'A', 'A'+rune(Num); c < ch; c++ {
		heap = append(heap, newObject(string(c), avgSize, list))
	}
	//root->A
	root.AddRef(heap[0])
	fmt.Println(heap[0])
	fmt.Println(heap[1])
	fmt.Println(heap[2])
	//root->C
	root.AddRef(heap[2])
	//A->B
	heap[0].AddRef(heap[1])
	fmt.Println("创建书本图3.2(a)中update_prt()函数执行时的情况")
	fmt.Println(root)
	fmt.Println(heap[0])
	fmt.Println(heap[1])
	fmt.Println(heap[2])
	//让A->B => A->C
	heap[0].updatePtr(heap[1], heap[2], list)
	fmt.Println("最终结果显示正确，执行图3.2(b)的结果")
	fmt.Println(root)
	fmt.Printf("%p", heap[0])
	fmt.Println(heap[0])
	fmt.Printf("%p", heap[1])
	fmt.Println(heap[1])
	fmt.Printf("%p", heap[2])
	fmt.Println(heap[2])
}

//BASE_SPACE：初始可分配空间大小，以一个链表节点的形式出场
func InitData(BASE_SPACE int) *FreeLinkedList {
	head := &Node{data: &Object{no: "head"}}
	node0 := &Object{no: "No.0", data: make([]byte, 10)}
	list := &FreeLinkedList{head: head}
	list.insertAtEnd(node0)
	list.printLinkedList()
	return list
}

func Execute(BASE_SPACE int) {
	list := InitData(BASE_SPACE)
	Example1(Num, Size, list)
	mergeChunk(list)
}
