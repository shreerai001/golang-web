package models

type BusinessUser struct {
	BaseModel
	FullName string `gorm:"size:255;not null;unique" json:"full_name"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
}

func (BusinessUser) TableName() string {
	return "BusinessUser"
}
