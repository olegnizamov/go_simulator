/*

	У вас есть переменная n, которая содержит входные пользовательские данные.
	n всегда содержит четное число.
	Напишите код, который возвращает сумму всех четных чисел до n (включительно) и записывает результат в переменную result.
	Например:
		Eсли n = 4 тогда 2 + 4 = 6
		Eсли n = 6 тогда 2 + 4 + 6 = 12
		Eсли n = 8 тогда 2 + 4 + 6 + 8 = 20
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
	var result int
	for i := 2; i <= n; i++ {
		if i%2 != 0 {
			continue
		}

		result = result + i
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
