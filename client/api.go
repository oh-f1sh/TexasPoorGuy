package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
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
			"room_id": id,
			"msg":     msg,
		},
		Token: common.TOKEN,
	}
	err := common.CONN.WriteJSON(req)
	if err != nil {
		log.Fatal("fatal:", err)
	}
}

func UserAction(action string, amount int) {
	req := Request{
		Type: "user_action",
		Data: map[string]interface{}{
			"action": action,
			"amount": amount,
		},
		Token: common.TOKEN,
	}
	common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 正在将操作发送给服务器，操作内容为:%v\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID, req))
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
		_ = json.Unmarshal(message, &resp)
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
		case "room_chat":
			HandleRoomChatResp(resp)
		case "game_info":
			HandleGameInfoResp(resp)
		case "game_info_error":
			HandleGameInfoErrorResp(resp)
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
	LOGIN_SIGNAL <- 1
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
	LIST_ROOM_SIGNAL <- 1
}

func HandleCreateRoomResp(resp map[string]interface{}) {
	data := resp["data"].(map[string]interface{})
	roomID := int(data["id"].(float64))
	common.ROOMID = roomID
	common.ROOMOWNERID = common.USERID
	CREATE_ROOM_SIGNAL <- 1
}

func HandleJoinRoomResp(resp map[string]interface{}) {
	data := resp["data"].(map[string]interface{})
	var roomID string
	for k := range data {
		roomID = k
	}
	common.ROOMID = int(data[roomID].(map[string]interface{})["id"].(float64))
	common.ROOMOWNERID = int(data[roomID].(map[string]interface{})["owner"].(map[string]interface{})["ID"].(float64))
}

func HandleRoomChatResp(resp map[string]interface{}) {
	data := resp["data"].([]interface{})
	msgs := []string{}
	names := []string{}
	for _, v := range data {
		msg := v.(map[string]interface{})
		names = append(names, msg["user_name"].(string))
		msgs = append(msgs, msg["message"].(string))
	}
	content := make([]string, 0)
	for i := range msgs {
		if names[i] == common.USERNAME {
			content = append(content, POOR_GUY_CLIENT.ChatView.senderStyle.Render(names[i])+": "+msgs[i])
		} else {
			content = append(content, POOR_GUY_CLIENT.ChatView.otherSenderStyle.Render(names[i])+": "+msgs[i])
		}
	}
	common.ROOM_CHAT = strings.Join(content, "\n")
	SEND_ROOM_MSG_SIGNAL <- 1
}

func HandleGameInfoResp(resp map[string]interface{}) {
	common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 开始处理游戏对局信息\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID))
	data := resp["data"].(map[string]interface{})
	game := data["game"].(map[string]interface{})
	game_gameInfo := game["game_info"].(map[string]interface{})
	player := data["player"].(map[string]interface{})

	// update game msg
	common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 游戏对局信息更新: %v\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID, resp["message"].(string)))
	GAME_MSG_CHAN <- resp["message"].(string)

	// update hand card
	hcUpdate := handCardUpdate{
		card:  []string{"", ""},
		color: []string{black, black},
		bg:    []string{darkgrey, darkgrey},
	}
	if player["hand"] != nil {
		cards := player["hand"].([]interface{})
		if len(cards) == 2 {
			hcUpdate.card[card1] = cards[card1].(map[string]interface{})["desc"].(string) + "\n" + suitMap[int(cards[card1].(map[string]interface{})["suit"].(float64))]
			hcUpdate.card[card2] = cards[card2].(map[string]interface{})["desc"].(string) + "\n" + suitMap[int(cards[card2].(map[string]interface{})["suit"].(float64))]
			hcUpdate.color[card1] = suitColor[int(cards[card1].(map[string]interface{})["suit"].(float64))]
			hcUpdate.color[card2] = suitColor[int(cards[card2].(map[string]interface{})["suit"].(float64))]
			hcUpdate.bg = []string{white, white}
		}
	}
	common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 手牌更新: %+v\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID, hcUpdate))
	HAND_CARD_CHAN <- hcUpdate

	// update community card
	ccUpdate := communityCardUpdate{
		card:  []string{"", "", "", "", ""},
		color: []string{black, black, black, black, black},
		bg:    []string{darkgrey, darkgrey, darkgrey, darkgrey, darkgrey},
	}
	if game_gameInfo["board"] != nil {
		for i, card := range game_gameInfo["board"].([]interface{}) {
			c := card.(map[string]interface{})
			POOR_GUY_CLIENT.CardView.card[i] = c["desc"].(string) + "\n" + suitMap[int(c["suit"].(float64))]
			POOR_GUY_CLIENT.CardView.color[i] = suitColor[int(c["suit"].(float64))]
			POOR_GUY_CLIENT.CardView.background[i] = white
		}
	}
	common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 牌河更新: %+v\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID, ccUpdate))
	COMMUNITY_CARD_CHAN <- ccUpdate

	// update scoreboard
	scoreUpdate := scoreboardUpdate{scores: []string{}}
	game_players := game["players"].([]interface{})
	for i, player := range game_players {
		p := player.(map[string]interface{})
		s := strings.Builder{}
		if i == int(game_gameInfo["action_player"].(float64)) {
			s.WriteString(waitingPlayerStyle.Render(strconv.Itoa(i+1) + "."))
			s.WriteString(waitingPlayerNameStyle.Render(p["player"].(map[string]interface{})["Username"].(string) + ":"))
			s.WriteString(strconv.Itoa(int(p["player"].(map[string]interface{})["chip"].(float64))) + ", ")
			s.WriteString("bets " + strconv.Itoa(int(p["bets"].(float64))) + ", ")
			s.WriteString("bets " + strconv.Itoa(int(p["bets"].(float64))) + ", ")
			if p["folded"].(bool) {
				s.WriteString("state: folded")
			} else if p["allIn"].(bool) {
				s.WriteString("state: all in")
			} else {
				s.WriteString("state: betting")
			}
			scoreUpdate.scores = append(scoreUpdate.scores, s.String())
		} else {
			s.WriteString(playingPlayerStyle.Render(strconv.Itoa(i+1) + "."))
			s.WriteString(playingPlayerNameStyle.Render(p["player"].(map[string]interface{})["Username"].(string) + ":"))
			s.WriteString(strconv.Itoa(int(p["player"].(map[string]interface{})["chip"].(float64))) + ", ")
			s.WriteString("bets " + strconv.Itoa(int(p["bets"].(float64))) + ", ")
			s.WriteString("bets " + strconv.Itoa(int(p["bets"].(float64))) + ", ")
			if p["folded"].(bool) {
				s.WriteString("state: folded")
			} else if p["allIn"].(bool) {
				s.WriteString("state: all in")
			} else {
				s.WriteString("state: betting")
			}
			scoreUpdate.scores = append(scoreUpdate.scores, s.String())
		}
	}
	scoreUpdate.scores = append(scoreUpdate.scores, "Pot: "+strconv.Itoa(int(game_gameInfo["pot"].(float64))))
	common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 下注信息更新: %+v\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID, scoreUpdate))
	SCOREBOARD_CHAN <- scoreUpdate

	// update control area

	common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 游戏对局信息已全部更新完毕\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID))
}

func HandleGameInfoErrorResp(resp map[string]interface{}) {
	// update game error msg
	common.LOG_FILE.WriteString(fmt.Sprintf("Now: %v, 当前用户:%v, %v, 游戏操作有误，返回错误提示: %v\n", time.Now().Format(time.RFC3339Nano), common.USERNAME, common.USERID, resp["message"].(string)))
	GAME_MSG_CHAN <- resp["message"].(string)
}
