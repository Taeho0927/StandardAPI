package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/Taeho0927/Standard_API/conf"
	"github/Taeho0927/Standard_API/services"
	"time"
)

func JSONMiddleware() gin.HandlerFunc {
	front := conf.FrontDomain()
	fmt.Println("connection front to : ", front)
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", front)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, X-Requested-With, Content-Type, Access-Control-Request-Headers")
	}
}

func RunAPI() (*gin.Engine, error) {
	r := gin.Default()
	r.Use(JSONMiddleware())
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	h, err := services.NewHandler()
	if err != nil {
		return nil, err
	}

	FileParamGroup := r.Group("/v1")       // FileGroup 사용 시
	FileParamGroup.Use(h.BindingFileParam) // FileGroup이 Binding을 사용하도록

	mainGroup := r.Group("/v1")
	mainGroup.Use(h.CheckDBConnection)
	{
		// 메인 그룹 라우팅 추가
		// 하위 그룹 생성 후 그룹라우팅 선호
	}
	return r, nil
}
