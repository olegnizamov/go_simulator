/*
	У вас есть переменная data которая содержит входные пользовательские данные.
	data - массив из элементов типа данных float64.
	Напишите код, который выбирает максимальный по модулю элемент из массива data и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
)

func findMaxAbsolute(data []float64) float64 {
	var maxElement float64 = 0.0

	for _, v := range data {
		if math.Abs(v) > math.Abs(maxElement) {
			maxElement = v

		}
	}

	return maxElement
}

func main() {
	data := ReadInput()
	result := findMaxAbsolute(data)
	fmt.Println(result)
}

func ReadInput() []float64 {
	var data []float64
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
