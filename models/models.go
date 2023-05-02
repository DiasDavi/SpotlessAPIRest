package models

type User struct {
	Id                    int64  `gorm:"primaryKey" json:"id"`
	Name                  string `gorm:"type:varchar(255)" json:"name"`
	Password              string `gorm:"type:varchar(255)" json:"password"`
	City                  string `gorm:"type:varchar(255)" json:"city"`
	State                 string `gorm:"type:varchar(255)" json:"state"`
	Street                string `gorm:"type:varchar(255)" json:"street"`
	HouseNumber           string `gorm:"type:varchar(255)" json:"houseNumber"`
	AdditionalObservation string `gorm:"type:varchar(255)" json:"additionalObservation"`
	Service               bool   `gorm:"type:tinyint(1)" json:"service"`
	Price                 int    `gorm:"type:int(5)" json:"price"`
}
