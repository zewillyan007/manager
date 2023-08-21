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

type PersonTypeHandlerRest struct {
	adapter_shared.Handler
	resource *resource.ServerResource
}

func NewPersonTypeHandlerRest(resource *resource.ServerResource) *PersonTypeHandlerRest {
	return &PersonTypeHandlerRest{
		resource: resource,
	}
}

func (h *PersonTypeHandlerRest) MakeRoutes() {

	h.ConfigError(constant.HDR_PERSON_TYPE, h.resource.Herror)

	router := h.resource.DefaultRouter("/person-types", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
}

func (h *PersonTypeHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var persontypeDtoIn *dto.PersonTypeDtoIn
		var persontypeDtoOut *dto.PersonTypeDtoOut

		persontypeDtoIn = dto.NewPersonTypeDtoIn()
		h.resource.Restful.BindDataReq(w, r, &persontypeDtoIn)
		persontypeDtoOut, err = service.NewPersonTypeService(h.resource.Provider(r)).Get(persontypeDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_PERSON_TYPE_GET, codeErr))
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, persontypeDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *PersonTypeHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		persontypes := service.NewPersonTypeService(h.resource.Provider(r)).GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, persontypes)
	})
}
