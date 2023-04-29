package main

import "fmt"

func main() {
	nums:=[]int{0,0,0,2,0,23,12,23,0}
	moveZeros(nums)
	fmt.Println(nums)
}
func moveZeros(nums []int){
	j:=0
	for i:=0;i<len(nums);i++{
		
		if nums[i]!=0{
nums[j],nums[i]=nums[i],nums[j]
		}
		if nums[j]!=0{
			j++
		}
	}

}