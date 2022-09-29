package hungarianAlgolang

// general useful things
func min(input []int) int {
	min := 0

	for i, v := range input {
		if i == 0 {
			min = v
		} else if v < min {
			min = v
		}
	}

	return min
}

func arrSum(s []int) int {

	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum

}

func getCol(j int, m [][]int) []int {

	col := []int{}

	for _, v := range m {
		col = append(col, v[j])
	}

	return col

}

func copyMatrix(m [][]int) [][]int {
	copy := [][]int{}
	for _, r := range m {
		row := []int{}
		for _, v := range r {
			row = append(row, v)
		}
		copy = append(copy, row)
	}
	return copy
}

// relating to step 1
func subtractSmallest(m [][]int) [][]int {

	newM := [][]int{}

	for _, r := range m {
		s := min(r)
		newR := []int{}
		for _, v := range r {
			newR = append(newR, v-s)
		}
		newM = append(newM, newR)
	}

	return newM
}

func checkValidStar(i int, j int, a [][]int) bool {

	r := a[i]
	c := getCol(j, a)

	if arrSum(r) == 0 && arrSum(c) == 0 {
		return true
	}
	return false

}

func findStars(m [][]int, c [][]int) [][]int {

	for i, r := range c {
		for j, v := range r {
			if v == 0 {
				valid := checkValidStar(i, j, m)
				if valid {
					m[i][j] = 1
				}
			}
		}
	}

	return m
}

func smallestUncovered(c [][]int, covR []int, covC []int) int {

	min := 999
	for i, r := range c {
		for j, v := range r {
			if covR[i] != 1 && covC[j] != 1 && v < min {
				min = v
			}
		}
	}

	return min
}

func subSmallestUnc(c [][]int, s int, covC []int) [][]int {

	for _, r := range c {
		for j, v := range r {
			if covC[j] != 1 {
				v = v - s
			}
		}
	}

	return c
}

func hasUncZero(c [][]int, covC []int) bool {

	for _, r := range c {
		for j, v := range r {
			if covC[j] != 1 && v == 0 {
				return true
			}
		}
	}

	return false
}

func rowHasStar(r []int) bool {

	for _, v := range r {
		if v == 1 {
			return true
		}
	}

	return false
}

func getRowPrime(s []int) int {

	idx := -1

	for i, v := range s {
		if v == 2 {
			idx = i
		}
	}

	return idx

}

func getColStar(s []int) int {

	idx := -1

	for i, v := range s {
		if v == 1 {
			idx = i
		}
	}

	return idx

}

func starPrimeSeries(i int, j int, a [][]int) [][]int {

	series := [][]int{{i, j}}
	lastPrime := false
	// curR := i
	curC := j

	for !lastPrime {

		col := getCol(curC, a)
		if getColStar(col) == -1 {
			lastPrime = true
		} else {
			starRow := getColStar(getCol(curC, a))
			starCoord := []int{starRow, curC}

			colPrime := getRowPrime(a[starRow])
			primeCoord := []int{starRow, colPrime}

			series = append(series, starCoord, primeCoord)
			// curR = starRow
			curC = colPrime
		}
	}

	return series
}
