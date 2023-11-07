package main

import (
	"fmt"
	"hogwarts-test/calculate"
)

func main() {
	coins := []int{1, 5, 10, 25, 50, 100, 200}
	target := 350 // G3.5

	ways := calculate.CountWaysToMakeChange(coins, target)
	fmt.Printf("Number of ways to make G3.5: %d\n", ways)
}
