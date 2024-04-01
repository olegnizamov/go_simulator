/*
	У вас есть переменная message, которая принимает входные пользовательские данные.
	Напишите код, который проверяет является ли строка палиндромом и записывает логический результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func isPalindrome(message string) bool {

	var values []string

	for len(message) > 0 {
		r, size := utf8.DecodeRuneInString(message)
		values = append(values, string(r))
		message = message[size:]
	}

	isPalindrome := true
	for i := 0; i < len(values)/2; i++ {
		if values[i] != values[len(values)-1-i] {
			isPalindrome = false
			break
		}
	}
	return isPalindrome
}

func main() {
	message := ReadInput()
	result := isPalindrome(message)
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
