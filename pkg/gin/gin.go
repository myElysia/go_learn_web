package gin

import (
	"go_learn_web/configs"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	"go_learn_web/logs"
	"net/http"
)

/**
swagger 接口文档信息参考 https://www.shouxicto.com/article/969.html
*/

// @Summary health
// @Description 项目健康检查接口，检查项目是否正常运行。(Author: henglong)
// @Tags 健康检查接口
// @Accept json
// @Produce json
// @Success 200 {string} json "{"message": "health"}"
// @Router /health [get]
func ginHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "health",
	})
	return
}

func Init() {
	r := gin.Default()
	gin.SetMode("release")
	// 使用日志中间件
	r.Use(logs.LoggerToFile())

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = "elysia.arknights.top"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		{
			health.GET("", ginHealth)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(configs.GIN_RUN_HOST + ":" + configs.GIN_RUN_PORT)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
