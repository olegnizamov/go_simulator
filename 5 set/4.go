/*
	У вас есть переменные A, e, которые содержат входные пользовательские данные.
	Напишите код, который проверяет есть ли элемент e в множестве A и записывает логический результат в переменную result.
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
	var resultSlice []int
	var result string
	var hasAtSlice bool
	valuesText := []string{}

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
