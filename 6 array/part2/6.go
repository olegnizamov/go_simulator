/*
	У вас есть переменныe size, position, data, которые содержат входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который заполняет массив data до нужного размера size нулями (0) в зависимости от значения position, которое может принимать значения: left или right.

	Важно!
	1 Если размер массива data  больше размера size, тогда в переменную result записываем сообщение: Неверный размер
    2 Если размер массива data  равен значению переменной size, тогда в переменную result записываем: Массив data в виде строки через запятую.
	3 Если значение переменной position не равно left или right, тогда в переменную result записываем сообщение: Неверная позиция
 	4 Сначала проверяем на корректность размера, потом проверяем на корректность позиции.

	Результат записать в виде строки через запятую в переменную result.
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
	size, position, data := ReadInput()
	var result string

	if len(data) > size {
		result = "Неверный размер"
	}

	if position != "left" && position != "right" {
		result = "Неверная позиция"
	}

	if len(data) == size {
		var strArr []string
		for _, num := range data {
			strNum := strconv.Itoa(num)
			strArr = append(strArr, strNum)
		}
		result = strings.Join(strArr, ", ")
	}

	var resultArray []int = make([]int, size)
	diff := size - len(data)
	diffIndex := size - len(data)
	currentArrayIndex := len(data)

	for i := 0; i < size; i++ {
		if position == "left" {
			if diffIndex > 0 {
				resultArray[i] = 0
				diffIndex--
			} else {
				resultArray[i] = data[i-diff]
			}
		} else if position == "right" {
			if currentArrayIndex == 0 && diffIndex > 0 {
				resultArray[i] = 0
				diffIndex--
			} else {
				resultArray[i] = data[i]
				currentArrayIndex--
			}
		}
	}

	var strArr []string
	for _, num := range resultArray {
		strNum := strconv.Itoa(num)
		strArr = append(strArr, strNum)
	}
	result = strings.Join(strArr, ", ")

	fmt.Println(result)
}

func ReadInput() (int, string, []int) {
	var size int
	var position string
	var data []int
	var values []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	size, _ = strconv.Atoi(values[0])
	position = values[1]
	json.Unmarshal([]byte(values[2]), &data)
	return size, position, data
}
