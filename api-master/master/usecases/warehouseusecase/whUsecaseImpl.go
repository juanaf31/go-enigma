package warehouseusecase

import (
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/repositories/warehouserepository"
	"log"
)

type WHUsecaseImpl struct {
	whRepo warehouserepository.WHRepository
}

func InitWHUsecase(whRepo warehouserepository.WHRepository) WHUsecase {
	return &WHUsecaseImpl{whRepo: whRepo}
}

func (c *WHUsecaseImpl) GetWarehouses() ([]*models.Warehouse, error) {
	warehouses, err := c.whRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (c *WHUsecaseImpl) GetWarehouseByID(id string) (*models.Warehouse, error) {
	warehouse, err := c.whRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}

func (c *WHUsecaseImpl) DeleteWarehouse(id string) error {
	err := c.whRepo.Delete(id)
	return err
}

func (c *WHUsecaseImpl) AddWarehouse(data *models.Warehouse) (*models.Warehouse, error) {
	warehouse, err := c.whRepo.Add(data)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}

func (c *WHUsecaseImpl) UpdateWarehouse(id string, data *models.Warehouse) (*models.Warehouse, error) {
	warehouse, err := c.whRepo.Update(id, data)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}

func (c *WHUsecaseImpl) GetWHInfo() ([]*models.WHInfo, error) {
	whInfo, err := c.whRepo.Info()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return whInfo, nil
}
