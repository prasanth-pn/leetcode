package main

import "fmt"

func main() {
	prices := []int{1,2} //[7,6,4,3,1],7,1,5,3,6,4
	fmt.Println(maxProfit(prices))

}
func maxProfit(prices []int) int {
	var line int
	count := 0
    count1:=0
	fmt.Println(len(prices))
    if len(prices)==1{
        return 0
    }
	min := prices[1]
	for i, v := range prices {

		if v < min && count < 1 {
			min = v
			line = i
			count++
			break
		} 
		if v>min&&count1<1{
            min=v
            count1++
			break
        }
        

	}
	fmt.Println(count)
	max := min
	prices = prices[line:]
	if len(prices) == 1 {
		return 0
	}
	fmt.Println(prices)
	for _, v := range prices {
		if v > max {
			max = v

		}
	}
	return max - min
}

