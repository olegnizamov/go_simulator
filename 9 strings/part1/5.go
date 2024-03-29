/*
	У вас есть переменная message, которая содержит входные пользовательские данные.
	Напишите код, который меняет первые буквы слов на заглавные и записывает результат в переменную result.
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

func convertString(message string) string {
	var resultString string
	var nextSymbolIsFirst bool = true

	for _, ch := range message {
		if nextSymbolIsFirst {
			nextSymbolIsFirst = false
			resultString += string(unicode.ToUpper(ch))
		} else if unicode.IsSpace(ch) {
			nextSymbolIsFirst = true
			resultString += string(unicode.ToLower(ch))
		} else {
			resultString += string(unicode.ToLower(ch))
		}
	}

	return resultString
}

func main() {
	message := ReadInput()
	result := convertString(message)
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
