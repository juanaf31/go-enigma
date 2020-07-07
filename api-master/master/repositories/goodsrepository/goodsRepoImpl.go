package goodsrepository

import (
	"database/sql"
	"liveCodeAPI/api-master/master/models"
)

type GoodsRepoImpl struct {
	db *sql.DB
}

func InitGoodsRepoImpl(db *sql.DB) GoodsRepository {
	return &GoodsRepoImpl{db: db}
}

func (s *GoodsRepoImpl) GetAll() ([]*models.Goods, error) {
	query := `select * from goods`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listGoods []*models.Goods
	for rows.Next() {
		goods := models.Goods{}
		err := rows.Scan(&goods.ID, &goods.Goods_name, &goods.Quantity, &goods.Capacity, &goods.WarehouseID)
		if err != nil {
			return nil, err
		}
		listGoods = append(listGoods, &goods)
	}
	return listGoods, nil

}

func (s *GoodsRepoImpl) GetByID(id string) (*models.Goods, error) {

	query := `select * from goods where id=?`
	row := s.db.QueryRow(query, id)
	var goods = models.Goods{}
	err := row.Scan(&goods.ID, &goods.Goods_name, &goods.Quantity, &goods.Capacity, &goods.WarehouseID)
	if err != nil {
		return nil, err
	}
	return &goods, nil

}

func (s *GoodsRepoImpl) Delete(id string) error {
	query := "delete from goods where id=?"
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}

func (s *GoodsRepoImpl) Add(goods *models.Goods) (*models.Goods, error) {

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(`insert into goods values(?,?,?,?,?)`, &goods.ID, &goods.Goods_name, &goods.Quantity, &goods.Capacity, &goods.WarehouseID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return goods, nil
}

func (s *GoodsRepoImpl) Update(id string, goods *models.Goods) (*models.Goods, error) {
	tx, err := s.db.Begin()
	_, err = tx.Exec("update goods set id=?, goods_name=?,quantity=?, capacity=?, warehouse_id=? where id=?", &goods.ID, &goods.Goods_name, &goods.Quantity, &goods.Capacity, &goods.WarehouseID, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return goods, tx.Commit()
}
