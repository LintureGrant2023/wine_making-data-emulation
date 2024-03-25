package utils

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/tealeg/xlsx"
)

// 等距采样
const Delta = 72

var Ph_wb *xlsx.File
var Alcohol_wb *xlsx.File
var Temp_wb *xlsx.File
var CO2_wb *xlsx.File
var O2_wb *xlsx.File

func OpenAllFiles() {
	//start := time.Now()

	// 打开一个xlsx文件
	var err error
	Ph_wb, err = xlsx.OpenFile("ph_5min.xlsx")
	if err != nil {
		log.Fatalf("Cannot open ph_5min.xlsx: %v", err)
	}

	Alcohol_wb, err = xlsx.OpenFile("alcohol_5min.xlsx")
	if err != nil {
		log.Fatalf("Cannot open alcohol_5min.xlsx: %v", err)
	}

	Temp_wb, err = xlsx.OpenFile("temp_5min.xlsx")
	if err != nil {
		log.Fatalf("Cannot open temp_5min.xlsx: %v", err)
	}

	CO2_wb, err = xlsx.OpenFile("co2_5min.xlsx")
	if err != nil {
		log.Fatalf("Cannot open co2_5min.xlsx: %v", err)
	}

	O2_wb, err = xlsx.OpenFile("o2_5min.xlsx")
	if err != nil {
		log.Fatalf("Cannot open o2_5min.xlsx: %v", err)
	}
	//fmt.Println("OpenAllFiles用时: ", time.Since(start))
}

func InitPH(index int) float64 {

	// 遍历文件中的所有工作表
	sheet := Ph_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[index*Delta]
	value, _ := row.Cells[0].Float()

	res := math.Floor((value+rand.NormFloat64()*0.02)*10) / 10
	//res := math.Floor(value*10) / 10
	//fmt.Println(res)
	return res
}

func InitAlcohol(index int) int {

	// 遍历文件中的所有工作表
	sheet := Alcohol_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[index*Delta]
	value, _ := row.Cells[0].Float()

	res := int(math.Floor(value + rand.NormFloat64()*0.03))
	//fmt.Println(res)
	return res
}

func InitTemp(index int) float64 {

	// 遍历文件中的所有工作表
	sheet := Temp_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[index*Delta]
	value, _ := row.Cells[0].Float()

	res := math.Floor((value+rand.NormFloat64()*0.02)*10) / 10
	//res := math.Floor(value*10) / 10
	//fmt.Println(res)
	return res
}

func InitCO2(index int) float64 {

	// 遍历文件中的所有工作表
	sheet := CO2_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[index*Delta]
	value, _ := row.Cells[0].Float()

	res := math.Floor((value+rand.NormFloat64()*0.03)*10) / 10
	//res := math.Floor(value*10) / 10
	//fmt.Println(res)
	return res
}

func InitO2(index int) float64 {

	// 遍历文件中的所有工作表
	sheet := O2_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[index*Delta]
	value, _ := row.Cells[0].Float()

	res := math.Floor((value+rand.NormFloat64()*0.02)*10) / 10
	//res := math.Floor(value*10) / 10
	//fmt.Println(res)
	return res
}

func GeneratePH(index int) float64 {
	// 遍历文件中的所有工作表
	sheet := Ph_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[1420+index]
	value, _ := row.Cells[0].Float()

	res := math.Floor((value+rand.NormFloat64()*0.02)*10) / 10
	//res := math.Floor(value*10) / 10
	//fmt.Println(res)
	return res
}

func GenerateAlcohol(index int) int {
	// 遍历文件中的所有工作表
	sheet := Alcohol_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[1420+index]
	value, _ := row.Cells[0].Float()

	res := int(math.Floor(value + rand.NormFloat64()*0.03))
	//fmt.Println(res)
	return res
}

func GenerateTemp(index int) float64 {
	// 遍历文件中的所有工作表
	sheet := Temp_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[1420+index]
	value, _ := row.Cells[0].Float()

	res := math.Floor((value+rand.NormFloat64()*0.02)*10) / 10
	//res := math.Floor(value*10) / 10
	//fmt.Println(res)
	return res
}

func GenerateCO2(index int) float64 {
	// 遍历文件中的所有工作表
	sheet := CO2_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[1420+index]
	value, _ := row.Cells[0].Float()

	res := math.Floor((value+rand.NormFloat64()*0.03)*10) / 10
	//res := math.Floor(value*10) / 10
	//fmt.Println(res)
	return res
}

func GenerateO2(index int) float64 {
	// 遍历文件中的所有工作表
	sheet := O2_wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[1420+index]
	value, _ := row.Cells[0].Float()

	res := math.Floor((value+rand.NormFloat64()*0.02)*10) / 10
	//res := math.Floor(value*10) / 10
	//fmt.Println(res)
	return res
}

func GenerateAbnormal() {
	count := randInt(1, 2)
	for i := 0; i < count; i++ {
		data := SystemHistoryData{
			Model:      0,
			CreateTime: time.Now(),
			//温度过高
			TemperatureMax: math.Floor(randFloat(35, 38)*10) / 10,
			TemperatureMin: math.Floor(randFloat(20, 35)*10) / 10,
			TemperatureAvg: math.Floor(randFloat(20, 35)*10) / 10,
			//PH过低
			PHMax: math.Floor(randFloat(3, 7)*10) / 10,
			PHMin: math.Floor(randFloat(1, 3)*10) / 10,
			PHAvg: math.Floor(randFloat(3, 7)*10) / 10,
			//o2过高
			O2Max: math.Floor(randFloat(21, 25)*10) / 10,
			O2Min: math.Floor(randFloat(0.1, 21)*10) / 10,
			O2Avg: math.Floor(randFloat(0.1, 21)*10) / 10,
			//co2过高
			CO2Max: math.Floor(randFloat(25, 28)*10) / 10,
			CO2Min: math.Floor(randFloat(0.1, 25)*10) / 10,
			CO2Avg: math.Floor(randFloat(0.1, 25)*10) / 10,
			//alcohol过低
			AlcoholMax: randInt(200, 2200),
			AlcoholMin: randInt(0, 200),
			AlcoholAvg: randInt(200, 1800),
		}
		WriteIntoMysql(data)
	}
}

func PrintFloatSlice(slice []float64) {
	var count = 0
	for i, v := range slice {
		count++
		if i == 0 || i%20 != 0 {
			fmt.Printf("%4.1f, ", v)
		} else {
			fmt.Printf("\n%4.1f, ", v)
		}
	}
	fmt.Println("..., 共计1000个数据")
}

func PrintIntSlice(slice []int) {
	var count = 0
	for i, v := range slice {
		count++
		if i == 0 || i%20 != 0 {
			fmt.Printf("%4d, ", v)
		} else {
			fmt.Printf("\n%4d, ", v)
		}
	}
	fmt.Println("..., 共计1000个数据")
}
