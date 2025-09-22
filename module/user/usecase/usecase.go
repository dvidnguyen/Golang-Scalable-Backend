package usecase

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"
	"errors"
)

type UseCase interface {
	Register(ctx context.Context, dto EmailPasswordRegistration) error
	Login(ctx context.Context, dto EmailPasswordLogin) (*TokenResponse, error)
}

type useCase struct {
	repo          UserRepository
	hasher        Hasher
	tokenProvider TokenProvider
	sessionRepo   SessionRepository
}

// 3

func NewUseCase(repo UserRepository, hasher Hasher, tokenProvider TokenProvider, sessionRepo SessionRepository) UseCase {
	return &useCase{repo: repo, hasher: hasher, tokenProvider: tokenProvider, sessionRepo: sessionRepo}
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

// 4
func (uc *useCase) Register(ctx context.Context, dto EmailPasswordRegistration) error {
	// 1. Find User by email
	// 1.1 Found : return error (email exited)
	// 2. Generate salt
	// 3. Hash password + salt
	// 4. Create user entity

	user, err := uc.repo.FindByEmail(ctx, dto.Email)

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
	if err := uc.repo.Create(ctx, userEntity); err != nil {
		return err
	}

	return nil
}

//1

type UserRepository interface {
	Create(ctx context.Context, data *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	//Update(ctx context.Context, data *domain.User) error
	//Find(ctx context.Context, id uuid.UUID) (*domain.User, error)
	//Delete(ctx context.Context, data *domain.User) error
}

type SessionRepository interface {
	Create(ctx context.Context, data *domain.Session) error
}
