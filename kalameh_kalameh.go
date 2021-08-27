package main

import (
	"fmt"
	"math"
)

func main() {
	var myStuff = map[int32]int{
		'a': 2,
		'e': 2,
		'i': 2,
		'o': 2,
		'u': 2,
	}

	var input string
	fmt.Scan(&input)

	counter := 0
	for _, i := range input {
		if _, ok := myStuff[i]; ok {
			counter += 1
		}
	}

	fmt.Println(int(math.Pow(2, float64(counter))))
	return
}
