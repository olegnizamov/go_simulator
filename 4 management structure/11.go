/*
	У вас есть переменная message, которая содержит входные пользовательские данные.
	Напишите код, который в зависимости от длины строки message записывает количество * в переменную result.
	Важно!
	Учитывайте то что ваш код должен работать как с кириллицей так и с латиницей.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	message := ReadInput()
	var result string
	length := utf8.RuneCountInString(message)
	for i := 1; i <= length; i++ {
		result = result + "*"
	}

	fmt.Println(result)
}

func ReadInput() string {
	var message string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message = scanner.Text()

	return message
}
