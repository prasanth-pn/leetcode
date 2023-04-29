package main

import "fmt"

func main() {
	c:=isPalindrome(1221)
	fmt.Println(c)
}
func isPalindrome(x int) bool {
    rev:=0
	k:=x
for k>0{
    rem:=k%10
    rev=rev*10+rem
	fmt.Println(rev)
    k=k/10
}
//fmt.Printf(x,rev)
return x==rev
}