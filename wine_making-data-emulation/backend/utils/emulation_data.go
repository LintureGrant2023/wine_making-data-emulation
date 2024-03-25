package utils

import (
	"crypto/rand"
	"math"
	"math/big"
)

type SensorData struct {
	Temp    float64 `json:"temp"`
	CO2     float64 `json:"co2"`
	O2      float64 `json:"o2"`
	PH      float64 `json:"ph"`
	Alcohol float64 `json:"alcohol"`
}

type SensorDataArray struct {
	Temps    []float64 `json:"temps"`
	CO2s     []float64 `json:"co2s"`
	O2s      []float64 `json:"o2s"`
	PHs      []float64 `json:"phs"`
	Alcohols []float64 `json:"alcohols"`
}

// 温度
type TempData struct {
	Temps    []float64 `json:"temps"`
	TempsMin []float64 `json:"tempsMin"`
	TempsMax []float64 `json:"tempsMax"`
}

// ph
type PhData struct {
	Phs    []float64 `json:"phs"`
	PhsMin []float64 `json:"phsMin"`
	PhsMax []float64 `json:"phsMax"`
}

// alcohol
type AlcoholData struct {
	Alcohols    []float64 `json:"alcohols"`
	AlcoholsMin []float64 `json:"alcoholsMin"`
	AlcoholsMax []float64 `json:"alcoholsMax"`
}

// o2
type O2 struct {
	O2    float64 `json:"o2"`
	O2Min float64 `json:"o2Min"`
	O2Max float64 `json:"o2Max"`
	//Time   int
}

// co2
type CO2 struct {
	CO2    float64 `json:"co2"`
	CO2Min float64 `json:"co2Min"`
	CO2Max float64 `json:"co2Max"`
	//Time    i
}

// 温度
type Temp struct {
	Temp    float64 `json:"temp"`
	TempMin float64 `json:"tempMin"`
	TempMax float64 `json:"tempMax"`
}

// ph
type Ph struct {
	Ph    float64 `json:"ph"`
	PhMin float64 `json:"phMin"`
	PhMax float64 `json:"phMax"`
}

// alcohol
type Alcohol struct {
	Alcohol    float64 `json:"alcohol"`
	AlcoholMin float64 `json:"alcoholMin"`
	AlcoholMax float64 `json:"alcoholMax"`
}

type SensorStatus struct {
	Temp    Temp    `json:"temp"`
	Ph      Ph      `json:"ph"`
	Alcohol Alcohol `json:"alcohol"`
	CO2     CO2     `json:"co2"`
	O2      O2      `json:"o2"`
}

// abnormData
type AbnormData struct {
}

// reactor status info
type ReactorStatus struct {
	FermentTime float32
	TempOutdoor float32
	TempEnv     float32
	HumidEnv    float32
	//SensorData  SensorData
}

type PerdictData struct {
	Temps []float64 `json:"temps"`
	Phs   []float64 `json:"phs"`
}

type UpdatedData struct {
	Temp float64 `json:"temp"`
	Ph   float64 `json:"ph"`
}

// 获取温度
func GetTemp() Temp {
	//res := make([]TempData, count)ar res TempData
	var res Temp
	res.Temp = GenerateRandomFloat64Range(0, 35)
	res.TempMin = GenerateRandomFloat64Range(0, 35)
	res.TempMax = GenerateRandomFloat64Range(0, 35)
	// res.Time = time.Now().Seco()
	return res
}

func GetTemps(count int) TempData {
	temps := make([]float64, count)
	tempsMin := make([]float64, count)
	tempsMax := make([]float64, count)
	for i := 0; i < count; i++ {
		temps[i] = GenerateRandomFloat64Range(0, 35)
		tempsMin[i] = GenerateRandomFloat64Range(0, 35)
		tempsMax[i] = GenerateRandomFloat64Range(0, 35)
	}
	res := TempData{
		Temps:    temps,
		TempsMin: tempsMin,
		TempsMax: tempsMax,
	}
	return res
}

// 获取ph
func GetPh() Ph {
	//res := make([]PhData, count)
	var res Ph
	res.Ph = GenerateRandomFloat64Range(0, 14)
	res.PhMin = GenerateRandomFloat64Range(0, 14)
	res.PhMax = GenerateRandomFloat64Range(0, 14)
	// res.Time = time.Now().Second()
	return res
}

func GetPhs(count int) PhData {
	phs := make([]float64, count)
	phMin := make([]float64, count)
	phMax := make([]float64, count)
	for i := 0; i < count; i++ {
		phs[i] = GenerateRandomFloat64Range(0, 14)
		phMin[i] = GenerateRandomFloat64Range(0, 14)
		phMax[i] = GenerateRandomFloat64Range(0, 14)
	}
	res := PhData{
		Phs:    phs,
		PhsMin: phMin,
		PhsMax: phMax,
	}
	return res
}

// 获取酒精浓度
func GetAlcohol() Alcohol {
	//res := make([]AlcoholData, count)
	var res Alcohol
	res.Alcohol = GenerateRandomFloat64Range(0, 100)
	res.AlcoholMin = GenerateRandomFloat64Range(0, 100)
	res.AlcoholMax = GenerateRandomFloat64Range(0, 100)
	// res.Time = time.Now().Second()
	return res
}

func GetAlcohols(count int) AlcoholData {
	alcohols := make([]float64, count)
	alcoholMin := make([]float64, count)
	alcoholMax := make([]float64, count)
	for i := 0; i < count; i++ {
		alcohols[i] = GenerateRandomFloat64Range(0, 100)
		alcoholMin[i] = GenerateRandomFloat64Range(0, 100)
		alcoholMax[i] = GenerateRandomFloat64Range(0, 100)
	}
	res := AlcoholData{
		Alcohols:    alcohols,
		AlcoholsMin: alcoholMin,
		AlcoholsMax: alcoholMax,
	}
	return res
}

// 获取氧气浓度
func GetO2() O2 {
	var res O2
	res.O2 = GenerateRandomFloat64Range(0, 100)
	res.O2Min = GenerateRandomFloat64Range(0, 100)
	res.O2Max = GenerateRandomFloat64Range(0, 100)
	//res.Time = time.Now().Second()
	return res
}

func GetO2s(count int) []O2 {
	res := make([]O2, count)
	//var res O2Datas
	for i := 0; i < count; i++ {
		res[i].O2 = GenerateRandomFloat64Range(0, 100)
		res[i].O2Min = GenerateRandomFloat64Range(0, 100)
		res[i].O2Max = GenerateRandomFloat64Range(0, 100)
		//res[i].Time = i
	}
	return res
}

// 获取二氧化碳浓度
func GetCO2() CO2 {
	var res CO2
	res.CO2 = GenerateRandomFloat64Range(0, 100)
	res.CO2Min = GenerateRandomFloat64Range(0, 100)
	res.CO2Max = GenerateRandomFloat64Range(0, 100)
	return res
}

// 获取二氧化碳浓度
func GetCO2s(count int) []CO2 {
	res := make([]CO2, count)
	//var res CO2Datas
	for i := 0; i < 12; i++ {
		res[i].CO2 = GenerateRandomFloat64Range(0, 100)
		res[i].CO2Min = GenerateRandomFloat64Range(0, 100)
		res[i].CO2Max = GenerateRandomFloat64Range(0, 100)
		//res[i].Time = i
	}
	return res
}

// 获取传感器数据
func GetSensorData() SensorData {
	sensor_data := SensorData{}
	sensor_data.Temp = GenerateRandomFloat64Range(0, 35)
	sensor_data.CO2 = GenerateRandomFloat64Range(0, 100)
	sensor_data.O2 = GenerateRandomFloat64Range(0, 100)
	sensor_data.PH = GenerateRandomFloat64Range(0, 14)
	sensor_data.Alcohol = GenerateRandomFloat64Range(0, 100)
	// current := time.Now()
	// fmt.Println("year = ", current.Year(), ", hour = ", current.Hour())
	return sensor_data
}

// 获取预测的温度
func GetPerdictedTemps(count int) TempData {
	temps := make([]float64, count)
	tempsMin := make([]float64, count)
	tempsMax := make([]float64, count)
	for i := 0; i < count; i++ {
		temps[i] = GenerateRandomFloat64Range(0, 35)
		tempsMin[i] = GenerateRandomFloat64Range(0, 35)
		tempsMax[i] = GenerateRandomFloat64Range(0, 35)
	}
	res := TempData{
		Temps:    temps,
		TempsMin: tempsMin,
		TempsMax: tempsMax,
	}
	return res
}

// 获取预测的ph
func GetPerdictedPhs(count int) PhData {
	phs := make([]float64, count)
	phMin := make([]float64, count)
	phMax := make([]float64, count)
	for i := 0; i < count; i++ {
		phs[i] = GenerateRandomFloat64Range(0, 14)
		phMin[i] = GenerateRandomFloat64Range(0, 14)
		phMax[i] = GenerateRandomFloat64Range(0, 14)
	}
	res := PhData{
		Phs:    phs,
		PhsMin: phMin,
		PhsMax: phMax,
	}
	return res
}

// 获取异常信息
func GetAbnorms() AbnormData {
	//to do
	var res AbnormData
	return res
}

// #web2界面接口
// 获取反应器状态信息
func GetReactorStatus() ReactorStatus {
	var res ReactorStatus
	//res.SensorData = GetSensorData()
	res.FermentTime = float32(GenerateRandomFloat64Range(0, 120))
	res.TempOutdoor = float32(GenerateRandomFloat64Range(0, 38))
	res.TempEnv = float32(GenerateRandomFloat64Range(10, 35))
	res.HumidEnv = float32(GenerateRandomFloat64Range(0.3, 0.8))
	return res
}

// 获取当前的传感器数据
func GetCurrentSensor() SensorDataArray {
	//var res SensorDataArray
	res := SensorDataArray{
		Alcohols: make([]float64, 12),
		CO2s:     make([]float64, 12),
		O2s:      make([]float64, 12),
		PHs:      make([]float64, 12),
		Temps:    make([]float64, 12),
	}
	for i := 0; i < 12; i++ {
		res.Alcohols[i] = GenerateRandomFloat64Range(0, 100)
		res.CO2s[i] = GenerateRandomFloat64Range(0, 100)
		res.O2s[i] = GenerateRandomFloat64Range(0, 100)
		res.PHs[i] = GenerateRandomFloat64Range(0, 14)
		res.Temps[i] = GenerateRandomFloat64Range(0, 35)
	}
	return res
}

// 获取历史的传感器数据
func GetHistorySensor() SensorDataArray {
	res := SensorDataArray{
		Alcohols: make([]float64, 12),
		CO2s:     make([]float64, 12),
		O2s:      make([]float64, 12),
		PHs:      make([]float64, 12),
		Temps:    make([]float64, 12),
	}
	for i := 0; i < 12; i++ {
		res.Alcohols[i] = GenerateRandomFloat64Range(0, 100)
		res.CO2s[i] = GenerateRandomFloat64Range(0, 100)
		res.O2s[i] = GenerateRandomFloat64Range(0, 100)
		res.PHs[i] = GenerateRandomFloat64Range(0, 14)
		res.Temps[i] = GenerateRandomFloat64Range(0, 35)
	}
	return res
}

func GetSensorStatus() SensorStatus {
	var res SensorStatus
	res.Alcohol = GetAlcohol()
	res.CO2 = GetCO2()
	res.O2 = GetO2()
	res.Ph = GetPh()
	res.Temp = GetTemp()
	return res
}

func GetPerdictedSensor(count int) PerdictData {
	res := PerdictData{
		Temps: make([]float64, count),
		Phs:   make([]float64, count),
	}
	for i := 0; i < count; i++ {
		res.Phs[i] = GenerateRandomFloat64Range(0, 14)
		res.Temps[i] = GenerateRandomFloat64Range(0, 35)
	}
	return res
}

func GetUpdatedSensor(count int) UpdatedData {
	var res UpdatedData
	res.Ph = GenerateRandomFloat64Range(0, 14)
	res.Temp = GenerateRandomFloat64Range(0, 35)

	return res
}

// 新增
func GetCurrentAll() SensorStatus {
	var res SensorStatus
	res.Alcohol = GetAlcohol()
	res.CO2 = GetCO2()
	res.O2 = GetO2()
	res.Ph = GetPh()
	res.Temp = GetTemp()
	return res
}

func GetEnv() ReactorStatus {
	var res ReactorStatus
	//res.SensorData = GetSensorData()
	res.FermentTime = float32(GenerateRandomFloat64Range(0, 120))
	res.TempOutdoor = float32(GenerateRandomFloat64Range(0, 38))
	res.TempEnv = float32(GenerateRandomFloat64Range(10, 35))
	res.HumidEnv = float32(GenerateRandomFloat64Range(0.3, 0.8))
	return res
}

func GetPerdictUpdated() UpdatedData {
	var res UpdatedData
	res.Ph = GenerateRandomFloat64Range(0, 14)
	res.Temp = GenerateRandomFloat64Range(0, 35)

	return res
}

// 生成随机数
func GenerateRandomFloat64Range(min, max float64) float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}

	//返回2位小数的float64
	return math.Floor((min+(max-min)*float64(n.Int64())/float64(math.MaxInt64))*100) / 100
}
