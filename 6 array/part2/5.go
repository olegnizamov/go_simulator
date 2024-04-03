/*
	У вас есть переменныe data1, data2, которые содержат входные пользовательские данные.
	data1 - массив из элементов типа данных int.
	data2 - массив из элементов типа данных int.
	Напишите код, который сравнивает массив data1 с data2 и записывает разницу между ними в порядке возрастания в виде строки через запятую в переменную result.
	Если два списка одинаковые, тогда переменная result должна содержать сообщение: Массивы одинаковые
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
	data1, data2 := ReadInput()
	var result string
	var arrResult []int

	for _, v1 := range data1 {
		var findElement bool = false
		for _, v2 := range data2 {
			if v1 == v2 {
				findElement = true
				break
			}
		}
		if !findElement {
			arrResult = append(arrResult, v1)
		}
	}

	if len(arrResult) == 0 {
		fmt.Println("Массивы одинаковые")
	}

	sort.Ints(arrResult)

	var strArr []string
	for _, num := range arrResult {
		strNum := strconv.Itoa(num)
		strArr = append(strArr, strNum)
	}
	result = strings.Join(strArr, ", ")

	fmt.Println(result)
}

func ReadInput() ([]int, []int) {
	var data1 []int
	var data2 []int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	json.Unmarshal([]byte(values[0]), &data1)
	json.Unmarshal([]byte(values[1]), &data2)
	return data1, data2
}
