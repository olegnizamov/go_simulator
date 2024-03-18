/*
	У вас есть переменные message1, message2 которые содержат входные пользовательские данные.
	Напишите код, который выполняет конкатенацию message1, message2 и записывает результат в переменную result.
	Важно! Между сообщениями message1 и message2 должен быть один пробел.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	message1, message2 := ReadInput()
	result := message1 + " " + message2
	fmt.Println(result)
}

func ReadInput() (string, string) {
	var message1, message2 string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	values = strings.Split(scanner.Text(), " | ")
	message1 = values[0]
	message2 = values[1]
	return message1, message2
}
