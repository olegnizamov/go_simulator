/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	data - массив из элементов типа данных string.
	Напишите код, который обращает порядок следования элементов массива data и записывает результат через запятую в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := ReadInput()
	var result string
	var resultArray []string

	for i := len(data) - 1; i >= 0; i-- {
		resultArray = append(resultArray, data[i])
	}
	result = strings.Join(resultArray, ", ")

	fmt.Println(result)
}

func ReadInput() []string {
	var data []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
