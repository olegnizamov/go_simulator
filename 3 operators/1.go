/*
	У вас есть переменные x, y, которые содержат входные пользовательские данные.
	Напишите код, который складывает эти переменные и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(x, y int) int {
	return x + y
}

func main() {
	x, y := ReadInput()
	result := sum(x, y)
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
