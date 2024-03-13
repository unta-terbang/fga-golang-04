package models

type Item struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Code        string `json:"code" gorm:"not null;type:varchar(10)"`
	Description string `json:"description" gorm:"not null;type:varchar(50)"`
	Quantity    uint   `json:"quantity"` // Corrected "quatity" to "quantity"
	OrderID     uint   `json:"order_id"`
}

var ItemDatas = []Item{}
