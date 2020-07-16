/*
@Time : 2020/7/16 12:40
@Author : xuyiqing
@File : ws.py
*/

package handlers

import (
	"fmt"
	"gin-template/models"
	"gin-template/pkg/jwt"
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
		Message: make(chan map[string]interface{}, 1024),
		Receive: make(chan map[string]interface{}, 1024),
	}
	models.ClientManagerInstance.Clients[client.ID] = &client
	go client.ReadMessage()
	go client.WriteMessage()
}

func TestWsMessageHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	currentUser := jwt.AssertUser(ctx)
	fmt.Println(currentUser)
	if client, ok := models.ClientManagerInstance.Clients[currentUser.ID]; ok {
		client.Message <- map[string]interface{}{
			"id": uint64(310974693714692804),
			"msg": "okokokok",
		}
	}
	response.Response(nil)
}

func Send()  {
	
}