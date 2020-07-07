package models

type Goods struct {
	ID          string `json:"goods_id"`
	Goods_name  string `json:"goods_name"`
	Quantity    string `json:"goods_quantity"`
	Capacity    string `json:"goods_capacity"`
	WarehouseID string `json:"warehouse_id"`
}
