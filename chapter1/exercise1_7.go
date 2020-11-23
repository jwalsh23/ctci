package chapter1

func Rotate(matrix [][]int) [][]int{
	returnVal := make([][]int,0)
	for i := 0; i < len(matrix); i++ {
		matrixRow := make([]int, 0)
		for j := len(matrix[i]) -1; j > -1; j-- {
			matrixRow = append(matrixRow, matrix[j][i])
		}
		returnVal = append(returnVal, matrixRow)
	}
	return returnVal
}
func RotateInPlace(matrix [][]int) {
	//for layer := 0; layer < len(matrix) -1; layer ++ {
	//	first := layer
	//	last := len(matrix) -1 -layer
	//	for i := 0; i < last; i++ {
	//		top := matrix[first][i]
	//		matrix[first][i] = matrix[last][i]
	//	}
	//}
}