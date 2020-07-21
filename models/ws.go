/*
@Time : 2020/7/16 13:04
@Author : xuyiqing
@File : ws.py
*/

package models

import "C"
import (
	"fmt"
	"github.com/gorilla/websocket"
)

type FormalMsg struct {
	ID 	 uint64 `json:"id"`
	Text  string `json:"text"`
}

type Client struct {
	ID      uint64
	WsConn  *websocket.Conn
	Receive chan *FormalMsg
	Message chan *FormalMsg
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

		if err := c.WsConn.WriteJSON(msg); err != nil {
			fmt.Println(err)
		}
	}
}

// 单向通知消息
func (c *Client) SendMessage() {
	for {
		msg := <- c.Message
		if err := c.WsConn.WriteJSON(msg); err != nil {
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
