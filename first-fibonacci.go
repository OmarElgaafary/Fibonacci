package main

import (
	"errors"
	"fmt"
)

func Fibonacci(x int) {
	tmp := []int{0, 1, 0}

	j := 0
	for i := 1; i < x; i++ {
		if i+1 < len(tmp) {
			tmp = append(tmp, tmp[i+1])
		} else {
			errors.New("Out of bounds")
		}
		if i+1 < len(tmp) {
			tmp[i+1] = tmp[i] + tmp[j]
			fmt.Println(tmp[i+1])

			j++
		} else {
			fmt.Println("Something went wrong here")
		}
	}
}

func main() {

	Fibonacci(15)

}
