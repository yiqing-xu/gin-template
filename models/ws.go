/*
@Time : 2020/7/16 13:04
@Author : xuyiqing
@File : ws.py
*/

package models

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type FormalMsg struct {
	ID 	 uint64 `json:"id"`
	Text string `json:"text"`
}

type Client struct {
	ID      uint64
	WsConn  *websocket.Conn
	Message chan map[string]interface{}
	Receive chan *FormalMsg
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
		var msg FormalMsg
		err := c.WsConn.ReadJSON(&msg)
		if err != nil {  // err则表示断开连接，删除conn连接
			fmt.Println(err)
			delete(ClientManagerInstance.Clients, c.ID)
			break
		}
		c.Receive <- &msg
	}
}

// 从chan获取发送消息
func (c *Client) WriteMessage() {
	for {
		msg := <- c.Receive
		go func() {
			message := Message{
				Text: msg.Text,
				SenderId: c.ID,
				ReceiverId: msg.ID,
			}
			DB.Create(&message)
		}()

		dstClient := ClientManagerInstance.Clients[msg.ID]
		if err := dstClient.WsConn.WriteJSON(msg); err != nil {
			fmt.Println(err)
		}
	}
}

type Message struct {
	BaseModel
	Text 		string 	`gorm:"comment:'文本'" json:"text"`
	Sender 		Account `gorm:"foreignkey:SenderId" json:"sender"`
	SenderId	uint64	`gorm:"comment:'发送者'" json:"-"`
	Receiver 	Account `gorm:"foreignkey:ReceiverId" json:"receiver"`
	ReceiverId  uint64	`gorm:"comment:'接收者'" json:"-"`
}

func (m *Message) TableName() string {
	return "messages"
}
