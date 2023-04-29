package main

import "fmt"

func main() {
	a:=[]int{3,3,4,4,5,2,3,2,1,1}
	fmt.Println(singleNumber(a))
}
func singleNumber(a []int)int{
	
	for i:=0;i<len(a);i++{
	a[0]^=a[i]
	fmt.Println(a[0])
	//n=i
	}
	return a[0]
}