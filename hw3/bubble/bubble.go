package main

import "fmt"

func main() {
	res := bubbleSort([]int{5, 3, 6, 8, 1, 2})
	fmt.Println(res)
}

func bubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			fmt.Println("before: arr[i+1] =", arr[i+1], "arr[i] =", arr[i])
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
			fmt.Println("after: arr[i+1] =", arr[i+1], "arr[i] =", arr[i], "arr = ", arr)
		}
	}
	return arr
}
