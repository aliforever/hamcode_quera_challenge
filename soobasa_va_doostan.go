package main

import (
	"fmt"
)

func main() {

	var n, a, b int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	var goals = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&goals[i])
	}

	firstHalf := []int{0, 45 + a}
	secondHalf := []int{45, 90 + b}
	var nextGoalTime int

	for index, goalTime := range goals {
		if goalTime > secondHalf[1] {
			fmt.Println("NO")
			return
		}

		if index != len(goals)-1 {
			nextGoalTime = goals[index+1]
		}

		if goalTime > firstHalf[1] {
			// Happened In Second Half
			if nextGoalTime < goalTime {
				fmt.Println("NO")
				return
			}
		}

	}
	fmt.Println("YES")
	return
}
