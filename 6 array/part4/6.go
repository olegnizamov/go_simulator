/*
У вас есть переменная data которая содержит входные пользовательские данные.
data - массив из элементов типа данных int.
Напишите код, который умножает все числовые значения элементов массива data на минимальный элемент массива data и записывает новый массив в переменную result.
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func processData(data []int) []int {
	var minElement int = 10000
	var result []int

	for _, v := range data {
		if v < minElement {
			minElement = v
		}
	}

	for _, v := range data {
		result = append(result, v*minElement)
	}

	return result
}

func main() {
	data := ReadInput()
	result, _ := json.Marshal(processData(data))
	fmt.Println(string(result))
}

func ReadInput() []int {
	var data []int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
