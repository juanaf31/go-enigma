package warehouseusecase

import "liveCodeAPI/api-master/master/models"

type WHUsecase interface {
	GetWarehouses() ([]*models.Warehouse, error)
	GetWarehouseByID(id string) (*models.Warehouse, error)
	DeleteWarehouse(id string) error
	AddWarehouse(*models.Warehouse) (*models.Warehouse, error)
	UpdateWarehouse(string, *models.Warehouse) (*models.Warehouse, error)
	GetWHInfo() ([]*models.WHInfo, error)
}
