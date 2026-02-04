package model

import "time"

type BaseModel struct {
	CreatedDate  time.Time `gorm:"column:created_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_date"`
	ModifiedDate time.Time `gorm:"column:modified_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"modified_date"`
}

type Region struct {
	RegionID   uint   `gorm:"primaryKey;column:region_id"  json:"region_id"`
	RegionName string `gorm:"column:region_name;type:varchar(25)" json:"region_name"`
	BaseModel         //embendded
}

func (Region) TableName() string { return "hr.regions" }

type Country struct {
	CountryID   string `gorm:"primaryKey;column:country_id;type:char(2)" json:"country_id"`
	CountryName string `gorm:"column:country_name;type:varchar(40)" json:"country_name"`
	RegionID    uint   `gorm:"column:region_id" json:"region_id"`
	BaseModel          //embendded
}

func (Country) TableName() string { return "hr.countries" }
