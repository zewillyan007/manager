package port

import (
	"manager/shared/connection/audit"

	"gorm.io/gorm"
)

type ITransaction interface {
	GetTransaction() *gorm.DB
	GetEnvelope() *audit.Envelope
}
