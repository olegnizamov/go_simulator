/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который определяет все ли числа в массиве data нечетные и записывает логический результат в переменную result.
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
	var result bool = true

	for i := range data {
		if data[i]%2 == 0 {
			result = false
		}
	}

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
