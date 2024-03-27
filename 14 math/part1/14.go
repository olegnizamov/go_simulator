/*
	У вас есть переменная n, которая содержит входные пользовательские данные.
	Напишите код, который определяет целое число, соответствующее количеству цифр в заданном целом числе n и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countDigits(n int) int {
	return len(strconv.Itoa(n))
}

func main() {
	n := ReadInput()
	result := countDigits(n)
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
