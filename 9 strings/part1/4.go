/*
	У вас есть переменная message, которая содержит входные пользовательские данные.
	Напишите код, который меняет строчные буквы на заглавные, а заглавные на строчные и записывает результат в переменную result.
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

func invertCase(message string) string {

	var resultString string
	for _, ch := range message {
		if unicode.IsUpper(ch) {
			resultString += string(unicode.ToLower(ch))
		} else if unicode.IsLower(ch) {
			resultString += string(unicode.ToUpper(ch))
		}
	}
	return resultString
}

func main() {
	message := ReadInput()
	result := invertCase(message)
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
