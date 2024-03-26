/*
	У вас есть переменная A, которая содержит входные пользовательские данные.
	Напишите код, который удаляет все четные числа из множества A и записывает отсортированный результат в порядке возрастания в переменную result.
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
	A := ReadInput()
	var result string
	var resultSlice []int
	var valuesText []string
	for i := range A {
		if i%2 == 0 {
			resultSlice = append(resultSlice, A[i])
		}
	}

	sort.Ints(resultSlice)
	for i := range resultSlice {
		number := resultSlice[i]
		valuesText = append(valuesText, strconv.Itoa(number))
	}
	result = strings.Join(valuesText, ", ")

	fmt.Println(result)
}

func ReadInput() []int {
	var A []int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &A)
	return A
}
