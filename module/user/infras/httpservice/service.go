package httpservice

import (
	"Ls04_GORM/module/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": "register successfully"})
	}
}
func (s *service) Routes(g *gin.RouterGroup) {
	g.POST("/register", s.handleRegister())
}
