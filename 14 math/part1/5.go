/*
	У вас есть переменная n, которая содержит входные пользовательские данные.
	Напишите код, который находит произведение цифр числа n и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func productOfDigits(n int) string {
	str := strconv.Itoa(n)
	var sum int = 1

	for i := 0; i < len(str); i++ {
		current, _ := strconv.Atoi(string(str[i]))
		sum = sum * current
	}

	return fmt.Sprintf("Произведение цифр числа %d равно %d", n, sum)
}

func main() {
	n := ReadInput()
	result := productOfDigits(n)
	fmt.Println(result)
}

func ReadInput() int {
	var n int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	n, _ = strconv.Atoi(input)
	return n
}
