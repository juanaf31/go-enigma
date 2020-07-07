package utils

import (
	"liveCodeAPI/api-master/master/models"
	"strconv"
)

func ValidationGoods(data *models.Goods) bool {
	cap, _ := strconv.Atoi(data.Capacity)
	qty, _ := strconv.Atoi(data.Quantity)
	totCap := cap * qty

	if data.ID == "" || data.Goods_name == "" || data.Quantity == "" {
		return false
	} else if totCap > 25000 || totCap == 0 {
		return false
	} else if data.WarehouseID != "WH0001" && data.WarehouseID != "WH0002" {
		return false
	} else {
		return true
	}
}
