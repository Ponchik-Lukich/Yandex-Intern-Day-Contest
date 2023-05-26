package main

import (
	"bufio"
	"fmt"
	"os"
)

func combinations(arr []int, size int, index int, data []int, i int) {
	if index == size {
		for j := 0; j < size; j++ {
			fmt.Print(data[j], " ")
		}
		fmt.Println()
		return
	}
	if i >= len(arr) {
		return
	}
	data[index] = arr[i]
	combinations(arr, size, index+1, data, i+1)
	combinations(arr, size, index, data, i+1)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	arr := make([]int, n)
	hash := make(map[int]int)
	elements := make(map[int][]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		for j := 0; j < a; j++ {
			var b int
			fmt.Fscan(in, &b)
			elements[i] = append(elements[i], b)
			hash[b] += arr[i]
		}
	}
	resArr := make([]int, len(hash))
	c := 0
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		if _, ok := hash[i+1]; ok {
			resArr[c] = hash[i+1]
		}
		c++
	}

	// Determine the maximum combination size based on the length of resArr and k
	maxSize := k
	if len(resArr) < k {
		maxSize = len(resArr)
	}

	// Generate all combinations of resArr of size maxSize down to 1
	for size := maxSize; size > 0; size-- {
		data := make([]int, size)
		combinations(resArr, size, 0, data, 0)
	}
}
