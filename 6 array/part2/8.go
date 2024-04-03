/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	data - массив из элементов типа данных int.
	Напишите код, который будет находить четные и нечетные числа и записывать результат в виде строки:
	(четные) (нечетные) в переменную result.
	Важно! Четные и нечетные числа должны быть отсортированы в порядке возрастания.
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
	data := ReadInput()
	var arrEven []int
	var arrOdd []int
	var result string

	for _, v := range data {
		if v%2 == 0 {
			arrEven = append(arrEven, v)
		} else if v%2 == 1 {
			arrOdd = append(arrOdd, v)
		}
	}
	sort.Ints(arrEven)
	sort.Ints(arrOdd)

	result = "("

	var strArrEven []string
	for _, num := range arrEven {
		strNum := strconv.Itoa(num)
		strArrEven = append(strArrEven, strNum)
	}
	result = result + strings.Join(strArrEven, ", ")
	result = result + ") ("

	var strArrOdd []string
	for _, num := range arrOdd {
		strNum := strconv.Itoa(num)
		strArrOdd = append(strArrOdd, strNum)
	}
	result = result + strings.Join(strArrOdd, ", ")
	result = result + ")"

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
