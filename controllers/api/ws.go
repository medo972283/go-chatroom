package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/medo972283/go-chatroom/models"
	"github.com/medo972283/go-chatroom/repositories"
	"github.com/medo972283/go-chatroom/utils"
	"golang.org/x/net/websocket"
)

// https://www.itread01.com/content/1546425722.html

type ClientManager struct {
	// record client's status
	clients map[*Client]bool
	// receive
	broadcast chan []byte
	// new client
	register chan *Client
	// client closed
	unregister chan *Client
}

type Client struct {
	// user ID
	id string
	// web socket connection
	socket *websocket.Conn
	// message get from client
	send chan []byte
	// room ID
	// roomID string
}

// type Message struct {
// 	Sender    string `json:"sender,omitempty"`
// 	Recipient string `json:"recipient,omitempty"`
// 	Content   string `json:"content,omitempty"`
// }

var Manager = ClientManager{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for client := range manager.clients {
		if client != ignore {
			client.send <- message
		}
	}
}

func (manager *ClientManager) Start() {
	for {
		select {
		// New connection
		case client := <-manager.register:
			manager.clients[client] = true
			fmt.Printf("\n- register client -\n")
		// Disconnection
		case client := <-manager.unregister:
			if _, ok := manager.clients[client]; ok {
				close(client.send)
				delete(manager.clients, client)
				fmt.Printf("\n- unregister client -\n")
				// TODO
				// repositories.DeleteParticipantByUk(session.Values["userID"],client.RoomID)
			}
		// Broadcast new message to each clients
		case message := <-manager.broadcast:
			for client := range manager.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(manager.clients, client)
					// TODO
					// repositories.DeleteParticipantByUk(session.Values["userID"],client.RoomID)
				}
			}
		}
	}
}

func (c *Client) read(session *sessions.Session) {
	defer func() {
		Manager.unregister <- c
		c.socket.Close()
	}()

	for {

		var message models.Message
		err := websocket.JSON.Receive(c.socket, &message)
		if err != nil {
			// Get error EOF when client terminate the web socket, e.g. refresh/close the page
			Manager.unregister <- c
			c.socket.Close()
			break
		}

		// Create new message
		message.ID = utils.ProduceV1UUID()
		message.UserID = utils.ConverseToUUID(session.Values["userID"])
		repositories.CreateMessage(&message)

		// TODO
		// repositories.CreateParticipant(&models.Participant{UserID: message.UerID ,RoomID: message.RoomID})

		// JsonEncode the message and broadcast to all clients
		jsonMessage, _ := json.Marshal(message)
		Manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message := <-c.send:
			// if !ok {
			// 	return
			// }

			// JsonDecode the message from the client's send channel
			var decodeMessage models.Message
			json.Unmarshal(message, &decodeMessage)

			// Sent response to the client
			sendMessage, _ := repositories.GetMessageByID(fmt.Sprintf("%v", decodeMessage.ID))
			websocket.JSON.Send(c.socket, sendMessage)
		}
	}
}

func MessageChannel(c echo.Context) error {

	websocket.Handler(func(conn *websocket.Conn) {
		defer conn.Close()

		client := &Client{
			id:     uuid.Must(uuid.NewRandom()).String(),
			socket: conn,
			send:   make(chan []byte)}

		// Register a new client
		Manager.register <- client

		session, _ := session.Get("session", c)

		// Keep receiving the data from web client
		go client.read(session)
		// Keep responsing the data to the web client
		go client.write()

		// Create a context without deadline to keep socket alive
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				return
			}
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
