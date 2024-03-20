/*
	У вас есть переменные x, y и direction которые содержат входные пользовательские данные.
	Игровое поле размера от 0 до 100 по x и от 0 до 100 по y.
	x, y содержат числа - стартовая позиция игрока.
	direction содержит направление движения, одного из: up, down, left, right.
	Напишите код, который высчитывает новую позицию игрока после перемещения в этом направлении на 1 и записывает результат в переменную result.
	Учитывайте то, что новая позиция игрока не должна быть меньше 0 или больше 100 как x так и по y.
	Например, если новая позиция игрока больше 100 тогда, устанавливаем ему значение 100, а если новая позиция игрока меньше 0 тогда устанавливаем ему значение 0 по  x или по y в зависимости от того какую границу игрок перешел.
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
		if y > 100 {
			y = 100
		}
		result = "x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y) + ", direction: " + direction
	case "left":
		x = x - 1
		if x < 0 {
			x = 0
		}
		result = "x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y) + ", direction: " + direction
	case "right":
		x = x + 1
		if x > 100 {
			x = 100
		}
		result = "x: " + strconv.Itoa(x) + ", y: " + strconv.Itoa(y) + ", direction: " + direction
	case "up":
		y = y - 1
		if y < 0 {
			y = 0
		}
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
