/*
	У вас есть переменные x, y и direction которые содержат входные пользовательские данные.
	x, y содержат числа - стартовая позиция игрока.
	direction содержит направление движения, одного из: up, down, left, right.
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
	x, y, direction := ReadInput()
	var result string

	switch direction {
	case "down":
		y = y + 1
		result = "x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y) + ", direction: " + direction
	case "left":
		x = x - 1
		result = "x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y) + ", direction: " + direction
	case "right":
		x = x + 1
		result = "x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y) + ", direction: " + direction
	case "up":
		y = y - 1
		result = "x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y) + ", direction: " + direction

	}

	fmt.Println(result)
}

func ReadInput() (int, int, string) {
	var x, y int
	var direction, input string
	var values []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	values = strings.Split(input, " ")
	x, _ = strconv.Atoi(values[0])
	y, _ = strconv.Atoi(values[1])
	direction = values[2]

	return x, y, direction
}
