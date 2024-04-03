/*
	У вас есть переменная message, которая принимает входные пользовательские данные.
	Данные такого формата: букваЧИСЛО,букваЧИСЛО,букваЧИСЛО и тд...
	Напишите код, который складывает все числа и записывает результат в переменную result.
	Важно!
	Учитывайте то что ваш код должен работать как с кириллицей так и с латиницей.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func sumNumbers(message string) int {

	nums := strings.Split(message, ",")
	result := 0

	for _, numStr := range nums {
		var newDigit string

		for _, ch := range numStr {
			if unicode.IsDigit(ch) {
				newDigit += ch
			}
		}

		num, _ := strconv.Atoi(numStr)
		result += num
	}

	result := 0

	for _, ch := range message {
		if unicode.IsDigit(ch) {
			num, _ := strconv.Atoi(ch)
			result += num
		}
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
