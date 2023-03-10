package PKG

import (
	"fmt"
	"sort"
)

/*合并操作*/
func coalescing() {
	//第一步：修改各对象的size
	intervals := changeSize()
	//第二步：合并区间操作
	m := merge(intervals)
	//第三步：利用切片和循环来合并成新的堆，更新成合并的新对象，与free_list的next_freeIndex
	newHeap(m)
}

func changeSize() [][]int {
	var count int = 0
	for index := range free_list {
		if index != 0 {
			i := free_list[index]
			if free_list[index-1] == i-1 {
				count++
				heap[i-count].size += heap[i].size
				intervals = append(intervals, []int{i - count, i})
			} else {
				count = 0
			}
		}
	}
	return intervals
}

//合并区间操作
//https://leetcode.cn/problems/merge-intervals/solution/shou-hua-tu-jie-56he-bing-qu-jian-by-xiao_ben_zhu/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newHeap(m [][]int) {
	//合并
	var heap_copy []Object
	for i := 0; i < len(m)-1; i++ {
		hj := heap[:m[i][0]+1]
		hk := heap[m[i][1]+1 : m[i+1][0]+1]
		for j := range hj {
			heap_copy = append(heap_copy, hj[j])
		}
		for k := range hk {
			heap_copy = append(heap_copy, hk[k])
		}
	}
	change_Index_list(heap_copy)
}

//更改next_freeIndex和更改free_list
func change_Index_list(heap_copy []Object) {
	var list []int
	var next_index int = -1
	for i := range heap_copy {
		if heap_copy[i].children[0] == -1 {
			list = append(list, i)
			heap_copy[i].next_freeIndex = next_index
			next_index = i
		}
	}
	heap = heap_copy
	fmt.Println(heap)
	free_list = list
	fmt.Println(free_list)
}
