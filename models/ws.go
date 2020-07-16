/*
@Time : 2020/7/16 13:04
@Author : xuyiqing
@File : ws.py
*/

package models

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
)

type Client struct {
	ID      uint64
	WsConn  *websocket.Conn
	Message chan map[string]interface{}
	Receive chan map[string]interface{}
}

type ClientManager struct {
	Clients map[uint64]*Client
}

var ClientManagerInstance = ClientManager{
	Clients: make(map[uint64]*Client, 1024),
}

// 获取消息，写入chan
func (c *Client) ReadMessage() {
	for {
		msg := make(map[string]interface{})
		err := c.WsConn.ReadJSON(&msg)
		if err != nil {  // err则表示断开连接，删除conn连接
			delete(ClientManagerInstance.Clients, c.ID)
			break
		}
		c.Receive <- msg
	}
}

// 从chan获取发送消息
func (c *Client) WriteMessage() {
	for {
		msg := <- c.Receive
		dstID := msg["id"]
		dstIDStr := dstID.(string)
		dstID64, _ := strconv.Atoi(dstIDStr)
		dstClient := ClientManagerInstance.Clients[uint64(dstID64)]
		if err := dstClient.WsConn.WriteJSON(msg); err != nil {
			fmt.Println(err)
		}
	}
}

type Message struct {
	BaseModel
	Text 		string 	`gorm:"comment:'文本'"`
	Sender 		Account `gorm:"foreignkey:SenderId"`
	SenderId	uint64	`gorm:"comment:'发送者'"`
	Receiver 	Account `gorm:"foreignkey:ReceiverId"`
	ReceiverId  uint64	`gorm:"comment:'接收者'"`
}
