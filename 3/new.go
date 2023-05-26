package main

import (
	"bufio"
	"fmt"
	"os"
)

var maxProfit int
var minLen int
var arr []int
var disCount map[int]int
var elements map[int][]int
var bestCombination []int

func combinations(eArr []int, arr []int, size int, index int, data []int, i int) {
	if index == size {
		copyEArr := make([]int, len(eArr))
		copy(copyEArr, eArr)
		profit := 0
		for _, number := range data {
			for i, e := range elements {
				for _, n := range e {
					if n == number {
						if copyEArr[i] != 0 {
							profit += copyEArr[i] * disCount[number] / 100
							copyEArr[i] = copyEArr[i] * (100 - disCount[number]) / 100
						}
					}
				}
			}
		}
		if profit > maxProfit || (profit == maxProfit && len(data) < minLen) {
			maxProfit = profit
			minLen = len(data)
			bestCombination = make([]int, len(data))
			copy(bestCombination, data)
		}
		return
	}
	if i >= len(arr) {
		return
	}
	data[index] = arr[i]
	combinations(eArr, arr, size, index+1, data, i+1)
	combinations(eArr, arr, size, index, data, i+1)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	arr = make([]int, n)
	disCount = make(map[int]int)
	hash := make(map[int]int)
	elements = make(map[int][]int)
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
			resArr[c] = i + 1
			disCount[i+1] = a
		}
		c++
	}
	maxSize := k
	if len(resArr) < k {
		maxSize = len(resArr)
	}

	for size := maxSize; size > 0; size-- {
		data := make([]int, size)
		combinations(arr, resArr, size, 0, data, 0)
	}

	fmt.Println(minLen)
	for _, number := range bestCombination {
		fmt.Print(number, " ")
	}
	fmt.Println()
}
