/*
	У вас есть переменные num, data которые содержат входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который определяет все ли числа в массиве data больше num и записывает логический результат в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	num, data := ReadInput()
	var result bool = true

	for i := range data {
		if data[i] < num {
			result = false
			break
		}
	}

	fmt.Println(result)
}

func ReadInput() (int, []int) {
	var num int
	var data []int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	num, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &data)
	return num, data
}
