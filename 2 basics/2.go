/*
	У вас есть переменная message, которая содержит входные пользовательские данные.
	Напишите код, который записывает значение переменной message в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	message := ReadInput()
	result := message
	fmt.Println(result)
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()
	return message
}
