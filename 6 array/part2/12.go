/*
	У вас есть переменные power и items, которые содержат входные пользовательские данные.
	items - массив из элементов типа данных string.
	Напишите код, который увеличивает значение переменной power в зависимости от того, сколько улучшений/бустов находится в списке items и записывается результат в переменную result.
	Улучшения:
		"Энергетик": power + 5
		"Кофе": power + 10
	Важно!

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
	power, items := ReadInput()
	var result int = power

	for _, item := range items {
		if item == "Кофе" {
			result = result + 10
		}
		if item == "Энергетик" {
			result = result + 5
		}
	}

	if result > 100 {
		result = 100
	}

	fmt.Println(result)
}

func ReadInput() (int, []string) {
	var power int
	var items []string
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	power, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &items)
	return power, items
}
