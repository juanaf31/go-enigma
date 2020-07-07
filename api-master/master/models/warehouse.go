package models

type Warehouse struct {
	ID         string `json:"warehouse_id"`
	WHName     string `json:"warehouse_name"`
	WHCapacity string `json:"warehouse_capasity"`
	Location   string `json:"location"`
	WHTypeID   string `json:"warehouse_type"`
}

type WHInfo struct {
	WHID        string `json:"warehouse_id"`
	WHName      string `json:"warehouse_name"`
	WHCap       string `json:"warehouse_capacity"`
	GoodsID     string `json:"goods_id"`
	GoodsName   string `json:"goods_name"`
	GoodsQty    string `json:"goods_quantity"`
	GoodsCap    string `json:"goods_capacity"`
	TotGoodsCap string `json:"total_goods_capacity"`
	WHType      string `json:"warehouse_type"`
}
