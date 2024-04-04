/*
	У вас есть переменные data1 и data2 которые содержат входные пользовательские данные.
	data1 - массив из элементов типа данных int.
	data2 - массив из элементов типа данных int.
	Напишите код, который складывает все числа одного массива с числами второго массива и записывает результат через запятую в переменную result.
	Размеры массивов одинаковые.
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
	data1, data2 := ReadInput()
	var result string

	var strArr []string
	for i := range data1 {
		sum := data1[i] + data2[i]
		strNum := strconv.Itoa(sum)
		strArr = append(strArr, strNum)
	}

	result = result + strings.Join(strArr, ", ")

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
