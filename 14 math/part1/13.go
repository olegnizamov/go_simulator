/*
	У вас есть переменная data, которая содержит входные пользовательские данные.
	data - это список из вложенных списков коробок, каждая переставлена в виде массива из трех элементов (длина, ширина, высота).
	Напишите код, который подсчитывает общий объем всех этих коробок вместе взятых и записывает результат в переменную result.
	Формула для подсчета объема коробки: V=l⋅w⋅h

*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func totalVolume(data [][]int) int {

	var resultVolume int

	for _, volumeIndex := range data {
		var currentVolume int = 1
		for _, volumeElement := range volumeIndex {
			currentVolume = currentVolume * volumeElement
		}

		resultVolume = resultVolume + currentVolume
	}

	return resultVolume
}

func main() {
	data := ReadInput()
	result := totalVolume(data)
	fmt.Println(result)
}

func ReadInput() [][]int {
	var data [][]int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
