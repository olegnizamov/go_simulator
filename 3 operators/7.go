/*
У вас есть переменная num, которая содержит входные пользовательские данные.

Напишите код, который определяет является ли число четным или нечетным и записывает результат в переменную result.

Если число четное result = "четное" иначе result = "нечетное"
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num := ReadInput1()
	var result string

	if num%2 == 0 {
		result = "четное"
	} else {
		result = "нечетное"
	}

	fmt.Println(result)
}

func ReadInput1() int {
	var num int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ = strconv.Atoi(scanner.Text())

	return num
}
