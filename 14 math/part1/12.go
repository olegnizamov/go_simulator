/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	Напишите код, который складывает минимальное и максимальное число в списке data и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func findMin(data []int) int {
	max := data[0]
	for _, num := range data {
		if num < max {
			max = num
		}
	}
	return max
}

func findMax(data []int) int {
	max := data[0]
	for _, num := range data {
		if num > max {
			max = num
		}
	}
	return max
}

func sumMinAndMax(data []int) int {
	return findMax(data) + findMin(data)
}

func main() {
	data := ReadInput()
	result := sumMinAndMax(data)
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
