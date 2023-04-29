package main

import "fmt"

func main() {
	two:=[][]int{
		{1,1,1},
		{0,1,1},
		{0,0,0},
	}
	for _,v:=range two{
		arr:=[]int{}
		fmt.Print(v)
		for i:=len(v)-1;i>=0;i--{
			if v[i]==0{
				v[i]=1
			}else{
				v[i]=0
			}
arr = append(arr, v[i])


		}
		v=arr
		fmt.Print(v)
		fmt.Println()
	}
	fmt.Println(two)
	//fmt.Println(two)
	
}
