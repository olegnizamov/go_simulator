/*
	У вас есть переменные n, data которые содержат входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который умножает каждый элемент массива на значение переменной n и записывает отсортированный результат в порядке возрастания в виде строки через запятую в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	n, data := ReadInput()
	var result string
	var strArr []string
	sort.Ints(data)

	for i := range data {
		data[i] = data[i] * n
	}

	for _, num := range data {
		strNum := strconv.Itoa(num)
		strArr = append(strArr, strNum)
	}
	result = strings.Join(strArr, ", ")
	fmt.Println(result)
}

func ReadInput() (int, []int) {
	var n int
	var data []int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	n, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &data)
	return n, data
}
