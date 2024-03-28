/*
	У вас есть переменные h, m, s, которые содержат входные пользовательские данные.
	h - часы
	m - минуты
	s - секунды
	Напишите код, определяет значение с наибольшей продолжительностью и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestTime(h, m, s int) int {

	if (h*60 > m) && (h*3600 > s) {
		return h
	}

	if (h*60 > m) && (h*3600 < s) {
		return s
	}

	if (m > h*60) && (m*60 > s) {
		return m
	}

	if (m > h*60) && (m*60 < s) {
		return s
	}

	return h
}

func main() {
	h, m, s := ReadInput()
	result := longestTime(h, m, s)
	fmt.Println(result)
}

func ReadInput() (int, int, int) {
	var h, m, s int
	var values []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	h, _ = strconv.Atoi(values[0])
	m, _ = strconv.Atoi(values[1])
	s, _ = strconv.Atoi(values[2])
	return h, m, s
}
