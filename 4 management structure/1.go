/*
	У вас есть переменная n, которая содержит входные пользовательские данные.
	n - натуральное число.
	Напишите код, который определяет, будет ли это число: нечётным, кратным 7.
	Логический результат записать в переменную result.
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
	var result bool = false
	if n%7 == 0 && n%2 != 0 {
		result = true
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
