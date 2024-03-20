/*
	У вас есть переменная stars, которая содержит входные пользовательские данные.
	Значение переменной stars от 1 до 5.
	1 — ★
	2 — ★★
	3 — ★★★
	4 — ★★★★
	5 — ★★★★★
	Напишите код, который проверяет значение переменной stars и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	stars := ReadInput()
	var result string

	switch stars {
	case 1:
		result = "★"
	case 2:
		result = "★★"
	case 3:
		result = "★★★"
	case 4:
		result = "★★★★"
	case 5:
		result = "★★★★★"

	}

	fmt.Println(result)
}

func ReadInput() int {
	var stars int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stars, _ = strconv.Atoi(scanner.Text())

	return stars
}
