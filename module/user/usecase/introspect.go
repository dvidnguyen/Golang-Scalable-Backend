package usecase

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/domain"
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenPaser interface {
	ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error)
}

type IntrospectUC struct {
	userQueryRepo    UserQueryRepository
	sessionQueryRepo SessionQueryRepository
	tokenPaser       TokenPaser
}

func NewIntrospectUC(userQueryRepo UserQueryRepository, sessionQueryRepo SessionQueryRepository, tokenPaser TokenPaser) *IntrospectUC {
	return &IntrospectUC{
		userQueryRepo:    userQueryRepo,
		sessionQueryRepo: sessionQueryRepo,
		tokenPaser:       tokenPaser,
	}
}
func (uc *IntrospectUC) IntrospectToken(ctx context.Context, accessToken string) (common.Requester, error) {

	claims, err := uc.tokenPaser.ParseToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	userId := uuid.MustParse(claims.Subject)
	sessionId := uuid.MustParse(claims.ID)
	// check session
	if _, err := uc.sessionQueryRepo.Find(ctx, sessionId.String()); err != nil {
		return nil, err
	}

	user, err := uc.userQueryRepo.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if user.Status() == "banned" {
		return nil, domain.ErrUserBanned
	}
	return common.NewRequester(userId, sessionId, user.FirstName(), user.LastName(), user.Role().String(), user.Status()), nil
}
