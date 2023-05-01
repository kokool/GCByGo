# GCMarkSweep（标记-清除）
## GCMarkSweep\achieve
只考虑书本的方案去实现书本提供的例子，而不考虑优化等问题，只要能运行就算赢。

- 存储结构采用的是切片、数组，而不是链表
- 在处理分块合并的这块内容，处理的过于麻烦

### GCMarkSweep\achieve\notConsidered
不考虑分配与合并的情况，实现中村成洋的《垃圾回收的算法与实现》的例子。
https://blog.csdn.net/kokool/article/details/129384458

### GCMarkSweep\achieve\Considered
在`notConsidered`代码的基础上，改成考虑分配与合并的情况，实现中村成洋的《垃圾回收的算法与实现》的书本例子
https://blog.csdn.net/kokool/article/details/129451004

## GCMarkSweep\Simple
忽略堆、链表、数组等数据结构，只考虑对象的处理，实现的简单示例，只剖析最核心的原理，由chatGpt提供。
https://blog.csdn.net/kokool/article/details/129384458

# ReferenceCount（引用计数法）
## ReferenceCount\achieve
只考虑书本的方案去正确实现书本提供的例子，而不考虑优化等问题。
### ReferenceCount\achieve\Considered
把图3.2该例子的所有涉及到的结构体实现都考虑了一遍

## ReferenceCount\TEST
对具体代码进行测试的内容，目前已经测试完成
### ReferenceCount\TEST\s1（chatGpt提供的一个简易实现）
### ReferenceCount\TEST\s2（chatGpt参考伪代码的初期版本）
### ReferenceCount\TEST\s3（可运行，但是不用在意）

# 总结
写的不完美，也不实用（性能堪忧、多线程环境不能运行），但是字符输出的结果较直观，例子选用的是《垃圾回收的算法与实现》，不过例子少，难免出现错漏。

## chatGpt使用经验
鉴于本人就是个小菜鸡，本想靠着AI起飞，结果发现还是靠自己｡

･ﾟﾟ*(>д<)*ﾟﾟ･｡

### 需要避免的地方：
- 经过实验，chatGpt3.5 不可以提供太复杂的示例，必须足够简单。
- 利用网上所谓的提示词和模板去限定它进行某个任务下的代码debug或者测试并不可靠。
- 在中文语境下（没测试英文语境），chatGpt3.5可能并不能理解伪代码是否存在隐藏的含义，就如同一句话在不同情景和不同的时代，有不同意思。而我在中村成洋的《垃圾回收的算法与实现》中，会经常遇到一个变量有多种含义的问题，如果不给我相关的文段内容，我都不理解什么意思，何况chatGpt呢，那么自然只能往错误的方向走到黑了。

### 可以积极使用的地方
- 它对于自己出现的代码错误提示，是真得可以给出修改方向，而且绝大多数都是正确的，不过前提是搞清自己的问题够不够清晰，有没有错误的地方！
- 程序员的百科全书，如果我们忘记了某个方法的使用或者自己希望能够实现一个几段代码可以做到的需求，它也能很好得处理，这适用于把一个大的问题拆分成多个小问题，这时它就能起到作用了。

总而言之，chatGpt3.5 在代码编程这块只能够做到**短小**，还有**快**就是了。