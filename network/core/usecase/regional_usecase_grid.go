package usecase

import (
	"manager/network/core/port"
	"manager/shared/grid"
	"strconv"
	"strings"
)

type RegionalUseCaseGrid struct {
	grid.Grid
	repository port.RegionalIRepository
}

func NewRegionalUseCaseGrid(repository port.RegionalIRepository) *RegionalUseCaseGrid {
	return &RegionalUseCaseGrid{repository: repository}
}

func (o *RegionalUseCaseGrid) table() string {

	return "network.regional"
}

func (o *RegionalUseCaseGrid) columns() []string {

	return []string{
		"id",
		"id_parent",
		"name",
		"short_name",
		"document",
		"document_type",
		"telephone",
		"address",
		"status",
		"type",
		"birthday",
	}
}

func (o *RegionalUseCaseGrid) mandatory() []string {

	return []string{}
}

func (o *RegionalUseCaseGrid) searchFields() map[string]string {

	return map[string]string{
		"name":          "string",
		"short_name":    "string",
		"document":      "string",
		"document_type": "string",
		"status":        "string",
		"type":          "string",
	}
}

func (o *RegionalUseCaseGrid) orderFields() map[string]string {

	return map[string]string{
		"name":          "string",
		"short_name":    "string",
		"document":      "string",
		"document_type": "string",
		"status":        "string",
		"type":          "string",
	}
}

func (o *RegionalUseCaseGrid) Execute(GridConfig *grid.GridConfig) (data map[string]interface{}, err error) {

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

	// where = append(where, "(status <> 'DELETED')")

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
