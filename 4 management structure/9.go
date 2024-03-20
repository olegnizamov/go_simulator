/*
	У вас есть переменная n, которая содержит входные пользовательские данные.
	n всегда содержит целое число.
	Напишите код, который возвращает произведение всех нечетных чисел до n (включительно) и записывает результат в переменную result.
	Например:
	Eсли n = 3 тогда 1 x 3 = 3
	Eсли n = 5 тогда 1 x 3 x 5 = 15
	Eсли n = 7 тогда 1 x 3 x 5 x 7 = 105
	и тд.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := ReadInput()
	var result int = 1
	for i := 1; i <= n; i++ {
		if i%2 != 1 {
			continue
		}

		result = result * i
	}

	fmt.Println(result)
}

func ReadInput() int {
	var n int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	return n
}
