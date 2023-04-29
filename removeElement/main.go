package main

import "fmt"

func main() {
	a:=[]int{3,2,2,3}
	c:=removeElement(a,3)
	fmt.Println(c)
}
func removeElement(a []int,v int)int{
	index:=0
	//fmt.Println(len(a),j)
for i:=0;i<len(a);i++{
	if a[i]!=v{
		a[index]=a[i]
		index++
	}
}

fmt.Println(a)
return index
}