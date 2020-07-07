package warehouserepository

import (
	"database/sql"
	"liveCodeAPI/api-master/master/models"
)

type WHRepoImpl struct {
	db *sql.DB
}

func InitWHRepoImpl(db *sql.DB) WHRepository {
	return &WHRepoImpl{db: db}
}

func (s *WHRepoImpl) GetAll() ([]*models.Warehouse, error) {
	query := `select * from warehouse`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listWH []*models.Warehouse
	for rows.Next() {
		warehouse := models.Warehouse{}
		err := rows.Scan(&warehouse.ID, &warehouse.WHName, &warehouse.WHCapacity, &warehouse.Location, &warehouse.WHTypeID)
		if err != nil {
			return nil, err
		}
		listWH = append(listWH, &warehouse)
	}
	return listWH, nil

}

func (s *WHRepoImpl) GetByID(id string) (*models.Warehouse, error) {

	query := `select * from warehouse where id=?`
	row := s.db.QueryRow(query, id)
	var warehouse = models.Warehouse{}
	err := row.Scan(&warehouse.ID, &warehouse.WHName, &warehouse.WHCapacity, &warehouse.Location, &warehouse.WHTypeID)
	if err != nil {
		return nil, err
	}
	return &warehouse, nil

}

func (s *WHRepoImpl) Delete(id string) error {
	query := "delete from warehouse where id=?"
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

func (s *WHRepoImpl) Add(warehouse *models.Warehouse) (*models.Warehouse, error) {

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(`insert into warehouse values(?,?,?,?,?)`, &warehouse.ID, &warehouse.WHName, &warehouse.WHCapacity, &warehouse.Location, &warehouse.WHTypeID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return warehouse, nil
}

func (s *WHRepoImpl) Update(id string, warehouse *models.Warehouse) (*models.Warehouse, error) {
	tx, err := s.db.Begin()
	_, err = tx.Exec("update warehouse set id=?, warehouse_name=?, warehouse_capacity=?,location=?, warehouse_type_id=? where id=?", &warehouse.ID, &warehouse.WHName, &warehouse.WHCapacity, &warehouse.Location, &warehouse.WHTypeID, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return warehouse, tx.Commit()
}

func (s *WHRepoImpl) Info() ([]*models.WHInfo, error) {
	query := `select tw.wh_id,
	tw.wh_name,
	tw.wh_cap,
	tw.goods_id,tw.goods_name,tw.goods_qty,tw.goods_cap,tw.goods_qty*tw.goods_cap as total_goods_cap, wt.warehouse_type
	 from(
	select w.id as wh_id,w.warehouse_name as wh_name ,w.warehouse_capacity as wh_cap, w.location as wh_loc, 
	w.warehouse_type_id as wh_type, g.id as goods_id, g.goods_name as goods_name, g.quantity as goods_qty, g.capacity as goods_cap
	from warehouse w 
	join goods g on w.id=g.warehouse_id)tw
	join wh_type wt on wt.id = tw.wh_type;`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listWH []*models.WHInfo
	for rows.Next() {
		whInfo := models.WHInfo{}
		err := rows.Scan(&whInfo.WHID, &whInfo.WHName, &whInfo.WHCap, &whInfo.GoodsID, &whInfo.GoodsName, &whInfo.GoodsQty, &whInfo.GoodsCap, &whInfo.TotGoodsCap, &whInfo.WHType)
		if err != nil {
			return nil, err
		}
		listWH = append(listWH, &whInfo)
	}
	return listWH, nil
}
