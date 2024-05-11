package client

var (
	LOGIN_SIGNAL         = make(chan int)
	LIST_ROOM_SIGNAL     = make(chan int)
	JOIN_ROOM_SIGNAL     = make(chan int)
	CREATE_ROOM_SIGNAL   = make(chan int)
	SEND_ROOM_MSG_SIGNAL = make(chan int, 1)

	GAME_MSG_CHAN       = make(chan string, 1)
	HAND_CARD_CHAN      = make(chan handCardUpdate, 1)
	COMMUNITY_CARD_CHAN = make(chan communityCardUpdate, 1)
)
