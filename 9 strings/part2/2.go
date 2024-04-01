/*
	У вас есть переменные uriId, message, которые принимают входные пользовательские данные.
	Напишите код, который увеличивает значение переменной message  на значение uriId и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func updateNumberInUri(uriId int, message string) string {
	parts := strings.Split(message, "/")
	messageNumberStr := parts[len(parts)-2]
	messageNumber, _ := strconv.Atoi(messageNumberStr)
	messageNumber += uriId
	result := strings.Replace(message, messageNumberStr, strconv.Itoa(messageNumber), 1)
	return fmt.Sprintf(result)
}

func main() {
	uriId, message := ReadInput()
	result := updateNumberInUri(uriId, message)
	fmt.Println(result)
}

func ReadInput() (int, string) {
	var uriId int
	var message string
	var values []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	uriId, _ = strconv.Atoi(values[0])
	message = values[1]
	return uriId, message
}
