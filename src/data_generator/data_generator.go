package data_generator

import (
	"math/rand"
	"time"
)

func Generate(length int) ([]float64, []float64) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ts1 := make([]float64, length)
	ts2 := make([]float64, length)

	for i := 0; i < length; i++ {
		ts1[i] = r.Float64()
		ts2[i] = r.Float64()
	}

	return ts1, ts2

}
