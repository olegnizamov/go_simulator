/*
	У вас есть переменная n, которая содержит входные пользовательские данные.
	Напишите код, который переводит десятичное число n в двоичное и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func toBinaryNumber(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func main() {
	n := ReadInput()
	result := toBinaryNumber(n)
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
