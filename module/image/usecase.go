package image

import (
	"Ls04_GORM/common"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/viettranx/service-context/core"
)

type Uploader interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) error
	GetName() string
	GetDomain() string
}
type UploadImageUC struct {
	repo     CmdRepository
	uploader Uploader
}

func (uc UploadImageUC) UploadImage(ctx context.Context, dto UploadDTO) (*Image, error) {
	dstFileName := fmt.Sprintf("%d_%s", time.Now().UTC().UnixNano(), dto.FileName)
	if err := uc.uploader.SaveFileUploaded(ctx, dto.FileData, dstFileName); err != nil {
		return nil, core.ErrInternalServerError.WithError(ErrCannotUploadImage.Error()).WithDebug(err.Error())
	}
	now := time.Now().UTC()
	image := Image{
		Id:              common.GenUUID(),
		Title:           dto.Name,
		FileName:        dstFileName,
		FileSize:        dto.FileSize,
		FileType:        dto.FileType,
		StorageProvider: uc.uploader.GetName(),
		Status:          StatusUploaded,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	if err := uc.repo.Create(ctx, &image); err != nil {
		return nil, core.ErrInternalServerError.WithError(ErrCannotUploadImage.Error()).WithDebug(err.Error())
	}

	return &image, nil

}

func NewUseCase(uploader Uploader, repo CmdRepository) UploadImageUC {
	return UploadImageUC{uploader: uploader, repo: repo}
}

type CmdRepository interface {
	Create(ctx context.Context, img *Image) error
}

var (
	ErrCannotUploadImage = errors.New("cannot upload image")
	ErrCannotFindImage   = errors.New("cannot find image")
)
