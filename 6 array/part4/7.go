/*
	У вас есть переменная data которая содержит входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который выбирает элементы массива data, которые больше среднего арифметического массива data и записывает новый массив в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func filterAboveAverage(data []int) []int {
	var average int = 0
	var result []int

	for _, v := range data {
		average = average + v
	}
	average = average / len(data)

	for _, v := range data {
		if average < v {
			result = append(result, v)
		}

	}

	return result
}

func main() {
	data := ReadInput()
	result, _ := json.Marshal(filterAboveAverage(data))
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
