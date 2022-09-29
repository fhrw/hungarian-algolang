package hungarianAlgolang

func validate(matrix [][]int) (message string, err int) {

	n := len(matrix)

	if n <= 0 {
		return "matrix is empty", 1
	}

	m := len(matrix[0])
	if m != n {
		return "matrix is not square", 1
	}

	return "matrix is valid", 0
}
