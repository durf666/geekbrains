package main

import (
	"fmt"
	"math/big"
)

func isEven(num int) (result bool) {
	if num%2 == 0 {
		result = true
	} else {
		result = false
	}
	return
}

func isDiv_3(num int) (result bool) {
	if num%3 == 0 {
		result = true
	} else {
		result = false
	}
	return
}

func fibi(n int) {
	var a, b = big.NewInt(0), big.NewInt(1)
	for i := 0; i < n; i++ {
		b.Add(a, b)
		a, b = b, a
		fmt.Println(a)
	}

}

func eratosfen(x int) {
	numbers := make([]bool, x)
	var result []int
	var p int = 2

	for i := 2; i < x; i++ {
		numbers[i] = true
	}
	for p < x {
		for i := p + p; i < x; i += p {
			numbers[i] = false
		}
		for i := p + 1; i < x; i++ {
			if numbers[i] {
				p = i
				break
			}
			p = x
		}
		if p == x-1 {
			break
		}
	}
	for ind, item := range numbers {
		if item {
			result = append(result, ind)
		}
	}
	fmt.Println(result)
	fmt.Println(len(result))
}

func main() {
	// fmt.Println(isEven(2))
	// fmt.Println(isDiv_3(7))
	// fibi(100)
	// eratosfen(542)
}
