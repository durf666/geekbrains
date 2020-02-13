package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, c, square, perim float64
	fmt.Println("Введите длину первого катета:")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println("Введите длину второго катета:")
	fmt.Fscan(os.Stdin, &b)
	square = (a * b) / 2
	c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	perim = a + b + c
	fmt.Println("Длина гипотенузы: " + fmt.Sprintf("%f", c))
	fmt.Println("Площадь треугольника: " + fmt.Sprintf("%f", square))
	fmt.Println("Периметр треугольника: " + fmt.Sprintf("%f", perim))
}
