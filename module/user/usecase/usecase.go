package usecase

import (
	"Ls04_GORM/module/user/domain"
	"context"

	"github.com/google/uuid"
)

type UseCase interface {
	Register(ctx context.Context, dto EmailPasswordRegistration) error
	Login(ctx context.Context, dto EmailPasswordLogin) (*TokenResponse, error)
}

type useCase struct {
	*LoginUC
	*RegisterUC
}

func NewUseCase(repo UserRepository, sessionRepo SessionRepository, hasher Hasher, tokenProvider TokenProvider) *useCase {
	return &useCase{
		LoginUC:    NewLoginUC(repo, sessionRepo, hasher, tokenProvider),
		RegisterUC: NewRegisterUC(repo, repo, hasher),
	}
}

// 3

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
}

type SessionRepository interface {
	SessionQueryRepository
	SessionCmdRepository
}
type SessionQueryRepository interface {
	Find(ctx context.Context, email string) (*domain.Session, error)
}
type SessionCmdRepository interface {
	Create(ctx context.Context, data *domain.Session) error
}
