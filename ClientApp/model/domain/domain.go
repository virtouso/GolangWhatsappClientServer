package domain

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	AccountId string
}
