/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	Напишите код, который находит максимальное число в списке data и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func findMax(data []int) int {

	var maxElement int = data[0]

	for i := 1; i < len(data); i++ {
		if maxElement < data[i] {
			maxElement = data[i]
		}
	}

	return maxElement
}

func main() {
	data := ReadInput()
	result := findMax(data)
	fmt.Println(result)
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
