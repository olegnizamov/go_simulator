/*
	У вас есть переменная n, которая содержит входные пользовательские данные.
	n - число.
	Напишите код, который определяет если число n положительное, тогда умножает его на 2 иначе оставляет как есть.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := ReadInput()
	var result int = n

	if n > 0 {
		result = n * 2
	}

	fmt.Println(result)
}

func ReadInput() int {
	var n int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	return n
}
