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

type NetworkStatusHandlerRest struct {
	adapter_shared.Handler
	resource *resource.ServerResource
}

func NewNetworkStatusHandlerRest(resource *resource.ServerResource) *NetworkStatusHandlerRest {
	return &NetworkStatusHandlerRest{
		resource: resource,
	}
}

func (h *NetworkStatusHandlerRest) MakeRoutes() {

	h.ConfigError(constant.HDR_NETWORK_STATUS, h.resource.Herror)

	router := h.resource.DefaultRouter("/network-status", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
}

func (h *NetworkStatusHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var networkstatusDtoIn *dto.NetworkStatusDtoIn
		var networkstatusDtoOut *dto.NetworkStatusDtoOut

		networkstatusDtoIn = dto.NewNetworkStatusDtoIn()
		h.resource.Restful.BindDataReq(w, r, &networkstatusDtoIn)
		networkstatusDtoOut, err = service.NewNetworkStatusService(h.resource.Provider(r)).Get(networkstatusDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_NETWORK_STATUS_GET, codeErr))
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, networkstatusDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *NetworkStatusHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		networkstatuss := service.NewNetworkStatusService(h.resource.Provider(r)).GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, networkstatuss)
	})
}
