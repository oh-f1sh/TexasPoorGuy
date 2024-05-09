package client

import "github.com/gorilla/websocket"

var (
	DEFAULTUSERNAME = ""
	DEFAULTPWD      = ""

	SCHEME      = "ws"
	SERVER_ADDR = ""
	PATH        = "/ws"
	PORT        = "8080"
	CONN        = new(websocket.Conn)

	USERNAME = ""
	USERID   = 0
	TOKEN    = ""
)
