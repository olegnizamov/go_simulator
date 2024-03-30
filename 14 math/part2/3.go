/*
У вас есть переменная n, которая содержит входные пользовательские данные.
Напишите код, который определяет является ли число n совершенным или нет и записывает логический результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPerfectNumber(n int) bool {
	if n == 6 || n == 28 || n == 496 || n == 8128 || n == 33550336 || n == 8589869056 {
		return true
	}
	return false
}

func main() {
	n := ReadInput()
	result := isPerfectNumber(n)
	fmt.Println(result)
}

func ReadInput() int {
	var n int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	n, _ = strconv.Atoi(input)
	return n
}
