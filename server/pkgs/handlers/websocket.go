package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow connections from any origin.
		},
	}
)
var todoList []string

func getCmd(input string) string {
	inputArr := strings.Split(input, " ")
	return inputArr[0]
}

func getMessage(input string) string {
	inputArr := strings.Split(input, " ")
	var result string
	for i := 1; i < len(inputArr); i++ {
		result += inputArr[i]
	}
	return result
}

func updateTodoList(input string) {
	tmpList := todoList
	todoList = []string{}
	for _, val := range tmpList {
		if val == input {
			continue
		}
		todoList = append(todoList, val)
	}
}

func (h Handles) WebSocGin(c *gin.Context) {
	roomId := c.Param("room_id")
	fmt.Println("roomId", roomId)
	// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}
	defer conn.Close()

	// Continuosly read and write message
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read failed:", err)
			break
		}
		input := string(message)
		cmd := getCmd(input)
		msg := roomId + ">>" + getMessage(input)
		if cmd == "add" {
			todoList = append(todoList, msg)
		} else if cmd == "done" {
			updateTodoList(msg)
		}
		output := "Current Todos: \n"
		for _, todo := range todoList {
			output += "\n - " + todo + "\n"
		}
		output += "\n----------------------------------------"
		message = []byte(output)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write failed:", err)
			break
		}
	}
}

func (h Handles) WebSocPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}
