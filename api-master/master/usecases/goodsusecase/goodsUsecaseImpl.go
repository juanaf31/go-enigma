package goodsusecase

import (
	"errors"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/repositories/goodsrepository"
	"liveCodeAPI/utils"
)

type GoodsUsecaseImpl struct {
	goodsRepo goodsrepository.GoodsRepository
}

func InitGoodsUsecase(goodsRepo goodsrepository.GoodsRepository) GoodsUsecase {
	return &GoodsUsecaseImpl{goodsRepo: goodsRepo}
}

func (c *GoodsUsecaseImpl) GetGoods() ([]*models.Goods, error) {
	goods, err := c.goodsRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (c *GoodsUsecaseImpl) GetGoodsByID(id string) (*models.Goods, error) {
	goods, err := c.goodsRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (c *GoodsUsecaseImpl) DeleteGoods(id string) error {
	err := c.goodsRepo.Delete(id)
	return err
}

func (c *GoodsUsecaseImpl) AddGoods(data *models.Goods) (*models.Goods, error) {
	isValid := utils.ValidationGoods(data)
	if isValid == false {
		err := errors.New("Fields are not allowed null")
		return nil, err
	} else {
		goods, err := c.goodsRepo.Add(data)
		if err != nil {
			return nil, err
		}
		return goods, nil
	}
}

func (c *GoodsUsecaseImpl) UpdateGoods(id string, data *models.Goods) (*models.Goods, error) {
	isValid := utils.ValidationGoods(data)
	if isValid == false {
		err := errors.New("Fields are not allowed null")
		return nil, err
	} else {
		goods, err := c.goodsRepo.Update(id, data)
		if err != nil {
			return nil, err
		}
		return goods, nil
	}
}
