package main

import "fmt"

func AppendInt(a, b int) (nums []int) {
	nums = append(nums, a)
	nums = append(nums, b)
	return
}

func main() {

	l := AppendInt(1, 3)
	fmt.Println(l)
}
