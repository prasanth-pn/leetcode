package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 2)
fmt.Println(a)
}
func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)

}
func reverse(a []int, start, end int) {
	for start < end {
		a[start], a[end] = a[end], a[start]
		start++
		end--
	}
}
