/*
	У вас есть переменные n и items, которые содержат входные пользовательские данные.
	items - массив из элементов типа данных int.
	Напишите код, который копирует элементы массива items из начала массива в конец массива и записывается результат в переменную result.
	n - число элементов которые необходимо скопировать.
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
	n, items := ReadInput()
	var result string
	firstNElements := items[:n]
	combinedArray := append(items, firstNElements...)
	var strArr []string
	for _, num := range combinedArray {
		strNum := strconv.Itoa(num)
		strArr = append(strArr, strNum)
	}
	result = result + strings.Join(strArr, ", ")

	fmt.Println(result)
}

func ReadInput() (int, []int) {
	var n int
	var items []int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	n, _ = strconv.Atoi(values[0])
	json.Unmarshal([]byte(values[1]), &items)
	return n, items
}
