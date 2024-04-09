/*
	У вас есть переменная files которая содержит входные пользовательские данные.
	files - массив из элементов типа данных string.
	Напишите код, который находит все файлы в массив files с расширением .png и записывает новый массив в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func findPNGFiles(files []string) []string {
	var results []string

	for _, file := range files {
		if strings.Contains(file, ".png") {
			results = append(results, file)
		}
	}
	return results
}

func main() {
	files := ReadInput()
	result, _ := json.Marshal(findPNGFiles(files))
	fmt.Println(string(result))
}

func ReadInput() []string {
	var files []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &files)
	return files
}
