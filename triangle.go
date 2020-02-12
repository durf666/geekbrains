package main

import (
	"fmt"
	"os"
)

func main() {
	var sum, percent float64
	var i int = 1
	fmt.Println("Введите сумму вклада:")
	fmt.Fscan(os.Stdin, &sum)
	fmt.Println("Введите годовой процент:")
	fmt.Fscan(os.Stdin, &percent)
	percent = percent / 100
	for i <= 5 {
		sum = sum * percent
		fmt.Println(sum)
	}

}
