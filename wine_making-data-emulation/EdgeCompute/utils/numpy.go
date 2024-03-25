package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().UnixNano())
}

func randInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func linspace(start, end, num int) []float64 {
	res := make([]float64, num)
	if num == 1 {
		res[0] = float64(start)
		return res
	}
	delta := float64(end-start) / float64(num-1)
	for i := range res {
		res[i] = float64(start) + delta*float64(i)
	}
	return res
}

func print(x []float64) {
	var count = 1
	for _, v := range x {
		fmt.Printf("%f,", v)
		if count%10 == 0 {
			fmt.Println()
		}
		count++
	}
}
