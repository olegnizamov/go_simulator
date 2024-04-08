/*
	У вас есть переменная data которая содержит входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который умножает все числовые значения элементов массива data на максимальный элемент массива data и записывает новый массив в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func multiplyMax(data []int) []int {

	var maxElement int = 0
	var result []int

	for _, v := range data {
		if v > maxElement {
			maxElement = v
		}
	}

	for _, v := range data {
		result = append(result, v*maxElement)
	}

	return result
}

func main() {
	data := ReadInput()
	result, _ := json.Marshal(multiplyMax(data))
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
