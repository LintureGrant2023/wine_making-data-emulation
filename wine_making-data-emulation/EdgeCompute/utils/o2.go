package utils

import (
	"math"
	"math/rand"
)

var o2Data []float64

// func InitO2(i int) float64 {
// 	return math.Floor((o2Data[i]+rand.NormFloat64()*0.09)*10) / 10
// }

func o2Simulate1(numPoints, start, end int, amplitude, rateA, rateB, frequency, phaseShift float64) []float64 {
	x := linspace(start, end, numPoints)
	//fmt.Println(x)
	y := make([]float64, numPoints)
	for i := range x {
		xPoint := float64(i + start)
		quadratic := rateA * math.Exp(-rateB*xPoint)
		sine := amplitude * math.Sin(frequency*xPoint+phaseShift)
		noise := rand.NormFloat64() * 0.09
		y[i] = math.Abs(quadratic + sine + noise)
		//fmt.Printf("[%d] = %f\n", i, InitO2(i))
	}
	//fmt.Println("func o2Simulate1: ", y[0])
	return y
}

func o2Simulate2(numPoints, start, end, numPoints1 int, amplitude, rateA, rateB, rateC, frequency, phaseShift float64) []float64 {
	x := linspace(start, end, numPoints)
	y := make([]float64, numPoints)
	for i := range x {
		xPoint := float64(i + start)
		quadratic := 1.5*math.Exp(-rateC*(xPoint-float64(numPoints))) + rateA*math.Exp(-rateB*float64(numPoints1)) - 1.5*math.Exp(-rateC*(float64(start-numPoints)))
		sine := amplitude * math.Sin(frequency*xPoint+phaseShift)
		noise := rand.NormFloat64() * 0.05
		y[i] = math.Abs(quadratic + sine + noise)
	}
	return y
}

func o2Simulate3(numPoints, start, end, numPoints1, numPoints2 int, amplitude, rateA, rateB, rateC, rateD, frequency, phaseShift float64) []float64 {
	x := linspace(start, end, numPoints)
	y := make([]float64, numPoints)
	for i := range x {
		xPoint := float64(i + start)
		//这个quadratic和python程序不一致，所以变化很大
		quadratic := 0.8*math.Exp(-rateD*(xPoint-float64(numPoints))) + 1.5*math.Exp(-rateC*float64(numPoints1)) + rateA*math.Exp(-rateB*float64(numPoints1)) - 1.5*math.Exp(-rateC*float64(numPoints1+1-numPoints2)) - 0.8*math.Exp(-rateD*(float64(start-numPoints)))
		sine := amplitude * math.Sin(frequency*xPoint+phaseShift)
		noise := rand.NormFloat64() * 0.03
		y[i] = math.Abs(quadratic + sine + noise)
	}
	return y
}

func generateO2() {
	// 设置数据点的数据量，总计140个
	var o2_numPoints1 = randInt(50, 70)
	var o2_numPoints2 = randInt(35, 45)
	var o2_numPoints3 = 140 - o2_numPoints1 - o2_numPoints2

	// 变化率
	var o2_a = randFloat(20, 22)
	var o2_b = randFloat(0.05, 0.06)
	var o2_c = randFloat(0.01, 0.015)
	var o2_d = randFloat(0.005, 0.01)

	// 振幅
	var o2_amplitude1 = randFloat(0.5, 0.6)
	var o2_amplitude2 = randFloat(0.1, 0.2)
	var o2_amplitude3 = randFloat(0.1, 0.2)

	// 波动频率
	var o2_frequency = randFloat(0.2, 0.4)

	// 相位便宜
	var o2_phaseShift = 0.0

	//使用三个函数产生数据
	y1 := o2Simulate1(o2_numPoints1, 1, o2_numPoints1, o2_amplitude1, o2_a, o2_b, 50, o2_phaseShift)
	y2 := o2Simulate2(o2_numPoints2, o2_numPoints1+1, o2_numPoints1+o2_numPoints2, o2_numPoints1, o2_amplitude2, o2_a, o2_b, o2_c, 50, o2_phaseShift)
	y3 := o2Simulate3(o2_numPoints3, o2_numPoints1+o2_numPoints2+1, 140, o2_numPoints1, o2_numPoints2, o2_amplitude3, o2_a, o2_b, o2_c, o2_d, o2_frequency, o2_phaseShift)

	// 合并x和y的数据
	//x := append(append(linspace(1, o2_numPoints1, o2_numPoints1), linspace(o2_numPoints1+1, o2_numPoints1+o2_numPoints2, o2_numPoints2)...), linspace(o2_numPoints1+o2_numPoints2+1, 140, o2_numPoints3)...)
	y := append(append(y1, y2...), y3...)

	o2Data = y
	//fmt.Println(o2Data)
	// fmt.Println("x:", x)
	//print(o2Data)
}
