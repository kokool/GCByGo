package PKG

type Object struct {
	size    int
	refCnt  int
	payload []byte
}

var freeList []*Object
