package PKG

type Object struct {
	size    int
	refCnt  int
	payload []byte
}

var free_list []*Object
