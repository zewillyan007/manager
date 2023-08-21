package port

import (
	port_shared "manager/shared/port"
)

type PersonTypeIRepository interface {
	port_shared.IRepositoryCRUD
	//GetAll(conditions ...interface{}) []*entity.PersonType
	//Remove(*entity.PersonType) error
	//Get(int64) (*entity.PersonType, error)
	//Save(*entity.PersonType) (*entity.PersonType, error)
	//SqlQueryRow(string) *sql.Row
	//SqlQueryRows(string) (*sql.Rows, error)
	//SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	//WithTransaction(transaction port_shared.ITransaction) PersonTypeIRepository
	//SqlQueryData(columns, table, where, sqlTemplate string, order ...string) (map[string]interface{}, error)
}
