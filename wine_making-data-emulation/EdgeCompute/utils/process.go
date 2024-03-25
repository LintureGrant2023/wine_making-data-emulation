package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

func ComputeNormalData(data *api.QueryTableResult) SystemHistoryData {

	fmt.Println("开始数据处理...")

	//res := make([]SystemHistoryData, 100)
	var temp SystemHistoryData
	// temp.TemperatureMax = 0
	// temp.TemperatureMin = math.MaxFloat64
	// temp.TemperatureAvg = 0

	//fmt.Println(data.Next())

	for data.Next() {
		//fmt.Println(data.Record().Value())
		value, _ := ConvertStringToSlice(data.Record().Value().(string))
		//fmt.Println("field = ", data.Record().Field(), " time = ", data.Record().Time(), " value = ", value)
		if data.Record().Field() == "Temperature" {
			temp.TemperatureMax = maxFloats(value)
			temp.TemperatureMin = minFloats(value)
			temp.TemperatureAvg = avgFloats(value)
		}
		if data.Record().Field() == "PH" {
			temp.PHMax = maxFloats(value)
			temp.PHMin = minFloats(value)
			temp.PHAvg = avgFloats(value)
		}
		if data.Record().Field() == "CO2" {
			temp.CO2Max = maxFloats(value)
			temp.CO2Min = minFloats(value)
			temp.CO2Avg = avgFloats(value)
		}
		if data.Record().Field() == "O2" {
			temp.O2Max = maxFloats(value)
			temp.O2Min = minFloats(value)
			temp.O2Avg = avgFloats(value)
		}
		if data.Record().Field() == "Alcohol" {
			temp.AlcoholMax = int(maxFloats(value))
			temp.AlcoholMin = int(minFloats(value))
			temp.AlcoholAvg = int(avgFloats(value))
		}
		if data.Record().Field() == "Model" {
			temp.Model = int(value[0])
		}
		//fmt.Println(temp)
		//res = append(res, temp)
	}

	fmt.Println("数据处理完成...")

	return temp
}

func ComputeInitData(data *api.QueryTableResult) SystemHistoryData {
	//res := make([]SystemHistoryData, 100)
	var temp SystemHistoryData
	// temp.TemperatureMax = 0
	// temp.TemperatureMin = math.MaxFloat64
	// temp.TemperatureAvg = 0

	//fmt.Println(data.Next())

	for data.Next() {
		//fmt.Println(data.Record().Value())
		value, _ := ConvertStringToSlice(data.Record().Value().(string))
		//fmt.Println("field = ", data.Record().Field(), " time = ", data.Record().Time(), " value = ", value)
		if data.Record().Field() == "Temperature" {
			a := math.Floor((value[0]+rand.NormFloat64()*0.02)*10) / 10
			b := math.Floor((value[0]+rand.NormFloat64()*0.02)*10) / 10
			temp.TemperatureMax, temp.TemperatureMin = Sort(a, b)
			temp.TemperatureAvg = math.Floor(10*(a+b)/2) / 10
		}
		if data.Record().Field() == "PH" {
			a := math.Floor((value[0]+rand.NormFloat64()*0.02)*10) / 10
			b := math.Floor((value[0]+rand.NormFloat64()*0.02)*10) / 10
			temp.PHMax, temp.PHMin = Sort(a, b)
			temp.PHAvg = math.Floor(10*(a+b)/2) / 10
		}
		if data.Record().Field() == "CO2" {
			a := math.Floor((value[0]+rand.NormFloat64()*0.02)*10) / 10
			b := math.Floor((value[0]+rand.NormFloat64()*0.02)*10) / 10
			temp.CO2Max, temp.CO2Min = Sort(a, b)
			temp.CO2Avg = math.Floor(10*(a+b)/2) / 10
		}
		if data.Record().Field() == "O2" {
			a := math.Floor((value[0]+rand.NormFloat64()*0.02)*10) / 10
			b := math.Floor((value[0]+rand.NormFloat64()*0.02)*10) / 10
			temp.O2Max, temp.O2Min = Sort(a, b)
			temp.O2Avg = math.Floor(10*(a+b)/2) / 10
		}
		if data.Record().Field() == "Alcohol" {
			a := value[0] + rand.NormFloat64()*0.03
			b := value[0] + rand.NormFloat64()*0.03
			if a > b {
				temp.AlcoholMax = int(a)
				temp.AlcoholMin = int(b)
			} else {
				temp.AlcoholMax = int(b)
				temp.AlcoholMin = int(a)
			}
			temp.AlcoholAvg = int((a + b) / 2)
		}
		if data.Record().Field() == "Model" {
			temp.Model = int(value[0])
		}
		//fmt.Println(temp)
		//res = append(res, temp)
	}

	return temp
}

func InitSimulationData(n int, index int) map[string]interface{} {
	//start := time.Now()

	res := make(map[string]interface{}, n)
	ReactorID := make([]int, n)
	EdgeID := make([]int, n)
	Temperature := make([]float64, n)
	CO2 := make([]float64, n)
	O2 := make([]float64, n)
	PH := make([]float64, n)
	Alcohol := make([]int, n)
	for i := 1; i <= n; i++ {
		ReactorID[i-1] = i
		EdgeID[i-1] = i
		Temperature[i-1] = InitTemp(index)
		PH[i-1] = InitPH(index)
		Alcohol[i-1] = InitAlcohol(index)
		CO2[i-1] = InitCO2(index)
		O2[i-1] = InitO2(index)

	}
	//fmt.Println()
	res["ReactorID"] = ReactorID
	res["EdgeID"] = EdgeID
	res["SystemID"] = "1"
	res["Model"] = strconv.Itoa(index + 1) //用来记录发酵时间
	//fmt.Println(res["Model"])
	res["Temperature"] = Temperature
	res["CO2"] = CO2
	res["O2"] = O2
	res["PH"] = PH
	res["Alcohol"] = Alcohol
	res["InsertTime"] = strconv.FormatInt(time.Now().UnixNano(), 10)

	//fmt.Println("InitSimulationData用时: ", time.Since(start))
	return res
}

func GetSimulationData(n int, index int) map[string]interface{} {
	res := make(map[string]interface{}, n)
	ReactorID := make([]int, n)
	EdgeID := make([]int, n)
	Temperature := make([]float64, n)
	CO2 := make([]float64, n)
	O2 := make([]float64, n)
	PH := make([]float64, n)
	Alcohol := make([]int, n)

	for i := 1; i <= n; i++ {
		ReactorID[i-1] = i
		EdgeID[i-1] = i

		Temperature[i-1] = GenerateTemp(index)
		PH[i-1] = GeneratePH(index)
		Alcohol[i-1] = GenerateAlcohol(index)
		CO2[i-1] = GenerateCO2(index)
		O2[i-1] = GenerateO2(index)

		if (i-1)%50 == 0 {
			fmt.Printf("%02d分区,反应器:", 1+(i-1)/50)
		}
		if i%50 != 0 {
			fmt.Printf("%02d,", 1+(i-1)%50)
		} else {
			fmt.Printf("%02d\n", 1+(i-1)%50)
		}

	}
	// fmt.Println("检测到[温度传感器], 采集的数据有: ")
	// PrintFloatSlice(Temperature)
	// fmt.Println("检测到[pH值传感器], 采集的数据有: ")
	// PrintFloatSlice(PH)
	// fmt.Println("检测到[酒精气体浓度传感器], 采集的数据有: ")
	// PrintIntSlice(Alcohol)
	// fmt.Println("检测到[二氧化碳浓度传感器], 采集的数据有: ")
	// PrintFloatSlice(CO2)
	// fmt.Println("检测到[氧气浓度传感器], 采集的数据有: ")
	// PrintFloatSlice(O2)
	//fmt.Println()
	res["ReactorID"] = ReactorID
	res["EdgeID"] = EdgeID
	res["SystemID"] = "1"
	res["Model"] = strconv.Itoa(index + 1) //用来记录发酵时间
	//fmt.Println(res["Model"])
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
		temp, err := strconv.ParseFloat(str, 64)
		floatSlice[i] = math.Floor(temp*10) / 10
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
	return math.Floor(sum/float64(len(slice))*10) / 10
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

func Sort(a, b float64) (float64, float64) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}
