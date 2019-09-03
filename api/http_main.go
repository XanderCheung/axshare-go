package api

import (
	"axshare_go/api/v1/attachment"
	"axshare_go/api/v1/axure"
	"axshare_go/api/v1/axure_group"
	app "axshare_go/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
)

func HttpServerRun() {
	// 设置日志文件
	f, _ := os.OpenFile("log/axshare_go_http.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	gin.DefaultWriter = io.MultiWriter(f)

	// 定义路由
	router := gin.Default()

	router.Use(app.TokenAuthMiddleware())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/axure_groups", axure_group.FetchList)
		v1.GET("/axure_groups/:axure_group_id/axures", axure.GetAxures)
		v1.GET("/axure_groups/:axure_group_id/axures/:axure_id/attachments", attachment.GetAttachments)
	}
	port := viper.GetString("http.port")
	_ = router.Run(":" + port)

}
