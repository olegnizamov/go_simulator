/*
	У вас есть переменная message, которая содержит входные пользовательские данные.
	Напишите код, который подсчитывает все заглавные буквы в переменной message и записывает результат в переменную result.
	Важно!
	Учитывайте то что ваш код должен работать как с кириллицей так и с латиницей.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func countUpperCase(message string) int {

	var count int

	for _, ch := range message {
		if unicode.IsUpper(ch) {
			count++
		}
	}
	return count
}

func main() {
	message := ReadInput()
	result := countUpperCase(message)
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
