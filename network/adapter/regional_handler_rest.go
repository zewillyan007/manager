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

type RegionalHandlerRest struct {
	adapter_shared.Handler
	resource *resource.ServerResource
}

func NewRegionalHandlerRest(resource *resource.ServerResource) *RegionalHandlerRest {
	return &RegionalHandlerRest{
		resource: resource,
	}
}

func (h *RegionalHandlerRest) MakeRoutes() {

	h.ConfigError(constant.HDR_REGIONAL, h.resource.Herror)

	router := h.resource.DefaultRouter("/regionals", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *RegionalHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var regionalDtoIn *dto.RegionalDtoIn
		var regionalDtoOut *dto.RegionalDtoOut

		regionalDtoIn = dto.NewRegionalDtoIn()
		h.resource.Restful.BindDataReq(w, r, &regionalDtoIn)
		regionalDtoOut, err = service.NewRegionalService(h.resource.Provider(r)).Get(regionalDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_REGIONAL_GET, codeErr))
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, regionalDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *RegionalHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		regionals := service.NewRegionalService(h.resource.Provider(r)).GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, regionals)
	})
}

func (h *RegionalHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		regionalDtoIn := dto.NewRegionalDtoIn()
		h.resource.Restful.BindDataReq(w, r, &regionalDtoIn)
		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Insert)

		err = service.NewRegionalService(h.resource.Provider(r)).WithTransaction(transaction).Save(regionalDtoIn)

		if err != nil {
			transaction.Rollback(err)
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_REGIONAL_CREATE, codeErr))
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *RegionalHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		RegionalDtoIn := dto.NewRegionalDtoIn()
		h.resource.Restful.BindDataReq(w, r, &RegionalDtoIn)
		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Update)

		err = service.NewRegionalService(h.resource.Provider(r)).WithTransaction(transaction).Save(RegionalDtoIn)

		if err != nil {
			transaction.Rollback(err)
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_REGIONAL_SAVE, codeErr))
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *RegionalHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		RegionalDtoIn := dto.NewRegionalDtoIn()
		h.resource.Restful.BindDataReq(w, r, &RegionalDtoIn)
		err = service.NewRegionalService(h.resource.Provider(r)).Remove(RegionalDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_REGIONAL_REMOVE, codeErr))
		} else {
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *RegionalHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		GridConfig := grid.NewGridConfig()
		h.resource.Restful.BindDataReq(w, r, &GridConfig)
		GridConfig = h.GridConfigData(GridConfig)
		dataGrid, err := service.NewRegionalService(h.resource.Provider(r)).Grid(GridConfig)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_REGIONAL_GRID, codeErr))
		} else {
			if GridConfig.Export != nil && len(GridConfig.Export.Value) > 0 {
				grid.ResponseDataGrid(w, GridConfig.Export.Type, dataGrid, "product")
			} else {
				h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
			}
		}
	})
}
