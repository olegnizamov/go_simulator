/*
	У вас есть переменная message, которая принимает входные пользовательские данные.
	Данные такого формата: ЧИСЛО,ЧИСЛО,ЧИСЛО и тд...
	Напишите код, который складывает все числа и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumNumbers(message string) int {
	nums := strings.Split(message, ",")
	result := 0

	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		result += num
	}

	return result
}

func main() {
	message := ReadInput()
	result := sumNumbers(message)
	fmt.Println(result)
}

func ReadInput() string {
	var message string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	message = input
	return message
}
