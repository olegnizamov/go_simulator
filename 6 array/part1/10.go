/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который сортирует числовые элементы массива data в порядке возрастания, отсеивает дубликаты и записывает результат через запятую в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	data := ReadInput()
	var result string
	slices.Sort(data)
	data = slices.Compact(data)
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
