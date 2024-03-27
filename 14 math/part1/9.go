/*
	У вас есть переменные a, b, которые содержат входные пользовательские данные.
	a – длина.
	b – ширина.
	Напишите код, который находит площадь прямоугльника и записывает результат в переменную result.
	Формула расчета площади прямоугольника: S=a⋅b, где S — площадь;  a,b — длина и ширина.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateSquareArea(a, b int) int {
	return a * b
}

func main() {
	a, b := ReadInput()
	result := calculateSquareArea(a, b)
	fmt.Println(result)
}

func ReadInput() (int, int) {
	var a, b int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	a, _ = strconv.Atoi(values[0])
	b, _ = strconv.Atoi(values[1])
	return a, b
}
