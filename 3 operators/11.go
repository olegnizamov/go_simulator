/*
	У вас есть переменные role, которая содержит входные пользовательские данные.
	Напишите код, который в зависимости от значения переменной role записывает соответствующий результат в переменную result.
	1 — admin
	2 — moderator
	3 — user
	default — guest
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	role := ReadInput()
	var result string
	switch role {
	case 1:
		result = "admin"
	case 2:
		result = "moderator"
	case 3:
		result = "user"
	default:
		result = "guest"
	}

	fmt.Println(result)
}

func ReadInput() int {
	var role int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	role, _ = strconv.Atoi(scanner.Text())

	return role
}
