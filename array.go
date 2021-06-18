package main

import "fmt"

func main() {

	array := []int{-3, 1, 2, 3, -1, -4, -2}

	fmt.Println(Solve(array))

}

func Solve(arr []int) int {
	visit := map[int]bool{}
	for i := 0; i < len(arr); i++ {
		visit[arr[i]] = true
	}
	for i := 0; i < len(arr); i++ {
		if !visit[-arr[i]] {
			return arr[i]
		}
	}
	return 0
}
