/*
	У вас есть переменные k, data которые содержат входные пользовательские данные.
	data - двумерный массив из элементов типа данных int.
	Напишите код, который находит число k в двумерном в массиве data и записывает логический результат в переменную result.
	Пример двумерного массива:

	[
  		[1,2,3],
  		[4,5,6],
  		[7,8,9]
	]
Важно! Массив data может быть любого размера.
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
	k, data := ReadInput()
	var result bool

	result = FindNumber(k, data)

	fmt.Println(result)
}

func FindNumber(k int, data [][]int) bool {

	var result bool = false

	for i := range data {
		for j := range data[i] {
			if k == data[i][j] {
				result = true
				break
			}
		}
	}

	return result
}

func ReadInput() (int, [][]int) {
	var k int
	var data [][]int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	k, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &data)
	return k, data
}
