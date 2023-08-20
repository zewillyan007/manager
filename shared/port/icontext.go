package port

import "manager/shared/types"

type IManagerContext interface {
	GetUser() *types.UserContext
}
