/*
	У вас есть переменная data которая содержит входные пользовательские данные.
	data - массив из элементов типа данных bool.
	Напишите код, который проверяет все ли элементы в массиве false и записывает логический результат в переменную result.
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

	for _, value := range data {
		result = result && !value
	}

	fmt.Println(result)
}

func ReadInput() []bool {
	var data []bool
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
