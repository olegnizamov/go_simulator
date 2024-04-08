/*
	У вас есть переменная data которая содержит входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который выбирает первый отрицательный элемент из списка data и записывает результат в переменную result.
	Если отрицательный элемент не найден, тогда возвращаем 0
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func findFirstNegative(data []int) int {

	var result int = 0

	for _, v := range data {
		if v < 0 {
			result = v
			break
		}
	}
	return result
}

func main() {
	data := ReadInput()
	result := findFirstNegative(data)
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
