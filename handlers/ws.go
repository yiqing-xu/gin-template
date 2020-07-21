/*
@Time : 2020/7/16 12:40
@Author : xuyiqing
@File : ws.py
*/

package handlers

import (
	"gin-template/models"
	"gin-template/pkg/jwt"
	"gin-template/serializers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsMessageHandler(ctx *gin.Context) {
	token := ctx.Query("token")
	claims, err := jwt.ValidateJwtToken(token)
	if err != nil {
		return
	}
	currentUser := claims.GetUserByID()
	wsConn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		panic(err)
	}
	client := models.Client{
		ID: currentUser.ID,
		WsConn: wsConn,
		Message: make(chan *models.FormalMsg, 1024),
		Receive: make(chan *models.FormalMsg, 1024),
	}
	models.ClientManagerInstance.Clients[client.ID] = &client
	go client.ReadMessage()
	go client.WriteMessage()
	go client.SendMessage()
}

func SendWsMessageHandler(user models.Account, text string) {
	if dstClient, ok := models.ClientManagerInstance.Clients[user.ID]; ok {
		dstClient.Message <- &models.FormalMsg{
			ID:   user.ID,
			Text: text,
		}
	}
}

// 查询消息记录
func GetWsMessageHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	currentUser := jwt.AssertUser(ctx)
	dstUserId := ctx.Param("id")
	var pager serializers.Pager
	pager.InitPager(ctx)
	if currentUser == nil {
		response.Unauthenticated("未认证")
		return
	}
	var messages []models.Message
	db := models.DB.Model(&messages).Limit(pager.PageSize).Where("sender_id = ?",
		currentUser.ID).Or("receiver_id = ?", dstUserId).Preload("Sender").Preload("Receiver")
	db.Count(&pager.Total)
	db.Offset(pager.OffSet).Find(&messages)
	pager.GetPager()
	response.Response(messages, pager)
}
