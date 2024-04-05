/*
	У вас есть переменная data которая содержит входные пользовательские данные.
	data - массив из элементов типа данных string.
	Напишите код, который проходит по массиву data, и записывает символы вместе с числом их повторений в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data := ReadInput()
	var result string = data[0]

	var prevChar string = data[0]
	var countElements int = 0
	for _, num := range data {
		if prevChar != num {
			result = result + strconv.Itoa(countElements)
			result = result + num
			countElements = 1
			prevChar = num
		} else {
			countElements++
		}
	}

	result = result + strconv.Itoa(countElements)

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
