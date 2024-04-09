/*
	У вас есть переменная data которая содержит входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который выбирает элементы массива data, которые больше, чем элементы, стоящие перед ними и записывает новый массив в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func getGreaterThanPrevious(data []int) []int {
	var result []int
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			result = append(result, data[i])
		}
	}

	return result
}

func main() {
	data := ReadInput()
	result, _ := json.Marshal(getGreaterThanPrevious(data))
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
