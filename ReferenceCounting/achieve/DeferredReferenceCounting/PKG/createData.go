package PKG

import "fmt"

//书本图3.2的例子
//Num 数量,Size 总需要的空间
var Num int = 3
var Size int = 2

//BASE_SPACE：初始可分配空间大小，以一个链表节点的形式出场
func initAllocateData(BASE_SPACE int) *FreeLinkedList {
	head := &Node{data: &Object{no: "head"}}
	node0 := &Object{no: "No.0", data: make([]*Object, BASE_SPACE)}
	list := &FreeLinkedList{head: head}
	list.insertAtEnd(node0)
	list.printLinkedList()
	return list
}

func initZCT() *ZCT {
	return nil
}

func Execute(BASE_SPACE int) {
	//创建一个可供分配的空闲链表
	list := initAllocateData(BASE_SPACE)
	//接着创建一个ZCT
	zct := newZCT(6)
	Example1(Num, Size, list, zct)
}

//图3.2的例子
func Example1(Num, Size int, list *FreeLinkedList, zct *ZCT) {
	//注意：根不存放自己的地址，只存放引用的地址
	root := &Object{no: "roots", refCnt: 1, data: make([]*Object, Size)}
	for c, ch := 'A', 'A'+rune(Num); c < ch; c++ {
		//注意：我只是逻辑上删除了某个对象，实际上在这个例子中，heap仍然保持着B的地址，即实际上没被回收！所以后续需要处理掉
		heap = append(heap, newObject(string(c), Size, list, zct))
	}
	//root->A
	root.AddRef(heap[0])
	//root->C
	root.AddRef(heap[2])
	//A->B
	heap[0].AddRef(heap[1])
	fmt.Println("创建书本图3.2(a)中")
	fmt.Println(root)
	fmt.Println(heap[0])
	fmt.Println(heap[1])
	fmt.Println(heap[2])
	//让A->B => A->C
	heap[0].updatePtr(heap[1], heap[2], list, zct)
	fmt.Println("最终结果显示正确，执行图3.2(b)的结果")
	fmt.Println(root)
	fmt.Println(heap[0])
	fmt.Println(heap[1])
	fmt.Println(heap[2])
	fmt.Println("--------此时空闲链表的情况------")
	list.printLinkedList()
	fmt.Println("--------未生成D之前的ZCT的情况------")
	fmt.Println(zct)
	fmt.Println(heap)
	//在空闲链表的可分配空间不足，接着在ZCT恰好有一个可以拿来用的值的情况下，再次创建多一个对象D
	heap = append(heap, newObject("D", Size, list, zct))
	fmt.Println("--------生成D之后的ZCT的情况------")
	fmt.Println(zct)
	fmt.Println(heap)
	list.printLinkedList()
}
