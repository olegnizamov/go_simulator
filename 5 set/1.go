/*
	У вас есть переменные A, e, которые содержат входные пользовательские данные.
	Напишите код, который добавляет элемент e в множество A и записывает отсортированный результат в порядке возрастания в переменную result.
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
	A, e := ReadInput()
	var result string
	valuesText := []string{}
	A = append(A, e)
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	for i := range A {
		number := A[i]
		valuesText = append(valuesText, strconv.Itoa(number))
	}
	result = strings.Join(valuesText, ", ")
	fmt.Println(result)
}

func ReadInput() ([]int, int) {
	var A []int
	var e int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	json.Unmarshal([]byte(values[0]), &A)
	e, _ = strconv.Atoi(values[1])
	return A, e
}
