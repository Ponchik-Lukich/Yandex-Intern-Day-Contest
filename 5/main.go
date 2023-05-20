package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Node struct {
	number  int
	freq    int
	desired int
	prev    int
}

type IntHeap []Node

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
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
	var n int
	fmt.Fscan(in, &n)
	hash := make(map[int]int)
	elements := make(map[int][]int)
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		hash[b] += 1
		elements[b] = append(elements[b], a)
	}
	for k, v := range hash {
		if n%v != 0 {
			heap.Push(h, Node{freq: v, number: k, desired: n/v + 1, prev: -1})
		} else {
			heap.Push(h, Node{freq: v, number: k, desired: n / v, prev: -1})
		}
		//println("Add: ", k, v)
	}
	queue := make([]Node, 0)
	flag := false
	i := 0
	done := len(hash)
	max := -1
	for i < n {
		node := heap.Pop(h).(Node)
		println(node.number, node.freq, node.desired, node.prev, "i:", i)
		if node.prev != -1 {
			if i-node.prev == node.desired {
				flag = true
				product := elements[node.number][0]
				fmt.Fprint(out, product)
				if len(elements[node.number]) >= 1 {
					fmt.Fprint(out, " ")
				}
				elements[node.number] = elements[node.number][1:]
				if node.freq > 1 {
					node.freq -= 1
					node.prev = i
					heap.Push(h, node)
				} else {
					done -= 1
				}
				i++
			} else {
				queue = append(queue, node)
				if node.desired > max {
					max = node.desired
				}
				if done == len(queue) {
					for _, node := range queue {
						if node.desired == max {
							node.desired -= 1
							max = -1
						}
						heap.Push(h, node)
					}
				}
			}
		} else {
			flag = true
			product := elements[node.number][0]
			fmt.Fprint(out, product)
			if len(elements[node.number]) >= 1 {
				fmt.Fprint(out, " ")
			}
			elements[node.number] = elements[node.number][1:]
			if node.freq > 1 {
				node.freq -= 1
				node.prev = i
				heap.Push(h, node)
			}
			i++
		}
		if flag {
			for _, node := range queue {
				heap.Push(h, node)
			}
			queue = make([]Node, 0)
			flag = false
		}
	}
}
