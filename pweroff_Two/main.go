package main

import (
	"fmt"
)

// given an interger n  return true if it is a power of two other wise ,return false
func main() {
	fmt.Println(poweroffTwo(17))
}
func poweroffTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n == 1 || (n%2 == 0 && poweroffTwo(n/2))

}
