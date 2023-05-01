package models

type User struct {
	Id       int64   `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"type:varchar(255)" json:"name"`
	Password string  `gorm:"type:varchar(255)" json:"password"`
	Address  Address `gorm:"foreignKey:AdressID" json:"address"`
	AdressID int64   `json:"-"`
}

type ServiceProvider struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	ServiceType int    `gorm:"type:int(5)" json:"password"`
	Price       int    `gorm:"type:int(5)" json:"password"`
}

type Address struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	City  string `gorm:"type:varchar(255)" json:"city"`
	State string `gorm:"type:varchar(255)" json:"state"`
}