package infras

import (
	"Ls04_GORM/common"
	"Ls04_GORM/module/category/query"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
)

type httpService struct {
	sctx sctx.ServiceContext
}

func NewHttpService(sctx sctx.ServiceContext) httpService {
	return httpService{sctx: sctx}
}

func (s httpService) handleRPCListCategories() gin.HandlerFunc {
	return func(c *gin.Context) {

		var param struct {
			Ids []uuid.UUID `json:"ids"`
		}

		if err := c.BindJSON(&param); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		result, err := query.NewCategoryById(s.sctx).Execute(c.Request.Context(), param.Ids)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(result))
	}
}

func (s httpService) Routes(g *gin.RouterGroup) {
	category := g.Group("category")

	rpc := category.Group("rpc")
	{
		rpc.GET("/query-categories-ids", s.handleRPCListCategories())
	}

}
