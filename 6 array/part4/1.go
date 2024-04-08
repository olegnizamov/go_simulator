/*
	У вас есть переменная data которая содержит входные пользовательские данные.
	data - массив из элементов произвольного типа данных.
	Напишите код, который заменяет все логические значения true на строковое "on", а все логические значения false на "off" и записывает результат в виде нового массива в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func booleanToString(data []interface{}) []interface{} {
	var result []interface{}

	for _, v := range data {
		if v == true {
			result = append(result, "on")
		} else if v == false {
			result = append(result, "off")
		} else {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	data := ReadInput()
	result, _ := json.Marshal(booleanToString(data))
	fmt.Println(string(result))
}

func ReadInput() []interface{} {
	var data []interface{}
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
