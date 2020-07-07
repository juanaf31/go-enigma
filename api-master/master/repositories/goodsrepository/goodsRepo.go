package goodsrepository

import "liveCodeAPI/api-master/master/models"

type GoodsRepository interface {
	GetAll() ([]*models.Goods, error)
	GetByID(string) (*models.Goods, error)
	Delete(string) error
	Add(*models.Goods) (*models.Goods, error)
	Update(string, *models.Goods) (*models.Goods, error)
}
