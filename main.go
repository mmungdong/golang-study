package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	var s sort.IntSlice = make([]int, 4)
	s = append(s, 5)
	s = append(s, 51)
	s = append(s, 99)
	s = append(s, 1)
	s.Sort()
	fmt.Println(s)
}

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return h.IntSlice[i] > h.IntSlice[j]
}

func (hp) Push(any)     {}
func (hp) Pop() (_ any) { return }

func maxKelements(nums []int, k int) (ans int64) {
	h := hp{nums}
	heap.Init(&h)
	for ; k > 0; k-- {
		ans += int64(h.IntSlice[0])
		h.IntSlice[0] = (h.IntSlice[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return ans
}
