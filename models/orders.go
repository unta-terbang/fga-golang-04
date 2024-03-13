package models

import(
	"time"
)

type Order struct{

	ID uint `json:"id" gorm:"primaryKey"`
	Costumer_name string `json:"costumer_name" gorm:"not null;tyype:varchar(50)"`
	Ordered_at time.Time  `json:"ordered_at"`
	Items []Item `json:"items" gorm:"foreignKey:OrderID"`

}

var OrderDatas = []Order{}