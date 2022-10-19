package models

type Item struct {
	Item_id     uint   `gorm:"primaryKey" json:"lineItemID"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_id    uint   `json:"orderID"`
}
