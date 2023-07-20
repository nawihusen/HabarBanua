package domain

import (
	"time"
)

type Saksi struct {
	ID           int64
	MemberNo     string
	RecruiterID  int64
	ProvCode     int64
	KabCode      int64
	KecCode      int64
	KelCode      int64
	TpsIDRequest int64
	TpsID        *int64
	Status       string `json:"status" form:"status" validate:"required"`
	DtmCrt       time.Time
	DtmUpd       time.Time
}

// AccountUsecase is Account usecase
type AccountUsecase interface {
}

// AccountMySQLRepository is Account repository in MySQL
type AccountMySQLRepository interface {
}

type AccountRedisRepository interface {
}

// StructureGRPCRepository is Saksi repository in gRPC
type SaksiGRPCRepository interface {
}

// SaksiGRPCAuthRepository is Structure repository in gRPC
type SaksiGRPCAuthRepository interface {
}

type UserGRPCUserRepository interface {
}

// AreaGRPCRepository is Area repository in gRPC
type AreaGRPCRepository interface {
}

// RealCountGRPCRepository is Real count mgt repository in gRPC
type RealCountGRPCRepository interface {
}
