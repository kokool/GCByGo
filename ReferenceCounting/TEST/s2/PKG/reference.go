package PKG

/*分配操作*/
// newObj creates a new Object with given size
func newObj(size int) *Object {
	//free_list不为空时就表示已经初始化了数据，空间分配就那么多，就只能用这些还可以用的空间
	if free_list != nil {
		obj := pickupChunk(size)
		if obj == nil {
			allocation_fail()
		}
		obj.refCnt = 1
		return obj
		//空闲链表表示空的话，没必要继续
	} else {
		allocation_fail()
	}
	return nil
}

func allocation_fail() {
	panic("allocation failed")
}

// pickupChunk searches the free_list for a chunk of memory
// that is at least the size requested.
// If it finds a chunk that is the exact size requested,
// it returns that chunk. If it finds a larger chunk,
// it splits the chunk into two parts (one of size `size`
// and one of the remaining size) and returns the
// chunk of size `size`. It also adds the remaining chunk
// to the free_list.
func pickupChunk(size int) *Object {
	for i, obj := range free_list {
		if obj.size >= size {
			// exact size match
			if obj.size == size {
				free_list = append(free_list[:i], free_list[i+1:]...)
				return obj
			}
			// split the chunk and return the requested size
			remainingSize := obj.size - size
			obj.size = size
			obj.refCnt = 0 // new chunk has 0 references
			obj.payload = obj.payload[:size]
			remainingChunk := &Object{
				size:    remainingSize,
				refCnt:  0, // new chunk has 0 references
				payload: obj.payload[size:],
			}
			free_list[i] = remainingChunk
			return obj
		}
	}
	return nil // no chunk of adequate size found
}
