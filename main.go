package main

import (
	"Ls04_GORM/builder"
	"Ls04_GORM/common"
	"Ls04_GORM/component"
	"Ls04_GORM/middleware"
	CatHTTTP "Ls04_GORM/module/category/infras"
	"Ls04_GORM/module/image"
	"Ls04_GORM/module/product/controller"
	"Ls04_GORM/module/product/domain/productusecase"
	ProductHTTP "Ls04_GORM/module/product/infras/http_service"
	"Ls04_GORM/module/product/productmysql"
	"Ls04_GORM/module/user/infras/httpservice"
	"Ls04_GORM/module/user/infras/repository"
	"Ls04_GORM/module/user/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/component/gormc"
)

func newServiceContext() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("GSB"),
		sctx.WithComponent(gormc.NewGormDB(common.KeyGorm, "")),
		sctx.WithComponent(component.NewJWT(common.KeyJWT)),
		sctx.WithComponent(component.NewAWSS3Provider(common.KeyAWSS3)),
	)
}

func main() {
	service := newServiceContext()

	if err := service.Load(); err != nil {
		log.Fatalln(err)
	}
	db := service.MustGet(common.KeyGorm).(common.DBContext).GetDB()

	r := gin.Default()
	r.Use(middleware.Recovery())

	tokenProvider := service.MustGet(common.KeyJWT).(component.TokenProvider)
	authClient := usecase.NewIntrospectUC(repository.NewUserRepository(db), repository.NewSessionRepository(db), tokenProvider)

	r.GET("/ping", middleware.RequireAuth(authClient), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Setup dependencies
	repo := productmysql.NewMysqlRepository(db)
	useCase := productusecase.NewCreateProductUseCase(repo)
	api := controller.NewAPIController(useCase)

	v1 := r.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", api.CreateProductAPI(db))
		}
	}

	//userUC := usecase.NewUseCase(repository.NewUserRepository(db), repository.NewSessionRepository(db), &common.Hasher{}, tokenProvider)
	userUseCase := usecase.UseCaseWithBuilder(builder.NewSimpleBuilder(db, tokenProvider))
	httpservice.NewService(userUseCase, service).SetAuthClient(authClient).Routes(v1)
	image.NewHTTPService(service).Routes(v1)
	ProductHTTP.NewHttpService(service).Routes(v1)
	CatHTTTP.NewHttpService(service).Routes(v1)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//type MockRepository struct{}
//
//func (r MockRepository) Create(ctx context.Context, data *domain.Session) error {
//	return nil
//}
