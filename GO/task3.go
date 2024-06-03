package main

import (
	"fmt"
	"math"
)

// Функция для выбора главного элемента в текущем столбце
func findPivot(matrix [][]float64, col, startRow int) int {
	maxVal := 0.0
	pivotRow := startRow

	for i := startRow; i < len(matrix); i++ {
		if math.Abs(matrix[i][col]) > maxVal {
			maxVal = math.Abs(matrix[i][col])
			pivotRow = i
		}
	}
	return pivotRow + 1 // возвращаем номер строки, начиная с 1
}

// Функция для решения системы линейных уравнений
func solveLinearSystem(matrix [][]float64) []float64 {
	n, m := len(matrix), len(matrix[0])-1
	workingMatrix := make([][]float64, n)
	for i := range workingMatrix {
		workingMatrix[i] = make([]float64, len(matrix[i]))
		copy(workingMatrix[i], matrix[i])
	}

	// Прямой ход
	fmt.Println("Прямой ход:")
	for col := 0; col < m; col++ {
		pivotRow := findPivot(workingMatrix, col, col)
		pivotValue := workingMatrix[pivotRow-1][col]
		fmt.Printf("Выбран главный элемент %.2f в строке %d\n", pivotValue, pivotRow)

		if pivotRow != col+1 {
			workingMatrix[col], workingMatrix[pivotRow-1] = workingMatrix[pivotRow-1], workingMatrix[col]
			fmt.Printf("Меняем местами строки %d и %d\n", col+1, pivotRow)
		}

		for i := col + 1; i < n; i++ {
			factor := workingMatrix[i][col] / workingMatrix[col][col]
			fmt.Printf("Вычитаем из строки %d %.2f\n", i+1, factor)

			for j := col; j < m+1; j++ {
				workingMatrix[i][j] -= factor * workingMatrix[col][j]
			}
		}

		fmt.Println("Промежуточная матрица после", col+1, "-го столбца:")
		for _, row := range workingMatrix {
			for _, elem := range row {
				if math.Abs(elem) < 3e-17 {
					fmt.Print("0 ")
				} else {
					fmt.Printf("%.2f ", elem)
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	// Обратный ход
	fmt.Println("Обратный ход:")
	solution := make([]float64, m)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < m; j++ {
			sum += workingMatrix[i][j] * solution[j]
		}
		solution[i] = (workingMatrix[i][m] - sum) / workingMatrix[i][i]
		fmt.Printf("x%d = %.2f\n", i+1, solution[i])
	}

	return solution
}

func checkSpectralRadius(C [][]float64) {
	spectralRadius := 0.0
	for i := 0; i < 4; i++ {
		sum := 0.0
		for j := 0; j < 4; j++ {
			sum += math.Abs(C[i][j])
		}
		if sum > spectralRadius {
			spectralRadius = sum
		}
	}

	fmt.Printf("\nСпектральный радиус матрицы C: %.2f\n", spectralRadius)

	if spectralRadius >= 1.0 {
		fmt.Println("Условие сходимости метода простых итераций не выполняется.")
	} else {
		fmt.Println("Условие сходимости метода простых итераций выполняется.")
	}
}

// Преобразование системы Ах=b к виду х=Сх+f
func formCanonicalSystem(A [][]float64, b []float64, C [][]float64, f []float64, n int) {
	for i := 0; i < n; i++ {
		f[i] = b[i] / A[i][i]
		for j := 0; j < n; j++ {
			if j != i {
				C[i][j] = -A[i][j] / A[i][i]
			} else {
				C[i][j] = 0.0
			}
		}
	}
}

// Метод простых итераций
func simpleIteration(C [][]float64, f []float64, x []float64, n int, epsilon float64) bool {
	xNew := make([]float64, n)
	k := 0
	maxDiff := 0.0
	maxIteration := 10

	// Задаем начальное приближение
	for i := range x {
		x[i] = 0.0
	}

	fmt.Printf("N%10sx1%10sx2%10sx3%10sx4%10sεn\n", "", "", "", "", "")

	for {
		// Вычисляем новое приближение
		for i := 0; i < n; i++ {
			sum := f[i]
			for j := 0; j < n; j++ {
				sum += C[i][j] * x[j]
			}
			xNew[i] = sum
		}

		// Проверяем условие остановки
		maxDiff = 0.0
		for i := 0; i < n; i++ {
			if math.Abs(xNew[i]-x[i]) > maxDiff {
				maxDiff = math.Abs(xNew[i] - x[i])
			}
			x[i] = xNew[i]
		}
		k++

		// Вывод результатов в таблицу
		fmt.Printf("%d%10.3f%10.3f%10.3f%10.3f%10.3f\n", k, x[0], x[1], x[2], x[3], maxDiff)

		if maxDiff <= epsilon {
			fmt.Println("Число итераций:", k)
			fmt.Println("Метод сходится.")
			return false
		}

		if k >= maxIteration {
			fmt.Println("Число итераций:", k, "(максимум)")
			fmt.Println("Метод расходится.")
			return true
		}
	}
}

func main() {
	matrix := [][]float64{
		{-1.21, -0.04, 0.21, -18.0, -1.24},
		{0.25, -1.23, 0.2, -0.09, 0.88},
		{-0.21, 0.2, 0.80, -0.13, 2.56},
		{0.15, -1.31, 0.06, 0.88, -1.21},
	}

	fmt.Println("\tРешение СЛАУ с выбором главного элемента:\n")
	solution := solveLinearSystem(matrix)
	fmt.Println("Матрица решений СЛАУ:")
	for _, x := range solution {
		fmt.Printf("%.2f ", x)
	}
}
