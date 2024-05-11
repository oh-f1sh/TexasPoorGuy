package common

import (
	"os"

	"github.com/gorilla/websocket"
)

var (
	DEFAULTUSERNAME = ""
	DEFAULTPWD      = ""

	SCHEME      = "ws"
	SERVER_ADDR = "·"
	PATH        = "/ws"
	PORT        = "8080"
	CONN        = new(websocket.Conn)

	USERNAME    = ""
	USERID      = 0
	TOKEN       = ""
	ROOMID      = -1
	ROOMOWNERID = 0
	ROOM_CHAT   = ""

	LOG_FILE *os.File
)
