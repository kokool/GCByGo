package main

import (
	L "GCByGo/ReferenceCounting/achieve/DeferredReferenceCounting/PKG"
)

func main() {
	/*说明*/
	//我没测试当root->object时，出现object.refCnt=0的情况，使用的是《垃圾回收的算法与实现》图3.2例子，不同点在于，我处理掉B的引用后，新增一个跟B空间大小一致的对象D。
	//1. 当输入的值小于6，空间不够，报错
	//2. 当输入的值等于6，利用对象D，测试ZCT生效，回收对象B
	//3. 当输入的值大于6，不回收B，生成对象D
	L.Execute(6)
}
