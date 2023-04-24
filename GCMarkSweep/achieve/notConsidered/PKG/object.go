package PKG

type Object struct {
	//利用go的空接口类型表示指针数组，当然你也可以用[]int或者[]string，用map[]其实更合适
	//总之我们要实现的要求如下：
	//key为本对象的index
	//value为其他对象的index，但是如果不指向的话该怎么处理？
	children       []int
	node_type      NodeType // flag
	marked         bool     //是否被标记
	size           int      //域的大小，假设一个指针占1字节的域
	next_freeIndex int      //free_list的指向
}

var roots []int

var free_list int

//设个常量，自己看着办
const (
	HEAP_SIZE = 7 //就拿书本的例子
)

var heap [HEAP_SIZE]Object //书本1.4，堆存放的就是对象

type NodeType string

//专门识别到底是放pointer还是data
func newNodeType(node_type string) NodeType {
	if node_type == "Key" || node_type == "Data" {
		return NodeType(node_type)
	} else {
		panic("error type")
	}
}
