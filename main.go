package main

import (
	"Ls04_GORM/common"
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

	r.GET("/ping", func(c *gin.Context) {
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

	userUC := usecase.NewUseCase(repository.NewUserRepository(db), &common.Hasher{})
	httpservice.NewService(userUC).Routes(v1)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
