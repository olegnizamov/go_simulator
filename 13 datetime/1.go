/*
	У вас есть переменная year, которая содержит входные пользовательские данные.
	Напишите код, который определяет является ли значение переменной year високосным годом или нет и записывает логический результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func main() {
	year := ReadInput()
	result := isLeapYear(year)
	fmt.Println(result)
}

func ReadInput() int {
	var year int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	year, _ = strconv.Atoi(input)
	return year
}
