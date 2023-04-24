package PKG

/*对象/非活动对象*/
type Object struct {
	children       []int
	node_type      NodeType
	marked         bool
	size           int
	next_freeIndex int
}

type NodeType string

func newNodeType(node_type string) NodeType {
	if node_type == "Key" || node_type == "Data" || node_type == "Null" {
		return NodeType(node_type)
	} else {
		panic("error type")
	}
}

/*根*/
var roots []int

/*链表。*/

var chunk_index int

var free_list []int

/*堆*/

var heap []Object

var init_space int = 0

var intervals [][]int
