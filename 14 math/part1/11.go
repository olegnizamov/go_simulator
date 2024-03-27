/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	Напишите код, который находит минимальное число в списке data и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func findMin(data []int) int {
	var minElement int = data[0]

	for i := 1; i < len(data); i++ {
		if minElement > data[i] {
			minElement = data[i]
		}
	}

	return minElement
}

func main() {
	data := ReadInput()
	result := findMin(data)
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
