package main

import (
	"fmt"
)

func main() {

	var arr = []int{1, 2, 3, 4, 5, 6}
	//买卖股票的最佳时机 1
	// fmt.Print(MaxProfitOne(arr))

	//买卖股票的最佳时机 2
	fmt.Print(MaxProfitTwo(arr))

}
