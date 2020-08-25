/*
@Time : 2020/6/28 21:29
@Author : xuyiqing
@File : init.go
*/

package routers

import (
	"encoding/json"
	"gin-template/conf"
	"gin-template/middlewares"
	"github.com/gin-gonic/gin"
)

type UrlGroup func(group *gin.RouterGroup)

func InitRouter(isErrMsg, isDebugMode bool) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.ErrorHandleMiddleware(isErrMsg))

	if isDebugMode {
		var temp map[string]string
		if err := json.Unmarshal([]byte(conf.ProjectCfg.StaticUrlMapPath), &temp); err == nil {
			for url, path := range temp {
				router.Static(url, path)
			}
		}
	}
	router.LoadHTMLGlob(conf.ProjectCfg.TemplateGlob)

	routerGroupWithNoAuth := router.Group("api")  // 无需token验证路由
	RegisterUsersRouter(routerGroupWithNoAuth)
	RegisterSwaggerRouter(routerGroupWithNoAuth)
	RegisterWsHandler(routerGroupWithNoAuth)

	routerGroupWithAuth := router.Group("api/v1")  // token验证路由
	routerGroupWithAuth.Use(middlewares.AuthJwtTokenMiddleware())
	RegisterUsersRouterWithAuth(routerGroupWithAuth)
	RegisterCmsRouter(routerGroupWithAuth)
	RegisterWsAuthHandler(routerGroupWithAuth)

	return router
}
