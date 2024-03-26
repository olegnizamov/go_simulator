/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который сортирует числовые элементы массива data в порядке убывания и записывает результат через запятую в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data := ReadInput()
	var result string
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
	result = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data)), ", "), "[]")
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
