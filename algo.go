package hungarianAlgolang

import (
	"fmt"
	"log"
)

func solve(m [][]int) ([]int, error) {
	v, err := validate(m)
	if err != 0 {
		log.Fatal(v)
	}
	original := copyMatrix(m)

	costs := m
	mask := make([][]int, len(m))
	for i := range mask {
		mask[i] = make([]int, len(m))
	}
	covR := make([]int, len(m))
	covC := make([]int, len(m))
	minUnc := smallestUncovered(costs, covR, covC)
	starPrimeStart := []int{}

	step := 1
	done := false
	result := []int{}

Reentry:

	for !done {
		switch step {

		case 1:
			fmt.Println("step1")
			costs = subtractSmallest(costs)
			step = 2

		case 2:
			fmt.Println("step2")
			mask = findStars(mask, costs)
			step = 3

		case 3:
			fmt.Println("step3")
			for _, r := range mask {
				for j, v := range r {
					if v == 1 {
						covC[j] = 1
					}
				}
			}
			numCov := arrSum(covC)
			if numCov == len(costs) {
				for i, r := range mask {
					for j, v := range r {
						if v == 1 {
							result = append(result, original[i][j])
						}
					}
				}
				done = true
			} else {
				step = 4
			}

		case 4:
			fmt.Println("step4")
			// pre-process to make sure there are uncovered zeros
			uncZero := hasUncZero(costs, covC)
			if !uncZero {
				minUnc = smallestUncovered(costs, covR, covC)
				costs = subSmallestUnc(costs, minUnc, covC)
			}
			// step 4 proper
			for i, r := range costs {
				for j, v := range r {
					if v == 0 && covC[j] != 1 {
						mask[i][j] = 2
						if !rowHasStar(r) {
							step = 5
							starPrimeStart = []int{i, j}
							goto Reentry
						} else {
							covR[i] = 1
							covC[j] = 0
						}
					}
				}
			}
			minUnc = smallestUncovered(costs, covR, covC)
			step = 6

		case 5:
			fmt.Println("step5")
			s := starPrimeSeries(starPrimeStart[0], starPrimeStart[1], mask)
			if len(s) > 1 {
				for i, v := range s {
					if i%2 == 0 {
						mask[v[0]][v[1]] = 1
					} else {
						mask[v[0]][v[1]] = 0
					}
				}
			} else {
				for i, r := range mask {
					for j := range r {
						if s[0][0] != i && s[0][1] != j {
							mask[i][j] = 0
						}
					}
				}
			}
			for i := range covR {
				covR[i] = 0
				covC[i] = 0
			}
			step = 3

		case 6:
			fmt.Println("step6")
			for i, r := range costs {
				for j, v := range r {
					if covR[i] == 1 {
						costs[i][j] = v + minUnc
					} else if covR[i] == 0 && covC[j] == 0 {
						costs[i][j] = v - minUnc
					}
				}
			}
			step = 4

		}
	}

	return result, nil
}
