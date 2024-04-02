/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	data - думерный массив из элементов типа данных int.
	Напишите код, который находит максимальное число в двумерном массиве произвольного размера и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := ReadInput()
	var result int = 0

	for i := range data {
		for j := range data[i] {
			if result < data[i][j] {
				result = data[i][j]
			}
		}
	}

	fmt.Println(result)
}

func ReadInput() [][]int {
	var data [][]int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
