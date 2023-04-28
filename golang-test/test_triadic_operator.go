package main

import "fmt"

// Any 若expr成立，则返回a；否则返回b。
func Any[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}

func main() {
	bDebug := true
	log_pre := Any(bDebug, "debug_", "release_")
	fmt.Printf("log_pre :%v", log_pre)
}
