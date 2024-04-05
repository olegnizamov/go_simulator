/*
	У вас есть переменная message которая содержит входные пользовательские данные.
	Напишите код, который:
	Читает строку message  вида: a2b3c1
	Воссоздает массив: a2b3c1 -> ["a","a","b","b","b","c"]
	Записывает воссозданный массив через запятую в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	message := ReadInput()
	var result string

	var strArr []string

	for i := 0; i < len(message); i += 2 {
		char := string(message[i])
		count, _ := strconv.Atoi(string(message[i+1]))

		for j := 0; j < count; j++ {
			strArr = append(strArr, char)
		}
	}
	result = strings.Join(strArr, ", ")

	fmt.Println(result)
}

func ReadInput() string {
	var message string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message = scanner.Text()
	return message
}
