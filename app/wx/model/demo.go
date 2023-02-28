package model

type Demo struct {
	ID int `gorm:"id" json:"id"`
}

func (m *Demo) TableName() string {
	return "demo"
}
