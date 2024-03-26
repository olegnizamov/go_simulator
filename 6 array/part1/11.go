/*
	У вас есть переменная tags, которая содержит входные пользовательские данные.
	tags - массив из элементов типа данных string.
	Напишите код, который сортирует строковые элементы массива tags в порядке возрастания, отсеивает дубликаты и записывает результат через запятую в переменную result.
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
	tags := ReadInput()
	var result string
	slices.Sort(tags)
	tags = slices.Compact(tags)
	result = strings.Join(tags, ", ")
	fmt.Println(result)
}

func ReadInput() []string {
	var tags []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &tags)
	return tags
}
