/*
	У вас есть переменные message, score, bonus которые содержат входные пользовательские данные.
	Напишите код, который выполняет конкатенацию message, score и записывает результат в переменную result.
	Значение переменой score необходимо умножить на bonus.
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
	message, score, bonus := ReadInput()
	var result = message + " " + strconv.Itoa(score*bonus)
	fmt.Println(result)
}

func ReadInput() (string, int, int) {
	var message string
	var score, bonus int
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	values = strings.Split(scanner.Text(), " | ")
	message = values[0]
	score, _ = strconv.Atoi(values[1])
	bonus, _ = strconv.Atoi(values[2])
	return message, score, bonus
}
