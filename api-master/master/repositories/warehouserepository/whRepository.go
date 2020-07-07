package warehouserepository

import "liveCodeAPI/api-master/master/models"

type WHRepository interface {
	GetAll() ([]*models.Warehouse, error)
	GetByID(string) (*models.Warehouse, error)
	Delete(string) error
	Add(*models.Warehouse) (*models.Warehouse, error)
	Update(string, *models.Warehouse) (*models.Warehouse, error)
	Info() ([]*models.WHInfo, error)
}
