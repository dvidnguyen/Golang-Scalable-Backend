package usecase

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"

	"github.com/google/uuid"
)

type UseCase interface {
	Register(ctx context.Context, dto EmailPasswordRegistration) error
	Login(ctx context.Context, dto EmailPasswordLogin) (*TokenResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error)
}

type useCase struct {
	*LoginUC
	*RegisterUC
	*refreshTokenUC
}

func NewUseCase(repo UserRepository, sessionRepo SessionRepository, hasher Hasher, tokenProvider TokenProvider) *useCase {
	return &useCase{
		LoginUC:    NewLoginUC(repo, sessionRepo, hasher, tokenProvider),
		RegisterUC: NewRegisterUC(repo, repo, hasher),
	}
}

type Builder interface {
	BuildUserQueryRepo() UserQueryRepository
	BuildUserCmdRepo() UserCmdRepository
	BuildHasher() Hasher
	BuildTokenProvider() TokenProvider
	BuildSessionQueryRepo() SessionQueryRepository
	BuildSessionCmdRepo() SessionCmdRepository
	BuildSessionRepo() SessionRepository
}

// 3
func UseCaseWithBuilder(b Builder) UseCase {
	return &useCase{
		RegisterUC:     NewRegisterUC(b.BuildUserQueryRepo(), b.BuildUserCmdRepo(), b.BuildHasher()),
		LoginUC:        NewLoginUC(b.BuildUserQueryRepo(), b.BuildSessionRepo(), b.BuildHasher(), b.BuildTokenProvider()),
		refreshTokenUC: NewRefreshTokenUC(b.BuildUserQueryRepo(), b.BuildSessionRepo(), b.BuildTokenProvider(), b.BuildHasher()),
	}
}

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type TokenProvider interface {
	IssueToken(ctx context.Context, id, sub string) (token string, err error)
	TokenExpireInSeconds() int
	RefreshExpireInSeconds() int
}

type UserRepository interface {
	UserQueryRepository
	UserCmdRepository
}

type UserQueryRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
type UserCmdRepository interface {
	Create(ctx context.Context, data *domain.User) error
	Update(ctx context.Context, data *domain.User) error
}

type SessionRepository interface {
	SessionQueryRepository
	SessionCmdRepository
}
type SessionQueryRepository interface {
	Find(ctx context.Context, email string) (*domain.Session, error)
	FindByRefreshToken(ctx context.Context, rt string) (*domain.Session, error)
	CountSessionByUserId(ctx context.Context, userId uuid.UUID) (int64, error)
}
type SessionCmdRepository interface {
	Create(ctx context.Context, data *domain.Session) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ImgRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*common.Image, error)
	SetImageStatusActivated(ctx context.Context, id uuid.UUID) error
}
