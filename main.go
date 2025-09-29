package main

import (
	"Ls04_GORM/builder"
	"Ls04_GORM/component"
	"Ls04_GORM/middleware"
	"Ls04_GORM/module/product/controller"
	"Ls04_GORM/module/product/productdomain/productusecase"
	"Ls04_GORM/module/product/productmysql"
	"Ls04_GORM/module/user/infras/httpservice"
	"Ls04_GORM/module/user/infras/repository"
	"Ls04_GORM/module/user/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.Use(middleware.Recovery())
	tokenProvider := component.NewJWTProvider("very-important-please-change-it!",
		60*60*24*7, 60*60*24*14)
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
	httpservice.NewService(userUseCase).Routes(v1)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//type MockRepository struct{}
//
//func (r MockRepository) Create(ctx context.Context, data *domain.Session) error {
//	return nil
//}
