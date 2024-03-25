package utils

import (
	"fmt"
	"math"
	"math/rand"
)

// 设置数据点的数据量，总计140个
var co2_numPoints1 = randInt(40, 45)
var co2_numPoints2 = randInt(35, 40)
var co2_numPoints3 = 140 - co2_numPoints1 - co2_numPoints2

//initialHumidity := randInt(45, 55)

// 变化率
var co2_a = randFloat(0.009, 0.0095)
var co2_b = randFloat(0.005, 0.0055)
var co2_c = randFloat(0.003, 0.0035)

var co2_amplitude1 = randFloat(0.2, 0.3)
var co2_amplitude2 = randFloat(0.1, 0.15)
var co2_amplitude3 = randFloat(0.3, 0.4)

// frequency := randFloat(0.2, 0.4)
var co2_phaseShift = 0.0 // 或者其他你可能设置的非零值

func InitCO2(i int) float64 {
	xPoint := float64(i + 1)
	quadratic := -co2_a*math.Pow(xPoint-float64(co2_numPoints1), 2) + co2_a*math.Pow(float64(co2_numPoints1), 2)
	sine := co2_amplitude1 * math.Sin(50*xPoint+co2_phaseShift)
	noise := rand.NormFloat64() * 0.07
	return math.Floor((quadratic+sine+noise)*100) / 100
}

func co2Simulate1(numPoints, start, end int, amplitude, rateA, frequency, phaseShift float64) []float64 {
	x := linspace(start, end, numPoints)
	y := make([]float64, numPoints)
	for i := range x {
		xPoint := float64(i + start)
		quadratic := -rateA*math.Pow(xPoint-float64(end), 2) + rateA*math.Pow(float64(end), 2)
		sine := amplitude * math.Sin(frequency*xPoint+phaseShift)
		noise := rand.NormFloat64() * 0.07
		y[i] = quadratic + sine + noise
	}
	return y
}

func co2Simulate2(numPoints, start, end int, amplitude, rateA, rateB, frequency, phaseShift float64) []float64 {
	x := linspace(start, end, numPoints)
	y := make([]float64, numPoints)
	for i := range x {
		xPoint := float64(i + start)
		quadratic := -rateB*math.Pow(xPoint-float64(end)+5, 2) + rateA*math.Pow(float64(start)-1, 2) + rateB*math.Pow(5+float64(start)-float64(end), 2)
		sine := amplitude * math.Sin(frequency*xPoint+phaseShift)
		noise := rand.NormFloat64() * 0.05
		y[i] = quadratic + sine + noise
	}
	return y
}

func co2Simulate3(numPoints, start, end, numPoints1, numPoints2 int, amplitude, rateA, rateB, rateC, frequency, phaseShift float64) []float64 {
	x := linspace(start, end, numPoints)
	y := make([]float64, numPoints)
	for i := range x {
		xPoint := float64(i + start)
		//这个quadratic和python程序不一致，所以变化很大
		quadratic := rateC*math.Pow(xPoint-float64(end), 2) + rateA*math.Pow(float64(numPoints1), 2) + rateB*math.Pow(float64(6-numPoints2), 2) - rateC*math.Pow(float64(start)-float64(end), 2)
		sine := amplitude * math.Sin(frequency*xPoint+phaseShift)
		noise := rand.NormFloat64() * 0.03
		y[i] = quadratic + sine + noise
	}
	return y
}

func GenerateCO2() {
	//使用三个函数产生数据
	y1 := co2Simulate1(co2_numPoints1, 1, co2_numPoints1, co2_amplitude1, co2_a, 50, co2_phaseShift)
	y2 := co2Simulate2(co2_numPoints2, co2_numPoints1+1, co2_numPoints1+co2_numPoints2, co2_amplitude2, co2_a, co2_b, 50, co2_phaseShift)
	y3 := co2Simulate3(co2_numPoints3, co2_numPoints1+co2_numPoints2+1, 140, co2_numPoints1, co2_numPoints2, co2_amplitude3, co2_a, co2_b, co2_c, 0.2, co2_phaseShift)

	// 合并x和y的数据
	x := append(append(linspace(1, co2_numPoints1, co2_numPoints1), linspace(co2_numPoints1+1, co2_numPoints1+co2_numPoints2, co2_numPoints2)...), linspace(co2_numPoints1+co2_numPoints2+1, 140, co2_numPoints3)...)
	y := append(append(y1, y2...), y3...)

	fmt.Println("x:", x)
	print(y)
}
