package hungarianAlgolang

import (
	"fmt"
	"log"
)

func Solve(m [][]int) ([]int, error) {
	v, err := validate(m)
	if err != 0 {
		log.Fatal(v)
	}
	original := copyMatrix(m)
	_ = original

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
				for _, r := range mask {
					for j, v := range r {
						if v == 1 {
							result = append(result, j)
						}
					}
				}
				done = true
			} else {
				step = 4
			}

		case 4:
			fmt.Println("step4")
			// find an uncovered zero and prime it
			// If there are no uncovered zeros this should go to step 6
			// if there are this will have to go through till there are none
			// left uncovered.
			uncZero := hasUncZero(costs, covR, covC)

			// while there are uncovered zreos
			for uncZero {
				for i, r := range costs {
					for j, v := range r {
						if v == 0 && covR[i] == 0 && covC[j] == 0 {
							mask[i][j] = 2
							check, idx := rowHasStar(mask[i])
							if check {
								covR[i] = 1
								covC[idx] = 0
							} else {
								starPrimeStart = []int{i, j}
								step = 5
								goto Reentry
							}
						}
					}
				}
				uncZero = hasUncZero(costs, covR, covC)
			}

			minUnc = smallestUncovered(costs, covR, covC)
			step = 6

		case 5:
			fmt.Println("step5")
			s := starPrimeSeries(starPrimeStart[0], starPrimeStart[1], mask)
			// BUG: Should only unstar stars in the series
			// leave the other ones alone...
			// FIX: check prime rows for stars and unstar them
			for i, v := range s {
				if (i % 2) == 0 {
					check, starIdx := rowHasStar(mask[v[0]])
					if check {
						mask[v[0]][starIdx] = 0
					}
					mask[v[0]][v[1]] = 1
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
