package port

import (
	port_shared "manager/shared/port"
)

type RegionalIRepository interface {
	port_shared.IRepositoryCRUD
	//GetAll(conditions ...interface{}) []*entity.Regional
	//Remove(*entity.Regional) error
	//Get(int64) (*entity.Regional, error)
	//Save(*entity.Regional) (*entity.Regional, error)
	//SqlQueryRow(string) *sql.Row
	//SqlQueryRows(string) (*sql.Rows, error)
	//SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	//WithTransaction(transaction port_shared.ITransaction) RegionalIRepository
	//SqlQueryData(columns, table, where, sqlTemplate string, order ...string) (map[string]interface{}, error)
}
