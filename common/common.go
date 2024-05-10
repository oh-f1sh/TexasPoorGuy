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

	LOGIN_SIGNAL         = make(chan int)
	LIST_ROOM_SIGNAL     = make(chan int)
	JOIN_ROOM_SIGNAL     = make(chan int)
	CREATE_ROOM_SIGNAL   = make(chan int)
	SEND_ROOM_MSG_SIGNAL = make(chan int)

	LOG_FILE *os.File
)
