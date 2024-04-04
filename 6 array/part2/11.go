/*
	У вас есть переменные health и items, которые содержат входные пользовательские данные.
	items - массив из элементов типа данных string.
	Напишите код, который увеличивает значение переменной health в зависимости от того, сколько зелья находится в списке items и записывается результат в переменную result.

	Важно!
	Одно "Зелье" это +10 к health
	Итоговое значение переменной result не должно быть больше 100. Если значение переменной result получается больше 100, тогда устанавливаем значении переменной result равной 100.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	health, items := ReadInput()
	var result int = health

	for _, item := range items {
		if item == "Зелье" {
			result = result + 10
		}
	}

	if result > 100 {
		result = 100
	}

	fmt.Println(result)
}

func ReadInput() (int, []string) {
	var health int
	var items []string
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	health, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &items)
	return health, items
}
