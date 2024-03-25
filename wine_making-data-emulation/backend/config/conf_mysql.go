package config

import "strconv"

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Others   string `yaml:"others"`
}

func (mysql MySQL) GetDB() string {
	return mysql.User + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + strconv.Itoa(mysql.Port) + ")/" + mysql.DBName + "?" + mysql.Others
}
