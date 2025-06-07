package main

import (
	"fmt"
	"hedera/cp/sharedpkg"
	"slices"
	"strconv"
	"strings"
)

func main() {
	x := sharedpkg.Scanner()
	arr := []int{}
	for _, v := range x {
		v, _ := strconv.Atoi(v)
		arr = append(arr, v)
	}

	pivot := arr[len(arr)-1]
	j := 0
	for i := range arr {
		if arr[i] < pivot {
			swap(arr, i, j)
			j += 1
		}
		if i == len(arr)-1 {
			arr = slices.Insert(arr, j, arr[i])
			arr = slices.Delete(arr, len(arr)-1, len(arr))
		}
	}
	fmt.Print(strings.Trim(fmt.Sprintf("%v", arr), "[]"))
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
