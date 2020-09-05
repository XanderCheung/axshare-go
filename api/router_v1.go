package api

import (
	"axshare_go/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/acct"
)

func SetV1Router(router *gin.Engine) {
	r := router.Group("/api/v1")
	{
		r.POST("/sign_in", acct.Handler.SignIn)

		r.GET("/axures/:id", v1.GetAxureWebInfo)

		r.Use(acct.TokenAuthMiddleware())

		accounts := r.Group("/accounts")
		{
			accounts.GET("/", acct.Handler.FetchAccounts)
			accounts.POST("/", acct.Handler.CreateAccount)
			accounts.GET("/:id", acct.Handler.FetchAccount)
			accounts.POST("/:id", acct.Handler.UpdateAccount)
			accounts.DELETE("/:id", acct.Handler.DestroyAccount)
			accounts.POST("/:id/password", acct.Handler.UpdateAccountPassword)
		}

		r.GET("/account/info", acct.Handler.FetchCurrentAccountInfo)

		r.GET("/axure_groups", v1.GetAxureGroups)

		r.GET("/axure_groups/:axure_group_id/axures", v1.GetAxures)
		r.GET("/axure_groups/:axure_group_id/axure/:id", v1.GetAxure)
		r.POST("/axure_groups/:axure_group_id/axures", v1.CreateAxure)
		r.PUT("/axure_groups/:axure_group_id/axure/:id", v1.UpdateAxure)

		r.GET("/axure_groups/:axure_group_id/axures/:axure_id/attachments", v1.GetAttachments)

		r.GET("/attachments", v1.GetAllAttachments)

		r.GET("/upload/token", v1.CreateUploadToken)
	}
}
