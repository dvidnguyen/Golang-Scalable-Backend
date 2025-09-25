package usecase

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"
	"errors"
	"time"
)

type refreshTokenUC struct {
	userRepo      UserQueryRepository
	sessionRepo   SessionRepository
	tokenProvider TokenProvider
	hasher        Hasher
}

func NewRefreshTokenUC(userRepo UserQueryRepository, sessionRepo SessionRepository, tokenProvider TokenProvider, hasher Hasher) *refreshTokenUC {
	return &refreshTokenUC{userRepo: userRepo, sessionRepo: sessionRepo, tokenProvider: tokenProvider, hasher: hasher}
}

func (uc *refreshTokenUC) RefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error) {
	// 1. Find session by refresh token
	session, err := uc.sessionRepo.FindByRefreshToken(ctx, refreshToken)

	if err != nil {
		return nil, err
	}
	// 2. Check refresh token expiration
	if session.RefreshExpAt().UnixNano() < time.Now().UTC().UnixNano() {
		return nil, errors.New("refresh token has expired")
	}
	user, err := uc.userRepo.FindById(ctx, session.UserId())

	if err != nil {
		return nil, err
	}

	if user.Status() == "banned" {
		return nil, errors.New("user has been banned")
	}

	userId := user.Id()
	newSessionId := common.GenUUID()

	// 3. Generate new token
	accessToken, err := uc.tokenProvider.IssueToken(ctx, newSessionId.String(), userId.String())

	if err != nil {
		return nil, err
	}
	newRefreshToken, _ := uc.hasher.RandomStr(16)
	tokenExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.TokenExpireInSeconds()))
	refreshExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.RefreshExpireInSeconds()))
	newSession := domain.NewSession(newSessionId, userId, newRefreshToken, tokenExpAt, refreshExpAt)

	if err := uc.sessionRepo.Create(ctx, newSession); err != nil {
		return nil, err
	}
	go func() {
		_ = uc.sessionRepo.Delete(ctx, session.Id())
	}()

	return &TokenResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  uc.tokenProvider.TokenExpireInSeconds(),
		RefreshToken:          newRefreshToken,
		RefreshTokenExpiresAt: uc.tokenProvider.RefreshExpireInSeconds(),
	}, nil
}
