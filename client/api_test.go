package client

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/oh-f1sh/TexasPoorGuy/common"
)

func TestHandleListRoomResp(t *testing.T) {
	message := []byte(`{
		"success": true,
		"message": "List of rooms",
		"data": [
		  {
			"id": 1,
			"owner": {
			  "ID": 114466,
			  "Username": "szy",
			  "Password": "123123",
			  "Balance": 0,
			  "chip": 10000
			},
			"players": {
			  "114466": {
				"ID": 114466,
				"Username": "szy",
				"Password": "123123",
				"Balance": 0,
				"chip": 10000
			  }
			},
			"player_list": [
			  {
				"ID": 114466,
				"Username": "szy",
				"Password": "123123",
				"Balance": 0,
				"chip": 10000
			  }
			],
			"type": "texas",
			"limit_chip": 10000,
			"sb_mount": 10,
			"chat_list": null
		  }
		],
		"type": "list_room"
	  }`)
	fmt.Println(string(message))
	var resp map[string]interface{}
	_ = json.Unmarshal(message, &resp)
	go HandleListRoomResp(resp)
	<-common.LIST_ROOM_SIGNAL
	fmt.Printf("%+v", ROOM_LIST)
}
