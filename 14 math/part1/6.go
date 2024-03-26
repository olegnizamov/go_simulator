/*
	У вас есть переменные a, b, c, которые содержат входные пользовательские данные.
	a, b, c – стороны треугольника.
	Напишите код, который находит периметр треугольника и записывает результат в переменную result.
	Периметр треугольника равняется сумме всех его сторон
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(a, b, c int) int {
	return a + b + c
}

func main() {
	a, b, c := ReadInput()
	result := calculate(a, b, c)
	fmt.Println(result)
}

func ReadInput() (int, int, int) {
	var a, b, c int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	a, _ = strconv.Atoi(values[0])
	b, _ = strconv.Atoi(values[1])
	c, _ = strconv.Atoi(values[2])
	return a, b, c
}
