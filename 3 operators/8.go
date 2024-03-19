/*
У вас есть переменные x, y и op, которые содержат входные пользовательские данные.
Напишите код, который выполняет операции сложения/вычитания над переменными x, y  и записывает результат в переменную result.
Если op == "+" тогда result = x + y
Если op == "-" тогда result = x - y
Если op == "" тогда result = 0
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
	x, y, op := ReadInput()
	var result int

	switch op {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "":
		result = 0
	}

	fmt.Println(result)
}

func ReadInput() (int, int, string) {
	var x, y int
	var op, input string
	var values []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	values = strings.Split(input, " ")
	x, _ = strconv.Atoi(values[0])
	y, _ = strconv.Atoi(values[2])
	op = values[1]

	return x, y, op
}
