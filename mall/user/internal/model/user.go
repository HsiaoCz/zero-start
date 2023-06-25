package model

type User struct {
	Id     int64  `gorm:"clomun:id;type:int(11);" json:"id"`
	Name   string `gorm:"clomun:name;type:varchar(40);" json:"name"`
	Gender string `gorm:"clomun:gender;type:varchar(10);" json:"gender"`
}

func (u User) TableName() string {
	return "user"
}
