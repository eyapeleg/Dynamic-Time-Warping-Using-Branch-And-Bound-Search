package main

import (
	data_generator "AI/src/data_generator"
	dtw_branch_and_bound "AI/src/dtw_branch_and_bound"
	dtw_dynamic "AI/src/dtw_dynamic"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	smallTsLength   = 4
	largeTsLength   = 23
	numOfIterations = 10
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	outputDirPath := "." + string(filepath.Separator) + "output"
	os.Remove(outputDirPath)
	os.Mkdir(outputDirPath, 0777)

	for i := smallTsLength; i <= largeTsLength; i++ {

		filePath := "." + string(filepath.Separator) + "output" + string(filepath.Separator) + strconv.Itoa(i) + ".csv"
		f, err := os.Create(filePath)
		check(err)

		ts1, ts2 := data_generator.Generate(i)

		for j := 0; j < numOfIterations; j++ {

			//timeDiffStr := getExecutionTime(dtw_basic.Dtw, ts1, ts2)
			//f.WriteString(timeDiffStr)

			f.WriteString(",")
			timeDiffStr := getExecutionTime(dtw_dynamic.Dtw, ts1, ts2)
			f.WriteString(timeDiffStr)

			f.WriteString(",")
			timeDiffStr = getExecutionTime(dtw_branch_and_bound.Dtw, ts1, ts2)
			f.WriteString(timeDiffStr)
			f.WriteString("\n")
		}

		f.Close()
	}
}

type dtwFunc func(ts1 []float64, ts2 []float64) float64

func getExecutionTime(dtw dtwFunc, ts1 []float64, ts2 []float64) string {
	startTime := time.Now()
	distance := dtw(ts1, ts2)
	endTime := time.Now()
	fmt.Println(distance)
	return strconv.FormatInt(endTime.UnixNano()-startTime.UnixNano(), 10)
}
