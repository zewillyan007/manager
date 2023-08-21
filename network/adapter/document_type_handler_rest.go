package adapter

import (
	"manager/network/core/domain/dto"
	"manager/network/core/service"
	adapter_shared "manager/shared/adapter"
	"manager/shared/constant"
	"manager/shared/resource"
	"net/http"
	"strconv"
)

type DocumentTypeHandlerRest struct {
	adapter_shared.Handler
	resource *resource.ServerResource
}

func NewDocumentTypeHandlerRest(resource *resource.ServerResource) *DocumentTypeHandlerRest {
	return &DocumentTypeHandlerRest{
		resource: resource,
	}
}

func (h *DocumentTypeHandlerRest) MakeRoutes() {

	h.ConfigError(constant.HDR_DOCUMENT_TYPE, h.resource.Herror)

	router := h.resource.DefaultRouter("/document-types", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
}

func (h *DocumentTypeHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var documenttypeDtoIn *dto.DocumentTypeDtoIn
		var documenttypeDtoOut *dto.DocumentTypeDtoOut

		documenttypeDtoIn = dto.NewDocumentTypeDtoIn()
		h.resource.Restful.BindDataReq(w, r, &documenttypeDtoIn)
		documenttypeDtoOut, err = service.NewDocumentTypeService(h.resource.Provider(r)).Get(documenttypeDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_DOCUMENT_TYPE_GET, codeErr))
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, documenttypeDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *DocumentTypeHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		documenttypes := service.NewDocumentTypeService(h.resource.Provider(r)).GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, documenttypes)
	})
}
