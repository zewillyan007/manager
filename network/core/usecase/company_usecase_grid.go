package usecase

import (
	"manager/network/core/port"
	"manager/shared/grid"
	"strconv"
	"strings"
)

type CompanyUseCaseGrid struct {
	grid.Grid
	repository port.CompanyIRepository
}

func NewCompanyUseCaseGrid(repository port.CompanyIRepository) *CompanyUseCaseGrid {
	return &CompanyUseCaseGrid{repository: repository}
}

func (o *CompanyUseCaseGrid) table() string {

	return "sale.company"
}

func (o *CompanyUseCaseGrid) columns() []string {

	return []string{
		"id",
		"id_company",
		"id_person",
		"id_origin_registration",
		"document",
		"email",
		"ddd",
		"telephone",
		"active",
		"notification_key",
		"more_infor",
	}
}

func (o *CompanyUseCaseGrid) mandatory() []string {

	return []string{
		"document",
		"email",
	}
}

func (o *CompanyUseCaseGrid) searchFields() map[string]string {

	return map[string]string{
		"name":     "string",
		"document": "string",
		"email":    "string",
	}
}

func (o *CompanyUseCaseGrid) orderFields() map[string]string {

	return map[string]string{
		"name":     "string",
		"document": "string",
		"email":    "string",
	}
}

func (o *CompanyUseCaseGrid) Execute(GridConfig *grid.GridConfig) (data map[string]interface{}, err error) {

	var sql string = ""
	var order string = ""
	var page, limit float64
	var where []string = []string{}

	prepare, err := o.Prepare(GridConfig, o.columns(), o.mandatory(), o.searchFields())
	if err != nil {
		return nil, err
	}
	params := prepare["params"].(*grid.Params)
	orders := prepare["orders"].(*grid.Orders)

	where = append(where, "(status <> 'DELETED')")

	if len(params.ToString()) > 0 {
		where = append(where, params.ToString())
	}

	if GridConfig.UseSqlFieldsSearch() && len(params.ToStringSearch()) > 0 {
		where = append(where, "("+params.ToStringSearch()+")")
	}

	if len(orders.GetList()) > 0 {
		order = orders.ToStringTranslate(o.orderFields())
	}

	sql = "SELECT %s FROM %s"
	if len(where) > 0 {
		sql = sql + " WHERE %s"
	}

	page, _ = strconv.ParseFloat(GridConfig.Page, 64)
	limit, _ = strconv.ParseFloat(GridConfig.RowsPage, 64)

	if GridConfig.UseSqlQueryPaginator() {
		data, err = o.repository.SqlQueryPaginator(strings.Join(o.columns(), ","), o.table(), strings.Join(where, " AND "), sql, page, limit, order)
	} else {
		data, err = o.repository.SqlQueryData(strings.Join(o.columns(), ","), o.table(), strings.Join(where, " AND "), sql, order)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}
