/*
	У вас есть переменная data которая содержит входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который заменяет числовые значения элементов массива data на противоположные и записывает новый массив в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func negateData(data []int) []int {

	var result []int
	for _, v := range data {
		result = append(result, -v)
	}
	return result
}

func main() {
	data := ReadInput()
	result, _ := json.Marshal(negateData(data))
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
