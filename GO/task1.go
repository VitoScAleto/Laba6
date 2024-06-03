package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func task1() {
	fmt.Println("Task 1")
	quantityLine, quantityColumns := 5, 6
	maxSrednee, numberLine := 0, 0

	randomMatrix := make([][]int, quantityLine)
	for i := range randomMatrix {
		randomMatrix[i] = make([]int, quantityColumns)
		for j := range randomMatrix[i] {
			randomMatrix[i][j] = GetRandomNumber(0, 100)
		}
	}

	for _, i := range randomMatrix {
		for _, j := range i {
			fmt.Printf("%4d", j)
		}
		fmt.Println()
	}

	for i := 0; i < quantityLine; i++ {
		sredneeArithmetic := 0
		for j := 0; j < quantityColumns; j++ {
			sredneeArithmetic += randomMatrix[i][j]
		}
		if maxSrednee < sredneeArithmetic {
			maxSrednee = sredneeArithmetic
			numberLine = i
		}
		fmt.Printf("Номер строки %d среднее арифметическое = %.2f\n", i, float64(sredneeArithmetic)/float64(quantityColumns))
	}
	fmt.Printf("Наибольшее среднее арифметическое в строке %d\n", numberLine)
}

func task2() {
	fmt.Println("Task 2")

	sizeSquareMatrix := 15
	squareMatrix := make([][]int, sizeSquareMatrix)
	for i := range squareMatrix {
		squareMatrix[i] = make([]int, sizeSquareMatrix)
		for j := range squareMatrix[i] {
			squareMatrix[i][j] = GetRandomNumber(-100, 100)
			fmt.Printf("%4d", squareMatrix[i][j])
		}
		fmt.Println()
	}

	for i := 0; i < sizeSquareMatrix; i++ {
		columns := make([]int, sizeSquareMatrix)
		for j := 0; j < sizeSquareMatrix; j++ {
			columns[j] = squareMatrix[j][i]
		}
		if i%2 == 0 {
			Sort(columns, true)
		} else {
			Sort(columns, false)
		}
		for j := 0; j < sizeSquareMatrix; j++ {
			squareMatrix[j][i] = columns[j]
		}
	}

	fmt.Println()
	for _, i := range squareMatrix {
		for _, j := range i {
			fmt.Printf("%4d", j)
		}
		fmt.Println()
	}

	sumRight, sumLeft := 0, 0
	for i := 0; i < len(squareMatrix); i++ {
		for j := i; j < len(squareMatrix); j++ {
			if squareMatrix[i][j] > 0 {
				sumRight++
			}
		}
	}

	for i := len(squareMatrix) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if squareMatrix[i][j] > 0 {
				sumLeft++
			}
		}
	}

	fmt.Printf("Количество положительных элементов правой половины = %d\tколичество положительных элементов левой половины = %d\n", sumRight, sumLeft)
	if sumRight > sumLeft {
		fmt.Println("Правая половина содержит больше положительных элементов")
	} else {
		fmt.Println("Левая половина содержит больше положительных элементов")
	}
}

func Sort(arr []int, asc bool) {
	if asc {
		for i := 0; i < len(arr)-1; i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	} else {
		for i := 0; i < len(arr)-1; i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] < arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}
}

func main() {
	task1()
	fmt.Println()
	task2()
}
