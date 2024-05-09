package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Request struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func Login(name, pwd, addr string) {
	SERVER_ADDR = addr
	u := url.URL{Scheme: SCHEME, Host: SERVER_ADDR + ":" + PORT, Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("fatal:", err)
	}
	CONN = conn
	defer CONN.Close()

	createUserRequest := Request{
		Type: "create_user",
		Data: map[string]string{
			"username": name,
			"password": pwd,
		},
	}
	err = CONN.WriteJSON(createUserRequest)
	if err != nil {
		log.Fatal("fatal:", err)
	}
	go ListenResponse()
}

func ListenResponse() {
	for {
		_, message, err := CONN.ReadMessage()
		if err != nil {
			log.Fatal("fatal:", err)
		}

		fmt.Printf("收到服务器消息: %s\n", message)

		// 处理返回消息
		var response map[string]interface{}
		err = json.Unmarshal(message, &response)
		if err != nil {
			log.Println("解析返回消息错误:", err)
			continue
		}

		// 在这里处理服务器返回的响应
		fmt.Println("结果")
		fmt.Println(response)
	}
}

func MakeRequest() {
	// 连接到 WebSocket 服务器
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/ws"}
	fmt.Println("连接到服务器:", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("连接错误:", err)
	}
	defer conn.Close()

	// 创建用户请求
	createUserRequest := Request{
		Type: "create_user",
		Data: map[string]string{
			"username": "mtt",
			"password": "123123",
		},
	}

	// 发送创建用户请求
	err = conn.WriteJSON(createUserRequest)
	if err != nil {
		log.Fatal("发送请求错误:", err)
	}

	// 接收服务器返回的消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("读取消息错误:", err)
			return
		}

		fmt.Printf("收到服务器消息: %s\n", message)

		// 处理返回消息
		var response map[string]interface{}
		err = json.Unmarshal(message, &response)
		if err != nil {
			log.Println("解析返回消息错误:", err)
			continue
		}

		// 在这里处理服务器返回的响应
		fmt.Println("结果")
		fmt.Println(response)
	}
}
