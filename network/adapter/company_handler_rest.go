package adapter

import (
	"manager/network/core/domain/dto"
	"manager/network/core/service"
	adapter_shared "manager/shared/adapter"
	"manager/shared/connection/audit"
	"manager/shared/constant"
	"manager/shared/grid"
	"manager/shared/resource"
	"net/http"
	"strconv"
)

type CompanyHandlerRest struct {
	adapter_shared.Handler
	resource *resource.ServerResource
}

func NewCompanyHandlerRest(resource *resource.ServerResource) *CompanyHandlerRest {
	return &CompanyHandlerRest{
		resource: resource,
	}
}

func (h *CompanyHandlerRest) MakeRoutes() {

	h.ConfigError(constant.HDR_COMPANY, h.resource.Herror)

	router := h.resource.DefaultRouter("/companies", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *CompanyHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var companyDtoIn *dto.CompanyDtoIn
		var companyDtoOut *dto.CompanyDtoOut

		companyDtoIn = dto.NewCompanyDtoIn()
		h.resource.Restful.BindDataReq(w, r, &companyDtoIn)
		companyDtoOut, err = service.NewCompanyService(h.resource.Provider(r)).Get(companyDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_COMPANY_GET, codeErr))
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, companyDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *CompanyHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		companys := service.NewCompanyService(h.resource.Provider(r)).GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, companys)
	})
}

func (h *CompanyHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		companyDtoIn := dto.NewCompanyDtoIn()
		h.resource.Restful.BindDataReq(w, r, &companyDtoIn)
		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Insert)

		err = service.NewCompanyService(h.resource.Provider(r)).WithTransaction(transaction).Save(companyDtoIn)

		if err != nil {
			transaction.Rollback(err)
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_COMPANY_CREATE, codeErr))
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *CompanyHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		CompanyDtoIn := dto.NewCompanyDtoIn()
		h.resource.Restful.BindDataReq(w, r, &CompanyDtoIn)
		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Update)

		err = service.NewCompanyService(h.resource.Provider(r)).WithTransaction(transaction).Save(CompanyDtoIn)

		if err != nil {
			transaction.Rollback(err)
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_COMPANY_SAVE, codeErr))
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *CompanyHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		CompanyDtoIn := dto.NewCompanyDtoIn()
		h.resource.Restful.BindDataReq(w, r, &CompanyDtoIn)
		err = service.NewCompanyService(h.resource.Provider(r)).Remove(CompanyDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_COMPANY_REMOVE, codeErr))
		} else {
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *CompanyHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		GridConfig := grid.NewGridConfig()
		h.resource.Restful.BindDataReq(w, r, &GridConfig)
		GridConfig = h.GridConfigData(GridConfig)
		dataGrid, err := service.NewCompanyService(h.resource.Provider(r)).Grid(GridConfig)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_COMPANY_GRID, codeErr))
		} else {
			if GridConfig.Export != nil && len(GridConfig.Export.Value) > 0 {
				grid.ResponseDataGrid(w, GridConfig.Export.Type, dataGrid, "product")
			} else {
				h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
			}
		}
	})
}
