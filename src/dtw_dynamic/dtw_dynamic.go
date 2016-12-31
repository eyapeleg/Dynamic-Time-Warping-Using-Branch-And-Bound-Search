package dtw_dynamic

import (
	"fmt"
	"math"
)

const (
	MaxFloat64 = 1.797693134862315708145274237317043567981e+308
)

func Dtw(ts1 []float64, ts2 []float64) float64 {

	accCostMatrix := initAccCostMatrix(ts1, ts2)

	for i := 1; i < len(ts1)+1; i++ {
		for j := 1; j < len(ts2)+1; j++ {
			distance := math.Abs(ts1[i-1] - ts2[j-1])
			fmt.Println("i: ", i, ", j:", j)
			accCostMatrix[i][j] = distance + min(accCostMatrix[i-1][j], accCostMatrix[i][j-1], accCostMatrix[i-1][j-1])
		}
	}

	return accCostMatrix[len(ts1)][len(ts2)]
}

func initAccCostMatrix(ts1 []float64, ts2 []float64) [][]float64 {
	accCostMatrix := make([][]float64, len(ts1)+1)
	for i := range accCostMatrix {
		accCostMatrix[i] = make([]float64, len(ts2)+1)
		for j := range accCostMatrix[i] {
			accCostMatrix[i][j] = MaxFloat64
		}
	}
	accCostMatrix[0][0] = 0
	return accCostMatrix
}

func min(nums ...float64) float64 {
	min := nums[0]
	for _, value := range nums {
		if value < min {
			min = value
		}
	}
	return min
}
