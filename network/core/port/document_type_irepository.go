package port

import (
	port_shared "manager/shared/port"
)

type DocumentTypeIRepository interface {
	port_shared.IRepositoryCRUD
	//GetAll(conditions ...interface{}) []*entity.DocumentType
	//Remove(*entity.DocumentType) error
	//Get(int64) (*entity.DocumentType, error)
	//Save(*entity.DocumentType) (*entity.DocumentType, error)
	//SqlQueryRow(string) *sql.Row
	//SqlQueryRows(string) (*sql.Rows, error)
	//SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	//WithTransaction(transaction port_shared.ITransaction) DocumentTypeIRepository
	//SqlQueryData(columns, table, where, sqlTemplate string, order ...string) (map[string]interface{}, error)
}
