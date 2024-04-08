/*
	У вас есть переменная grid которая содержит входные пользовательские данные.
	grid - двумерный массив из элементов типа данных int.
	Напишите код, который находит сумму всех чисел по двум диагоналям двумерного массива grid и записывает результат в переменную result.
	Обратите внимание на то, что ваш код должен работать для любого размера двухмерного массива grid например: 2х2, 3х3, 4х4 и тд...
	Важно! Вам необходимо убедиться в том, что размер массива по высоте и ширине одинаковый! Если размеры массива grid тогда записываем -1 в переменную result.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	grid := ReadInput()
	var result int

	rows := len(grid)
	columns := len(grid[0])

	fmt.Println(rows)
	fmt.Println(columns)

	if rows != columns {
		result = -1
	} else {
		for i := range grid {
			for j := range grid[i] {
				fmt.Println(grid[i][j])
				result += grid[i][j]
			}
		}
	}

	fmt.Println(result)
}

func ReadInput() [][]int {
	var grid [][]int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &grid)
	return grid
}
