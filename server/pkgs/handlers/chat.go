package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	UserId    string
	Text      string
	Timestamp time.Time
}

type Client struct {
	Connection   *websocket.Conn
	Conversation []Message
}

type ChatSessions struct {
}

type PrivateChat struct {
	AuthUserId    int
	ChatUserId    int
	PrivateChatId int
}

var clients []Client
var conversation []Message

func parseMessage(input, userId string) Message {
	return Message{UserId: userId, Text: input, Timestamp: time.Now()}
}

func (h Handles) WebSocPrivate(c *gin.Context) {
	senderId, found := c.Params.Get("auth_user_id")
	if !found {
		log.Print("illegal ws connection request")
		return
	}
	fmt.Println("senderId", senderId)

	// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}
	defer conn.Close()

	clients = append(clients, Client{Connection: conn})

	// conversation := getSetConversation(ctx, userId)
	// h.Dao.RdClient.Set(ctx, "", )

	// output := fmt.Sprintf("%v joined the conversation", userId)
	// fmt.Println(output)
	// joinLog := []byte(output)
	// err = conn.WriteMessage(1, joinLog)
	// if err != nil {
	// 	log.Println("join log write failed:", err)
	// }

	// if conversation != nil && len(conversation) > 0 {
	// 	fmt.Println("old conversation: ", conversation)
	// 	message, _ := json.Marshal(conversation)
	// 	_ = conn.WriteMessage(1, message)
	// }

	// Continuously read and write message
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read failed:", err)
			break
		}
		if message == nil || string(message) == "" {
			fmt.Println("empty message received")
			continue
		}
		msg := parseMessage(string(message), senderId)
		conversation = append(conversation, msg)

		// output := fmt.Sprintf("Current conversation: %v", conversation)
		// message = []byte(output)

		//broadcast
		for _, client := range clients {
			message, _ = json.Marshal(conversation)
			err = client.Connection.WriteMessage(mt, message)
			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}

		// message, _ = json.Marshal(conversation)
		// err = conn.WriteMessage(mt, message)
		// if err != nil {
		// 	log.Println("write failed:", err)
		// 	break
		// }
	}

	fmt.Println(senderId, "left the conversation")
}

// func getSetConversation(ctx context.Context, userId int, conversation []string) {
// }
