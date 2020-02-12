package main

import (
	"fmt"
	"os"
)

func main() {
	const usd = 0.016
	var rub, sum float64
	fmt.Println("Введите желаемую сумму в рублях:")
	fmt.Fscan(os.Stdin, &rub)
	sum = usd * rub
	fmt.Println("Введенная вами сумма в пересчете на доллары со ставит: " + fmt.Sprintf("%f", sum))
}
