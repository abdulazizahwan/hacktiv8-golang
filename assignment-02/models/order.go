package models

import "time"

type Order struct {
	Order_id      uint      `gorm:"primaryKey" json:"orderID"`
	Customer_name string    `gorm:"notnull" json:"customerName"`
	Items         []Item    `json:"Items" gorm:"foreignKey:Order_id;constraint:OnDelete:CASCADE"`
	Ordered_at    time.Time `json:"orderedAt"`
}
