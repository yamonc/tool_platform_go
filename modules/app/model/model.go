package model

import "gorm.io/gorm"

type TestModel struct {
	Name string `json:"name"`
}
type Car struct {
	gorm.Model
	Uuid    string `gorm:"column:uuid"`
	Name    string `gorm:"column:name"`
	DailyKm string `gorm:"column:daily_km"`
	NowKm   string `gorm:"column:now_km"`
	IsAlarm bool   `gorm:"column:is_alarm"`
	Remark  string `gorm:"column:remark"`
}

type APICar struct {
	gorm.Model
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	DailyKm    string `json:"dailyKm"`
	NowKm      string `json:"now_km"`
	IsAlarm    bool   `json:"isAlarm"`
	Remark     string `json:"remark"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

type CarRecord struct {
	Uuid  string `json:"uuid"`
	CarId string `json:"carId"`
	Item  string `json:"item"`
	Km    string `json:"km"`
	Time  string `json:"time"`
}

type CarMaintain struct {
	Uuid       string `json:"uuid"`
	CarId      string `json:"carId"`
	CarName    string `json:"carName"`
	LastKm     string `json:"lastKm"`
	LastTime   string `json:"lastTime"`
	NowKm      string `json:"nowKm"`
	Remark     string `json:"remark"`
	CreateTime string `json:"createTime"`
}
