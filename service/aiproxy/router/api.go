package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/labring/sealos/service/aiproxy/common/env"
	"github.com/labring/sealos/service/aiproxy/controller"
	"github.com/labring/sealos/service/aiproxy/middleware"

	"github.com/gin-gonic/gin"
)

func SetAPIRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	if env.Bool("GZIP_ENABLED", false) {
		apiRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	apiRouter.Use(middleware.AdminAuth)
	{
		apiRouter.GET("/status", controller.GetStatus)
		apiRouter.GET("/models", controller.BuiltinModels)
		apiRouter.GET("/models/price", controller.ModelPrice)
		apiRouter.GET("/models/enabled", controller.EnabledModels)
		apiRouter.GET("/models/enabled/price", controller.EnabledModelsAndPrice)
		apiRouter.GET("/models/enabled/channel", controller.EnabledType2Models)
		apiRouter.GET("/models/enabled/channel/price", controller.EnabledType2ModelsAndPrice)
		apiRouter.GET("/models/enabled/default", controller.ChannelDefaultModels)
		apiRouter.GET("/models/enabled/default/:type", controller.ChannelDefaultModelsByType)

		groupsRoute := apiRouter.Group("/groups")
		{
			groupsRoute.GET("/", controller.GetGroups)
			groupsRoute.GET("/search", controller.SearchGroups)
		}
		groupRoute := apiRouter.Group("/group")
		{
			groupRoute.POST("/", controller.CreateGroup)
			groupRoute.GET("/:id", controller.GetGroup)
			groupRoute.DELETE("/:id", controller.DeleteGroup)
			groupRoute.POST("/:id/status", controller.UpdateGroupStatus)
			groupRoute.POST("/:id/qpm", controller.UpdateGroupQPM)
		}
		optionRoute := apiRouter.Group("/option")
		{
			optionRoute.GET("/", controller.GetOptions)
			optionRoute.PUT("/", controller.UpdateOption)
			optionRoute.PUT("/batch", controller.UpdateOptions)
		}
		channelsRoute := apiRouter.Group("/channels")
		{
			channelsRoute.GET("/", controller.GetChannels)
			channelsRoute.GET("/all", controller.GetAllChannels)
			channelsRoute.POST("/", controller.AddChannels)
			channelsRoute.GET("/search", controller.SearchChannels)
			channelsRoute.GET("/test", controller.TestChannels)
			channelsRoute.GET("/update_balance", controller.UpdateAllChannelsBalance)
		}
		channelRoute := apiRouter.Group("/channel")
		{
			channelRoute.GET("/:id", controller.GetChannel)
			channelRoute.POST("/", controller.AddChannel)
			channelRoute.PUT("/", controller.UpdateChannel)
			channelRoute.POST("/:id/status", controller.UpdateChannelStatus)
			channelRoute.DELETE("/:id", controller.DeleteChannel)
			channelRoute.GET("/test/:id", controller.TestChannel)
			channelRoute.GET("/update_balance/:id", controller.UpdateChannelBalance)
		}
		tokensRoute := apiRouter.Group("/tokens")
		{
			tokensRoute.GET("/", controller.GetTokens)
			tokensRoute.GET("/:id", controller.GetToken)
			tokensRoute.PUT("/:id", controller.UpdateToken)
			tokensRoute.POST("/:id/status", controller.UpdateTokenStatus)
			tokensRoute.POST("/:id/name", controller.UpdateTokenName)
			tokensRoute.DELETE("/:id", controller.DeleteToken)
			tokensRoute.GET("/search", controller.SearchTokens)
		}
		tokenRoute := apiRouter.Group("/token")
		{
			tokenRoute.GET("/:group/search", controller.SearchGroupTokens)
			tokenRoute.GET("/:group", controller.GetGroupTokens)
			tokenRoute.GET("/:group/:id", controller.GetGroupToken)
			tokenRoute.POST("/:group", controller.AddToken)
			tokenRoute.PUT("/:group/:id", controller.UpdateGroupToken)
			tokenRoute.POST("/:group/:id/status", controller.UpdateGroupTokenStatus)
			tokenRoute.POST("/:group/:id/name", controller.UpdateGroupTokenName)
			tokenRoute.DELETE("/:group/:id", controller.DeleteGroupToken)
		}
		logsRoute := apiRouter.Group("/logs")
		{
			logsRoute.GET("/", controller.GetLogs)
			logsRoute.DELETE("/", controller.DeleteHistoryLogs)
			logsRoute.GET("/stat", controller.GetLogsStat)
			logsRoute.GET("/search", controller.SearchLogs)
			logsRoute.GET("/consume_error", controller.SearchConsumeError)
		}
		logRoute := apiRouter.Group("/log")
		{
			logRoute.GET("/:group/search", controller.SearchGroupLogs)
			logRoute.GET("/:group", controller.GetGroupLogs)
		}
	}
}
