package leet16

import (
	"math"
	"sort"
)

func ThreeSumClosest(in []int, target int) int {
	if len(in) < 3 {
		return 0
	}

	closest := in[0] + in[1] + in[2]
	diff := math.Abs(float64(closest - target))

	sort.Ints(in)

	for k := 0; k < len(in)-2; k++ {
		i := k + 1
		j := len(in) - 1

		for i < j {
			sum := in[k] + in[i] + in[j]
			newDiff := math.Abs(float64(sum - target))
			if diff > newDiff {
				diff = newDiff
				closest = sum
			}
			if sum < target {
				i++
			} else {
				j--
			}
		}

	}

	return closest
}
