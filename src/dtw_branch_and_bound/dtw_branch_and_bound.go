package dtw_branch_and_bound

import (
	"math"
	"runtime"
	"sync"
)

type Direction string

const (
	first  Direction = "1"
	second Direction = "2"
	both   Direction = "both"
)

const (
	MaxFloat64 = 1.797693134862315708145274237317043567981e+308
)

var (
	currentUpperBound float64 = MaxFloat64
	mutex             sync.Mutex
)

func Dtw(ts1 []float64, ts2 []float64) float64 {
	runtime.GOMAXPROCS(4)

	currentUpperBound = calculateInitialUpperBound(0, ts1, ts2)
	//fmt.Println("Initial Upper Bound: ", currentUpperBound)

	return search(both, 0, ts1, ts2)
	//fmt.Println("DTW distance: ", dtwDistance)
}

func search(direction Direction, accDist float64, ts1 []float64, ts2 []float64) float64 {

	accDist += math.Abs(ts1[0] - ts2[0])

	if accDist >= getUpperBound() {
		//fmt.Println("Exceeded Upper Bound: ", accDist)
		return MaxFloat64
	}

	if len(ts1) == 1 && len(ts2) == 1 {
		registerUpperBound(accDist)
		return accDist
	}

	val1, val2, val3 := MaxFloat64, MaxFloat64, MaxFloat64
	var waitGrp sync.WaitGroup

	if len(ts1) > 1 && direction != second {
		waitGrp.Add(1)
		go func() {
			val1 = search(first, accDist, ts1[1:], ts2)
			waitGrp.Done()
		}()
	}

	if len(ts2) > 1 && direction != first {
		waitGrp.Add(1)
		go func() {
			val2 = search(second, accDist, ts1, ts2[1:])
			waitGrp.Done()
		}()
	}

	if len(ts1) > 1 && len(ts2) > 1 {
		waitGrp.Add(1)
		go func() {
			val3 = search(both, accDist, ts1[1:], ts2[1:])
			waitGrp.Done()
		}()
	}

	waitGrp.Wait()
	return min(val1, val2, val3)
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

func registerUpperBound(alternativeUpperBound float64) {
	mutex.Lock()
	if currentUpperBound > alternativeUpperBound {
		//fmt.Println("New Upper Bound: ", alternativeUpperBound)
		currentUpperBound = alternativeUpperBound
	}
	mutex.Unlock()
}

func getUpperBound() float64 {
	return currentUpperBound
}

func calculateInitialUpperBound(acc float64, ts1 []float64, ts2 []float64) float64 {
	acc += math.Abs(ts1[0] - ts2[0])

	if len(ts1) > 1 && len(ts2) > 1 {
		return calculateInitialUpperBound(acc, ts1[1:], ts2[1:])
	} else if len(ts1) > 1 {
		return calculateInitialUpperBound(acc, ts1[1:], ts2)
	} else if len(ts2) > 1 {
		return calculateInitialUpperBound(acc, ts1, ts2[1:])
	}
	return acc
}
