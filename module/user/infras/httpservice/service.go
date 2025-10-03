package httpservice

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
)

type service struct {
	uc usecase.UseCase
}

func NewService(uc usecase.UseCase) *service {
	return &service{uc: uc}
}
func (s *service) handleRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.EmailPasswordRegistration
		if err := c.BindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := s.uc.Register(c.Request.Context(), dto); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "register successfully"})
	}
}
func (s *service) handleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.EmailPasswordLogin
		if err := c.BindJSON(&dto); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}
		res, err := s.uc.Login(c.Request.Context(), dto)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": res})
	}
}
func (s *service) handleRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Data struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := c.BindJSON(&Data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data, err := s.uc.RefreshToken(c.Request.Context(), Data.RefreshToken)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": data})
	}
}

func (s *service) Routes(g *gin.RouterGroup) {
	g.POST("/register", s.handleRegister())
	g.POST("/authenticate", s.handleLogin())
	g.POST("/refresh-token", s.handleRefreshToken())
}
