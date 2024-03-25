package config

type Sensor struct {
	TempUpperThreshold   string `form:"temp_upper_threshold" json:"temp_upper_threshold" yaml:"temp_upper_threshold"`
	TempLowerThreshold   string `form:"temp_lower_threshold" json:"temp_lower_threshold" yaml:"temp_lower_threshold"`
	O2UpperThreshold     string `form:"o2_upper_threshold" json:"o2_upper_threshold" yaml:"o2_upper_threshold"`
	O2LowerThreshold     string `form:"o2_lower_threshold" json:"o2_lower_threshold" yaml:"o2_lower_threshold"`
	CO2UpperThreshold    string `form:"co2_upper_threshold" json:"co2_upper_threshold" yaml:"co2_upper_threshold"`
	CO2LowerThreshold    string `form:"co2_lower_threshold" json:"co2_lower_threshold" yaml:"co2_lower_threshold"`
	PHUpperThreshold     string `form:"ph_upper_threshold" json:"ph_upper_threshold" yaml:"ph_upper_threshold"`
	PHLowerThreshold     string `form:"ph_lower_threshold" json:"ph_lower_threshold" yaml:"ph_lower_threshold"`
	StarchUpperThreshold string `form:"starch_upper_threshold" json:"starch_upper_threshold" yaml:"starch_upper_threshold"`
	StarchLowerThreshold string `form:"starch_lower_threshold" json:"starch_lower_threshold" yaml:"starch_lower_threshold"`
}
