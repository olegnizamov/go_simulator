/*
	У вас есть переменные message, score которые содержат входные пользовательские данные.
	Напишите код, который выполняет конкатенацию message, score и записывает результат в переменную result.
	Значение переменой score необходимо умножить на 2.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	message, score := ReadInput()
	var result = message + " " + strconv.Itoa(score*2)
	fmt.Println(result)
}

func ReadInput() (string, int) {
	var message string
	var score int
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	values = strings.Split(scanner.Text(), " | ")
	message = values[0]
	score, _ = strconv.Atoi(values[1])
	return message, score
}
