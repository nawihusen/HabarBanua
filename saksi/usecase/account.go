package usecase

import (
	"be-service-saksi-management/domain"
)

// accountUsecase is struct SQL
type accountUsecase struct {
	accountMySQLRepo domain.AccountMySQLRepository
	// saksiGRPCRepo       domain.SaksiGRPCRepository
	// saksiGRPCAuthRepo   domain.SaksiGRPCAuthRepository
	// accountRedisRepo    domain.AccountRedisRepository
	// adsGRPCRepo         domain.AreaGRPCRepository
	// tpsMySQLRepo        domain.TpsMySQLRepository
	// attendanceMySQLRepo domain.AttendanceMySQLRepository
	// recruiterRepo       domain.RecruiterMySQLRepository
	// userGRPCAuthRepo    domain.UserGRPCUserRepository
	// motionPayRepo       domain.MotionPayRepository
}

// NewAccountUsecase is constructor of account usecase
func NewAccountUsecase(accountMySQLRepo domain.AccountMySQLRepository) domain.AccountUsecase {
	return &accountUsecase{
		accountMySQLRepo: accountMySQLRepo,
		// saksiGRPCRepo:       saksiGRPCRepo,
		// saksiGRPCAuthRepo:   saksiGRPCAuthRepo,
		// accountRedisRepo:    accountRedisRepo,
		// adsGRPCRepo:         adsGRPCRepo,
		// tpsMySQLRepo:        tpsMySQLRepo,
		// attendanceMySQLRepo: attendanceMySQLRepo,
		// recruiterRepo:       recruiterRepo,
		// userGRPCAuthRepo:    userGRPCAuthRepo,
		// motionPayRepo:       motionPayRepo,
	}
}

// DeleteSaksiValidation is constructor of account usecase
// func (au *accountUsecase) DeleteSaksiValidation(ctx context.Context, id int64) error {
// 	err := au.accountMySQLRepo.RemoveSaksiValidation(ctx, id)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	return nil
// }
