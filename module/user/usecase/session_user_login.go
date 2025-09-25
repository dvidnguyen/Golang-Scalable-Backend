package usecase

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"
	"time"
)

type LoginUC struct {
	userRepo    UserQueryRepository
	sessionRepo SessionRepository

	hasher        Hasher
	tokenProvider TokenProvider
}

func NewLoginUC(userRepo UserQueryRepository, sessionRepo SessionRepository, hasher Hasher, tokenProvider TokenProvider) *LoginUC {
	return &LoginUC{userRepo: userRepo, sessionRepo: sessionRepo, hasher: hasher, tokenProvider: tokenProvider}
}

func (uc *LoginUC) Login(ctx context.Context, dto EmailPasswordLogin) (*TokenResponse, error) {
	// 1.Find user by email

	user, err := uc.userRepo.FindByEmail(ctx, dto.Email)

	if err != nil {
		return nil, err
	}

	if user.Status() == "banned" {
		return nil, domain.ErrUserBanned
	}
	count, err := uc.sessionRepo.CountSessionByUserId(ctx, user.Id())
	if err != nil {
		return nil, err
	}
	if count < 0 || count > 5 {
		return nil, domain.ErrTooManyLogin

	}
	// 2.hash and compare password with password login and salt
	if ok := uc.hasher.CompareHashPassword(user.Password(), user.Salt(), dto.Password); !ok {
		return nil, domain.ErrInvalidPassword
	}

	userID := user.Id()
	sessionID := common.GenUUID()

	// 3. Generate token
	accessToken, err := uc.tokenProvider.IssueToken(ctx, sessionID.String(), userID.String())

	if err != nil {
		return nil, err
	}

	// 4. Insert session in repo
	refreshToken, _ := uc.hasher.RandomStr(16)
	tokenExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.TokenExpireInSeconds()))
	refreshExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.RefreshExpireInSeconds()))
	session := domain.NewSession(sessionID, userID, refreshToken, tokenExpAt, refreshExpAt)

	if err := uc.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}
	// 5. return token

	return &TokenResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  uc.tokenProvider.TokenExpireInSeconds(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: uc.tokenProvider.RefreshExpireInSeconds(),
	}, nil

}
