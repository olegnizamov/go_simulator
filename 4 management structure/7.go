/*
	У вас есть переменная n, которая содержит входные пользовательские данные.
	Напишите код, который проверяет является ли значение n негативным, позитивным или равно 0 и записывает результат в переменную result.
	Если n позитивное число тогда result = "Число позитивное"
	Если n негативное число тогда result = "Число негативное"
	Если n равно 0 число тогда result = "Число равно 0"
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
	var result string

	if n > 0 {
		result = "Число позитивное"
	} else if n < 0 {
		result = "Число негативное"
	} else {
		result = "Число равно 0"
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
