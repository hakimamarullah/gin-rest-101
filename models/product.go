package models

type Product struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(300)" json:"product_name" binding:"required"`
	Description string `gorm:"type:varchar(500)" json:"description" binding:"required"`

}