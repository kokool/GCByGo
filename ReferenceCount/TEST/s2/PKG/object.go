package PKG

type Object struct {
	size int
	//计数器
	refCnt  int
	payload []byte
}

var free_list []*Object
