# GCMarkSweep（标记-清除）
标记清楚算法的内容，目录与接下来的计划如下，尽请期待。

- 经过实验，chatGpt3.5 不太可以提供太复杂的示例，代码bug等问题，还是得靠自己，但是它至少提供了一个简易实现思路给我。

## GCMarkSweep\achieve
只考虑书本的方案去实现书本提供的例子，而不考虑优化等问题，只要能运行就算赢。

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

## ReferenceCount\TEST
对具体代码进行测试的内容，可能可以运行但是结果不对，也可能不能运行

# 总结
写的不完美，也不实用，但是较直观，例子选用的是《垃圾回溯的算法与实现》，例子少，难免出现错漏。