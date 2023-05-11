package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Node struct {
	value  interface{}
	number int
}

type IntHeap []Node

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].value.(int) > h[j].value.(int) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	arr := make([]int, n)
	hash := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		for j := 0; j < a; j++ {
			var b int
			fmt.Fscan(in, &b)
			hash[b] += arr[i]
		}
	}
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		hash[i+1] = hash[i+1] * a / 100
		if hash[i+1] == 0 {
			continue
		}
		heap.Push(h, Node{hash[i+1], i + 1})
	}
	var size int
	if h.Len() < k {
		size = h.Len()
	} else {
		size = k
	}
	fmt.Fprintln(out, size)
	for i := 0; i < size; i++ {
		node := heap.Pop(h).(Node)
		fmt.Fprint(out, node.number, " ")
	}
}
