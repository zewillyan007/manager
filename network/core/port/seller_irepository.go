package port

import (
	port_shared "manager/shared/port"
)

type SellerIRepository interface {
	port_shared.IRepositoryCRUD
	//GetAll(conditions ...interface{}) []*entity.Seller
	//Remove(*entity.Seller) error
	//Get(int64) (*entity.Seller, error)
	//Save(*entity.Seller) (*entity.Seller, error)
	//SqlQueryRow(string) *sql.Row
	//SqlQueryRows(string) (*sql.Rows, error)
	//SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	//WithTransaction(transaction port_shared.ITransaction) SellerIRepository
	//SqlQueryData(columns, table, where, sqlTemplate string, order ...string) (map[string]interface{}, error)
}
