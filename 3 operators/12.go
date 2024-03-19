/*
	У вас есть переменная age, которая принимает входные пользовательские данные.
	Напишите код, который проверяет переменную age и записывает результат в переменную result.
	Если age больше или равно 18 тогда записываем взрослый иначе записываем подросток.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	age := ReadInput()
	var result string
	if age >= 18 {
		result = "взрослый"
	} else {
		result = "подросток"
	}
	fmt.Println(result)
}

func ReadInput() int {
	var age int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text())

	return age
}
