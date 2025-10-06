package usecase

import (
	"Ls04_GORM/common"

	"github.com/google/uuid"
)

type EmailPasswordRegistration struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
type EmailPasswordLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken           string `json:"access_token"`
	AccessTokenExpiresAt  int    `json:"access_exp_at"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt int    `json:"refresh_exp_at"`
}
type SingleImgDTO struct {
	Requester common.Requester `json:"-"`
	ImageID   uuid.UUID        `json:"image_id"`
}
