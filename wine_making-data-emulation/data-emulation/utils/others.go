package utils

import (
	"log"
	"math"
	"math/rand"

	"github.com/tealeg/xlsx"
)

func InitPH(index int) float64 {
	// 打开一个xlsx文件
	wb, err := xlsx.OpenFile("ph.xlsx")
	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}

	// 遍历文件中的所有工作表
	sheet := wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[index]
	value, _ := row.Cells[1].Float()

	res := math.Floor((value+rand.NormFloat64()*0.08)*100) / 100
	//fmt.Println(res)
	return res
}

func InitAlcohol(index int) float64 {
	// 打开一个xlsx文件
	wb, err := xlsx.OpenFile("alcohol.xlsx")
	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}

	// 遍历文件中的所有工作表
	sheet := wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[index]
	value, _ := row.Cells[1].Float()

	res := math.Floor((value+rand.NormFloat64()*0.08)*100) / 100
	//fmt.Println(res)
	return res
}

func InitTemp(index int) float64 {
	// 打开一个xlsx文件
	wb, err := xlsx.OpenFile("temperature.xlsx")
	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}

	// 遍历文件中的所有工作表
	sheet := wb.Sheets[0]
	// 遍历工作表中的所有行
	row := sheet.Rows[index]
	value, _ := row.Cells[1].Float()

	res := math.Floor((value+rand.NormFloat64()*0.08)*100) / 100
	//fmt.Println(res)
	return res
}
