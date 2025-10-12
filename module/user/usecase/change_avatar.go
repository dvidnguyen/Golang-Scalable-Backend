package usecase

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"
	"fmt"

	"github.com/viettranx/service-context/core"
)

type ChangeAvatarUC struct {
	userQueryRepo UserQueryRepository
	userCmdRepo   UserCmdRepository
	imgRepo       ImgRepository
}

func NewChangeAvatarUC(userQueryRepo UserQueryRepository, userCmdRepo UserCmdRepository, imgRepo ImgRepository) *ChangeAvatarUC {
	return &ChangeAvatarUC{
		userQueryRepo: userQueryRepo,
		userCmdRepo:   userCmdRepo,
		imgRepo:       imgRepo,
	}
}

func (uc *ChangeAvatarUC) ChangeAvatar(ctx context.Context, dto SingleImgDTO) error {
	// 1. Find user by id
	userEntity, err := uc.userQueryRepo.Find(ctx, dto.Requester.UserId())
	if err != nil {
		fmt.Printf("Find user error: %+v\n", err)
		return core.ErrBadRequest.
			WithError(domain.ErrCannotChangeAvatar.Error()).
			WithDebug(err.Error())
	}

	fmt.Printf("Found user: %+v\n", userEntity)
	if err != nil {
		return core.ErrBadRequest.WithError(domain.ErrCannotChangeAvatar.Error()).WithDebug(err.Error())
	}

	img, err := uc.imgRepo.Find(ctx, dto.ImageID)

	if err != nil {
		return core.ErrBadRequest.WithError(domain.ErrCannotChangeAvatar.Error()).WithDebug(err.Error())
	}

	if err := userEntity.ChangeAvatar(img.FileName); err != nil {
		return core.ErrBadRequest.WithError(domain.ErrCannotChangeAvatar.Error()).WithDebug(err.Error())
	}

	if err := uc.userCmdRepo.Update(ctx, userEntity); err != nil {
		return core.ErrBadRequest.WithError(domain.ErrCannotChangeAvatar.Error()).WithDebug(err.Error())
	}
	go func() {
		defer common.Recover()
		_ = uc.imgRepo.SetImageStatusActivated(ctx, dto.ImageID)
	}()

	return nil

}
