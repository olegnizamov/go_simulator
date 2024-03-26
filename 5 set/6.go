/*
	У вас есть переменные A, B, которые содержат входные пользовательские данные.
	Напишите код, который находит симметрическую разность множеств A△B, и записывает отсортированный результат в порядке возрастания в переменную result.

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
	A, B := ReadInput()
	var result string
	var resultSlice []int
	var hasAtSlice bool
	var valuesText []string

	//добавляем slice a
	for i := range A {
		hasAtSlice = false
		for j := range B {
			if A[i] == B[j] {
				hasAtSlice = true
				continue
			}
		}
		if !hasAtSlice {
			resultSlice = append(resultSlice, A[i])
		}
	}

	//добавляем slice b
	for i := range B {
		hasAtSlice = false
		for j := range A {
			if B[i] == A[j] {
				hasAtSlice = true
				continue
			}
		}
		if !hasAtSlice {
			resultSlice = append(resultSlice, B[i])
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

func ReadInput() ([]int, []int) {
	var A []int
	var B []int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	json.Unmarshal([]byte(values[0]), &A)
	json.Unmarshal([]byte(values[1]), &B)
	return A, B
}
