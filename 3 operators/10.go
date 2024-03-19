/*
	У вас есть переменные x, y, которые содержат входные пользовательские данные.
	Напишите код, который сравнивает эти переменные и записывает результат в переменную result.
	Если x равно y тогда result = true
	Если x не равно y тогда result = false
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
	x, y := ReadInput()
	var result bool
	if x == y {
		result = true
	} else {
		result = false
	}
	fmt.Println(result)
}

func ReadInput() (int, int) {
	var x, y int
	var input string
	var values []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	values = strings.Split(input, " ")
	x, _ = strconv.Atoi(values[0])
	y, _ = strconv.Atoi(values[1])

	return x, y
}
