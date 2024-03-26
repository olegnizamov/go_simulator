/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	data - массив из элементов типа данных string.
	Напишите код, который сортирует строковые элементы массива data в порядке возрастания и записывает результат через запятую в переменную result.
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
	sort.Strings(data)
	result = strings.Join(data, ", ")
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
