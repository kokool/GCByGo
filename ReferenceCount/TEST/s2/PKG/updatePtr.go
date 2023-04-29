package PKG

// updatePtr updates the pointer `ptr` to point to `obj`,
// incrementing the reference count of `obj` and decrementing
// the reference count of the previous Object pointed to by `ptr`.
func updatePtr(ptr *Object, obj *Object) {
	incRefCnt(obj)
	if ptr != nil {
		decRefCnt(ptr)
	}
	*ptr = *obj
}

// incRefCnt increments the reference count of an Object
func incRefCnt(obj *Object) {
	obj.refCnt++
}

// decRefCnt decrements the reference count of an Object
// and frees it if the reference count drops to 0.
func decRefCnt(obj *Object) {
	obj.refCnt--
	if obj.refCnt == 0 {
		free_list = append(free_list, obj)
	}
}
