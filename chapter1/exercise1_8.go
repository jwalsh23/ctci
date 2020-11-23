package chapter1

func Zero(matrix [][]int) {
	rowZerod := make([]bool, len(matrix))
	columnZerod := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowZerod[i] = true
				columnZerod[j] = true
			}
		}
	}
	for i, shouldZero := range rowZerod {
		if shouldZero {
			zeroRow(matrix[i])
		}
	}
	for i, shouldZero := range columnZerod {
		if shouldZero{
			zeroColumn(matrix, i)
		}
	}
}

func zeroRow(row []int) {
	for i := range row {
		row[i] = 0
	}
}

func zeroColumn(matrix [][]int, colNum int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][colNum] = 0
	}
}
