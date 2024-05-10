package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
	"github.com/oh-f1sh/TexasPoorGuy/common"
)

type Request struct {
	Type  string      `json:"type"`
	Data  interface{} `json:"data"`
	Token string      `json:"token,omitempty"`
}

func ListRoom() {
	req := Request{
		Type:  "list_room",
		Data:  map[string]string{},
		Token: common.TOKEN,
	}
	err := common.CONN.WriteJSON(req)
	if err != nil {
		log.Fatal("fatal:", err)
	}
}

func CreateRoom() {
	req := Request{
		Type: "create_room",
		Data: map[string]interface{}{
			"room_type": "texas",
			"chip":      10000,
		},
		Token: common.TOKEN,
	}
	err := common.CONN.WriteJSON(req)
	if err != nil {
		log.Fatal("fatal:", err)
	}
}

func JoinRoom(id int) {
	req := Request{
		Type: "join_room",
		Data: map[string]interface{}{
			"room_id": id,
		},
		Token: common.TOKEN,
	}
	err := common.CONN.WriteJSON(req)
	if err != nil {
		log.Fatal("fatal:", err)
	}
}

func StartGame() {
	req := Request{
		Type:  "start_game",
		Data:  map[string]interface{}{},
		Token: common.TOKEN,
	}
	err := common.CONN.WriteJSON(req)
	if err != nil {
		log.Fatal("fatal:", err)
	}
}

func SendMsg(id int, msg string) {
	req := Request{
		Type: "send_room_msg",
		Data: map[string]interface{}{
			"room_id": 1,
			"msg":     msg,
		},
		Token: common.TOKEN,
	}
	err := common.CONN.WriteJSON(req)
	if err != nil {
		log.Fatal("fatal:", err)
	}
}

func Login(name, pwd, addr string) {
	common.SERVER_ADDR = addr
	common.USERNAME = name
	common.LOG_FILE, _ = tea.LogToFile("debug.log", common.USERNAME)
	u := url.URL{Scheme: common.SCHEME, Host: common.SERVER_ADDR + ":" + common.PORT, Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("fatal:", err)
	}
	common.CONN = conn

	createUserRequest := Request{
		Type: "login",
		Data: map[string]string{
			"username": name,
			"password": pwd,
		},
	}
	err = common.CONN.WriteJSON(createUserRequest)
	if err != nil {
		log.Fatal("fatal:", err)
	}
	go ListenResponse()
}

func ListenResponse() {
	defer common.CONN.Close()
	defer common.LOG_FILE.Close()
	for {
		_, message, err := common.CONN.ReadMessage()
		if err != nil {
			log.Fatal("fatal:", err)
		}

		var resp map[string]interface{}
		err = json.Unmarshal(message, &resp)
		if err != nil {
			common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 服务器返回错误：%v\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID, err))
			continue
		}
		common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 服务器已返回操作，操作类型：%v\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID, resp["type"]))
		switch resp["type"] {
		case "login":
			HandleLoginResp(resp)
		case "list_room":
			HandleListRoomResp(resp)
		case "join_room":
			HandleJoinRoomResp(resp)
		case "create_room":
			HandleCreateRoomResp(resp)
		case "send_room_msg":
			HandleSendRoomMsgResp(resp)
		case "game_info":
			HandleGameInfoResp(resp)
		case "list_room_player":

		case "quit_room":

		case "get_subsidy":

		case "room_user_join":

		}
	}
}

func HandleLoginResp(resp map[string]interface{}) {
	tk, _ := resp["data"].(map[string]interface{})["token"].(string)
	userID, _ := resp["data"].(map[string]interface{})["user_id"].(int)
	common.TOKEN = tk
	common.USERID = userID
	common.LOGIN_SIGNAL <- 1
}

func HandleListRoomResp(resp map[string]interface{}) {
	ROOM_LIST = make([]Room, 0)
	roomList := resp["data"].([]interface{})
	for _, roomObj := range roomList {
		room := roomObj.(map[string]interface{})
		id := int(room["id"].(float64))
		owner := room["owner"].(map[string]interface{})
		playerList := room["player_list"].([]interface{})
		typ := room["type"].(string)
		ROOM_LIST = append(ROOM_LIST, Room{
			id:          id,
			ownerID:     int(owner["ID"].(float64)),
			ownerName:   owner["Username"].(string),
			gameType:    typ,
			playerCount: len(playerList),
		})
	}
	common.LIST_ROOM_SIGNAL <- 1
}

func HandleCreateRoomResp(resp map[string]interface{}) {
	data := resp["data"].(map[string]interface{})
	roomID := int(data["id"].(float64))
	common.ROOMID = roomID
	common.ROOMOWNERID = common.USERID
	common.CREATE_ROOM_SIGNAL <- 1
}

func HandleJoinRoomResp(resp map[string]interface{}) {
	data := resp["data"].(map[string]interface{})
	var roomID string
	for k := range data {
		roomID = k
	}
	common.ROOMID = int(data[roomID].(map[string]interface{})["id"].(float64))
	common.ROOMOWNERID = int(data[roomID].(map[string]interface{})["owner"].(map[string]interface{})["ID"].(float64))
	common.JOIN_ROOM_SIGNAL <- 1
}

func HandleSendRoomMsgResp(resp map[string]interface{}) {
	data := resp["data"].([]interface{})
	POOR_GUY_CLIENT.ChatView.usernames = []string{}
	POOR_GUY_CLIENT.ChatView.messages = []string{}
	for _, v := range data {
		msg := v.(map[string]interface{})
		POOR_GUY_CLIENT.ChatView.usernames = append(POOR_GUY_CLIENT.ChatView.messages, msg["user_name"].(string))
		POOR_GUY_CLIENT.ChatView.messages = append(POOR_GUY_CLIENT.ChatView.messages, msg["message"].(string))
	}
	POOR_GUY_CLIENT.ChatView.Update(nil)
}

func HandleGameInfoResp(resp map[string]interface{}) {
	fmt.Println("\n\n\n\n\n\n\n\n\n", resp)
	data := resp["data"].(map[string]interface{})
	game := data["game"].(map[string]interface{})
	game_gameInfo := game["game_info"].(map[string]interface{})
	player := data["game"].(map[string]interface{})

	// update game msg
	gameMsg := resp["message"].(string)
	POOR_GUY_CLIENT.GameView.messages = append(POOR_GUY_CLIENT.GameView.messages, gameMsg)
	POOR_GUY_CLIENT.GameView.Update(nil)

	// update hand card
	if player["hand"] != nil {
		cards := player["hand"].([]interface{})
		POOR_GUY_CLIENT.HandCardView.card = []string{
			cards[0].(map[string]interface{})["desc"].(string) + "\n" + suitMap[int(cards[0].(map[string]interface{})["suit"].(float64))],
			cards[1].(map[string]interface{})["desc"].(string) + "\n" + suitMap[int(cards[1].(map[string]interface{})["suit"].(float64))],
		}
		POOR_GUY_CLIENT.HandCardView.color = []string{
			suitColor[int(cards[0].(map[string]interface{})["suit"].(float64))],
			suitColor[int(cards[1].(map[string]interface{})["suit"].(float64))],
		}
		POOR_GUY_CLIENT.HandCardView.background = []string{
			white,
			white,
		}
	} else {
		POOR_GUY_CLIENT.HandCardView.card = []string{
			"",
			"",
		}
		POOR_GUY_CLIENT.HandCardView.color = []string{
			black,
			black,
		}
		POOR_GUY_CLIENT.HandCardView.background = []string{
			darkgrey,
			darkgrey,
		}
	}
	POOR_GUY_CLIENT.HandCardView.Update(nil)

	// update community card
	if game_gameInfo["board"] != nil {
		for i, card := range game_gameInfo["board"].([]interface{}) {
			c := card.(map[string]interface{})
			POOR_GUY_CLIENT.CardView.card[i] = c["desc"].(string) + "\n" + suitMap[int(c["suit"].(float64))]
			POOR_GUY_CLIENT.CardView.color[i] = suitColor[int(c["suit"].(float64))]
			POOR_GUY_CLIENT.CardView.background[i] = white
		}
	} else {
		for i := range POOR_GUY_CLIENT.CardView.card {
			POOR_GUY_CLIENT.CardView.card[i] = ""
			POOR_GUY_CLIENT.CardView.color[i] = black
			POOR_GUY_CLIENT.CardView.background[i] = darkgrey
		}
	}
	POOR_GUY_CLIENT.CardView.Update(nil)

	// update scoreboard

	// update control area

}
