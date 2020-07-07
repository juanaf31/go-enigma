package goodsusecase

import "liveCodeAPI/api-master/master/models"

type GoodsUsecase interface {
	GetGoods() ([]*models.Goods, error)
	GetGoodsByID(id string) (*models.Goods, error)
	DeleteGoods(id string) error
	AddGoods(*models.Goods) (*models.Goods, error)
	UpdateGoods(string, *models.Goods) (*models.Goods, error)
}
