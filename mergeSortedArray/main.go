package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 6, 7, 12}
	brr := []int{1, 4, 8}
	sum := mergeTwoSortedArray(arr, brr)
	fmt.Println(sum)
}

// merge two sorted array using concadination
//select sort method is used to do this problem
func mergeTwoSortedArray(a, b []int) []int {
	result:=[]int{}
	i,j:=0,0

	for i<len(a)&&j<len(b){
		if a[i]<b[j]{
			result=append(result, a[i])
			i++
		}else{
			result=append(result, b[j])
			j++
		}
	}
	for i<len(a){
		result=append(result, a[i])
		i++
	}
	for j<len(b){
		result=append(result, b[j])
		j++
	}
return result
}
