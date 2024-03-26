/*
	У вас есть переменные a, h, которые содержат входные пользовательские данные.
	a – основание.
	h – высота.
	Напишите код, который находит площадь треугольника и записывает результат в переменную result.
	Самая простая формула для расчета площади это произведение основания и высоты треугольника, поделенное на 2
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateTriangleArea(a, h int) int {
	return (a * h) / 2
}

func main() {
	a, h := ReadInput()
	result := calculateTriangleArea(a, h)
	fmt.Println(result)
}

func ReadInput() (int, int) {
	var a, h int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	a, _ = strconv.Atoi(values[0])
	h, _ = strconv.Atoi(values[1])
	return a, h
}
