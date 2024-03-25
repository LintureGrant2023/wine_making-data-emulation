package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func GetSimulationData(n int, index int) map[string]interface{} {
	res := make(map[string]interface{}, n)
	ReactorID := make([]int, n)
	EdgeID := make([]int, n)
	Temperature := make([]float64, n)
	CO2 := make([]float64, n)
	O2 := make([]float64, n)
	PH := make([]float64, n)
	Alcohol := make([]float64, n)
	for i := 1; i <= n; i++ {
		ReactorID[i-1] = i
		EdgeID[i-1] = i

		Temperature[i-1] = InitTemp(index)
		CO2[i-1] = InitCO2(index)
		O2[i-1] = InitO2(index)
		PH[i-1] = InitPH(index)
		Alcohol[i-1] = InitAlcohol(index)
	}
	//fmt.Println()
	res["ReactorID"] = ReactorID
	res["EdgeID"] = EdgeID
	res["SystemID"] = "1"
	res["Model"] = "1"
	res["Temperature"] = Temperature
	res["CO2"] = CO2
	res["O2"] = O2
	res["PH"] = PH
	res["Alcohol"] = Alcohol
	res["InsertTime"] = strconv.FormatInt(time.Now().UnixNano(), 10)
	return res
}

func ConvertStringToSlice(stringValue string) ([]float64, error) {
	// 去掉字符串两端的方括号
	trimmedStringValue := strings.Trim(stringValue, "[]")

	// 分割字符串得到单独的数字字符串
	stringElements := strings.Split(trimmedStringValue, " ")

	// 创建一个浮点数切片来存储转换后的值
	floatSlice := make([]float64, len(stringElements))

	// 将字符串转换为浮点数
	for i, str := range stringElements {
		// 去除可能的空格
		str = strings.TrimSpace(str)

		// 将字符串转换为浮点数
		var err error
		floatSlice[i], err = strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("convertStringToSlice error, string = ", stringElements)
			return nil, fmt.Errorf("convertStringToSlice: %v", err)
		}
	}
	//fmt.Println(floatSlice)
	// 返回浮点数切片和nil作为错误（表示没有错误）
	return floatSlice, nil
}

// sumFloats 返回一个 float64 切片的平均值
func avgFloats(slice []float64) (sum float64) {
	for _, v := range slice {
		sum += v
	}
	return math.Floor(sum/float64(len(slice))*100) / 100
}

// maxFloats 返回一个 float64 切片的最大值
func maxFloats(slice []float64) (max float64) {
	max = math.Inf(-1)
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return
}

// minFloats 返回一个 float64 切片的最小值
func minFloats(slice []float64) (min float64) {
	min = math.Inf(1)
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return
}
