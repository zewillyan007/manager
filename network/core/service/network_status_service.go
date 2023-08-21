package service

import (
	"fmt"
	"manager/network/core/domain/dto"
	"manager/network/core/port"
	"manager/network/core/usecase"
	port_shared "manager/shared/port"
	"manager/shared/types"
	"strconv"
)

type NetworkStatusService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.NetworkStatusIRepository
	ucGet      *usecase.NetworkStatusUseCaseGet
	ucGetAll   *usecase.NetworkStatusUseCaseGetAll
}

func NewNetworkStatusService(provider port_shared.IResourceProvider) *NetworkStatusService {
	repo := types.GetConstructor((*port.NetworkStatusIRepository)(nil))(provider.GetDB()).(port.NetworkStatusIRepository)
	repo.SetContext(provider.Context())

	return &NetworkStatusService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewNetworkStatusUseCaseGet(repo),
		ucGetAll:   usecase.NewNetworkStatusUseCaseGetAll(repo),
	}
}

func (o *NetworkStatusService) WithTransaction(transaction port_shared.ITransaction) *NetworkStatusService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *NetworkStatusService) Get(dtoIn *dto.NetworkStatusDtoIn) (*dto.NetworkStatusDtoOut, error) {

	id, _ := strconv.Atoi(dtoIn.Id)
	NetworkStatus, err := o.ucGet.Execute(int64(id))
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewNetworkStatusDtoOut()

	dtoOut.Id = fmt.Sprintf("%d", NetworkStatus.Id)
	dtoOut.Name = NetworkStatus.Name
	dtoOut.Mnemonic = NetworkStatus.Mnemonic
	dtoOut.Hint = NetworkStatus.Hint

	if NetworkStatus.CreationDateTime != nil {
		dtoOut.CreationDateTime = NetworkStatus.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if NetworkStatus.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = NetworkStatus.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	if NetworkStatus.DisableDateTime != nil {
		dtoOut.DisableDateTime = NetworkStatus.DisableDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *NetworkStatusService) GetAll(conditions ...interface{}) []*dto.NetworkStatusDtoOut {

	var arrayNetworkStatusDto []*dto.NetworkStatusDtoOut

	arrayNetworkStatus := o.ucGetAll.Execute(conditions...)

	for _, NetworkStatus := range arrayNetworkStatus {

		dtoOut := dto.NewNetworkStatusDtoOut()

		dtoOut.Id = fmt.Sprintf("%d", NetworkStatus.Id)
		dtoOut.Name = NetworkStatus.Name
		dtoOut.Mnemonic = NetworkStatus.Mnemonic
		dtoOut.Hint = NetworkStatus.Hint

		if NetworkStatus.CreationDateTime != nil {
			dtoOut.CreationDateTime = NetworkStatus.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if NetworkStatus.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = NetworkStatus.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		if NetworkStatus.DisableDateTime != nil {
			dtoOut.DisableDateTime = NetworkStatus.DisableDateTime.Format("2006-01-02 15:04:05")
		}

		arrayNetworkStatusDto = append(arrayNetworkStatusDto, dtoOut)
	}
	return arrayNetworkStatusDto
}
