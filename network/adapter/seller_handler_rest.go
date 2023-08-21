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

type SellerHandlerRest struct {
	adapter_shared.Handler
	resource *resource.ServerResource
}

func NewSellerHandlerRest(resource *resource.ServerResource) *SellerHandlerRest {
	return &SellerHandlerRest{
		resource: resource,
	}
}

func (h *SellerHandlerRest) MakeRoutes() {

	h.ConfigError(constant.HDR_SELLER, h.resource.Herror)

	router := h.resource.DefaultRouter("/sellers", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *SellerHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var sellerDtoIn *dto.SellerDtoIn
		var sellerDtoOut *dto.SellerDtoOut

		sellerDtoIn = dto.NewSellerDtoIn()
		h.resource.Restful.BindDataReq(w, r, &sellerDtoIn)
		sellerDtoOut, err = service.NewSellerService(h.resource.Provider(r)).Get(sellerDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_SELLER_GET, codeErr))
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, sellerDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *SellerHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sellers := service.NewSellerService(h.resource.Provider(r)).GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, sellers)
	})
}

func (h *SellerHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		sellerDtoIn := dto.NewSellerDtoIn()
		h.resource.Restful.BindDataReq(w, r, &sellerDtoIn)
		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Insert)

		err = service.NewSellerService(h.resource.Provider(r)).WithTransaction(transaction).Save(sellerDtoIn)

		if err != nil {
			transaction.Rollback(err)
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_SELLER_CREATE, codeErr))
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *SellerHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		SellerDtoIn := dto.NewSellerDtoIn()
		h.resource.Restful.BindDataReq(w, r, &SellerDtoIn)
		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Update)

		err = service.NewSellerService(h.resource.Provider(r)).WithTransaction(transaction).Save(SellerDtoIn)

		if err != nil {
			transaction.Rollback(err)
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_SELLER_SAVE, codeErr))
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *SellerHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		SellerDtoIn := dto.NewSellerDtoIn()
		h.resource.Restful.BindDataReq(w, r, &SellerDtoIn)
		err = service.NewSellerService(h.resource.Provider(r)).Remove(SellerDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_SELLER_REMOVE, codeErr))
		} else {
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *SellerHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		GridConfig := grid.NewGridConfig()
		h.resource.Restful.BindDataReq(w, r, &GridConfig)
		GridConfig = h.GridConfigData(GridConfig)
		dataGrid, err := service.NewSellerService(h.resource.Provider(r)).Grid(GridConfig)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_SELLER_GRID, codeErr))
		} else {
			if GridConfig.Export != nil && len(GridConfig.Export.Value) > 0 {
				grid.ResponseDataGrid(w, GridConfig.Export.Type, dataGrid, "product")
			} else {
				h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
			}
		}
	})
}
