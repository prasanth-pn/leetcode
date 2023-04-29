package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	a := []int{3, 30, 34, 5, 9}
	value := largestNumber(a)
	fmt.Println(value)
}
func largestNumber(a []int) string {
	res := ""
	strNums := make([]string, len(a))
	for i, v := range a {
		strNums[i] = strconv.Itoa(v)

	}
	
	cmp := func(a, b string) bool {
		return a+b > b+a
	}
	

	sort.Slice(strNums, func(i, j int) bool {
		//fmt.Println(strNums[i],strNums[j])
		return cmp(strNums[i], strNums[j])
	})
	
	res=strings.Join(strNums,"")
	res=strings.TrimRight(res,"0")
	return res   //3, 30, 34, 5, 9
}
