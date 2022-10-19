package controllers

import (
	"assignment-02/database"
	"assignment-02/models"
	"fmt"
)

func QueryCreate(orderInput models.Order) models.Order {
	db := database.GetDB()

	newOrder := orderInput

	dberr := db.Debug().Create(&newOrder).Error

	if dberr != nil {
		panic(dberr)

	}

	return newOrder
}

func QueryGetAll() []models.Order {
	db := database.GetDB()

	var orders []models.Order

	dberr := db.Preload("Items").Find(&orders).Error

	if dberr != nil {
		panic(dberr)
	}

	return orders
}

func QueryUpdateByID(orderInput models.Order, id uint) models.Order {
	db := database.GetDB()

	updatedOrder := orderInput
	var err error

	for i := range updatedOrder.Items {
		err = db.Model(&updatedOrder.Items[i]).Where("Item_id=?", updatedOrder.Items[i].Item_id).Updates(&updatedOrder.Items[i]).Error
		if err != nil {
			panic(err)
		}
	}

	var updatedOnlyOrder models.Order
	updatedOnlyOrder.Customer_name = updatedOrder.Customer_name
	updatedOnlyOrder.Ordered_at = updatedOrder.Ordered_at

	err = db.Model(&updatedOnlyOrder).Where("Order_id=?", id).Updates(&updatedOnlyOrder).Error

	if err != nil {
		panic(err)
	}

	err = db.Preload("Items").Where("Order_id=?", id).Find(&updatedOrder).Error

	if err != nil {
		panic(err)
	}

	fmt.Println("Data updated")
	return updatedOrder
}

func QueryDeleteByID(id uint) {
	db := database.GetDB()

	dberr := db.Where("Order_id=?", id).Delete(&models.Item{}).Error

	if dberr != nil {
		panic(dberr)
	}

	dberr = db.Delete(&models.Order{}, id).Error

	if dberr != nil {
		panic(dberr)
	}

	fmt.Println("Data deleted")
}
