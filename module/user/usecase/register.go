package usecase

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"
	"errors"
)

type RegisterUC struct {
	userQueryRepo UserQueryRepository
	userCmdRepo   UserCmdRepository
	hasher        Hasher
}

func NewRegisterUC(userRepoQuery UserQueryRepository, userCmdRepo UserCmdRepository, hasher Hasher) *RegisterUC {
	return &RegisterUC{userQueryRepo: userRepoQuery, userCmdRepo: userCmdRepo, hasher: hasher}
}

func (uc *RegisterUC) Register(ctx context.Context, dto EmailPasswordRegistration) error {
	// 1. Find User by email
	// 1.1 Found : return error (email exited)
	// 2. Generate salt
	// 3. Hash password + salt
	// 4. Create user entity

	user, err := uc.userQueryRepo.FindByEmail(ctx, dto.Email)

	if user != nil {
		return domain.ErrEmailHasExisted
	}

	if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
		return err
	}

	salt, err := uc.hasher.RandomStr(30)
	if err != nil {
		return err
	}

	hashPassword, err := uc.hasher.HashPassword(salt, dto.Password)
	if err != nil {
		return err
	}
	userEntity, err := domain.NewUser(
		common.GenUUID(),
		dto.FirstName,
		dto.LastName,
		dto.Email,
		hashPassword,
		salt,
		domain.RoleUser,
	)

	if err != nil {
		return err
	}
	if err := uc.userCmdRepo.Create(ctx, userEntity); err != nil {
		return err
	}

	return nil
}
