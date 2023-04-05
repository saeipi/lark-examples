package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
	"log"
	"time"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte, 100000),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			fmt.Println(utils.ToString(message))
			time.Sleep(1 * time.Second)
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) Upgrade(ctx *gin.Context) {
	var (
		uid         any
		uidVal      int64
		platform    any
		platformVal int32
		exists      bool
		err         error
	)
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	uid, exists = ctx.Get(constant.USER_UID)
	if exists == false {
		xhttp.Error(ctx, 1003, "获取用户ID失败")
		return
	}
	uidVal, err = utils.ToInt64(uid)
	if err != nil || uidVal == 0 {
		xhttp.Error(ctx, 1003, "获取用户ID失败")
		return
	}

	platform, exists = ctx.Get(constant.USER_PLATFORM)
	if exists == false {
		xhttp.Error(ctx, 1004, "获取用户平台信息失败")
		return
	}
	platformVal, err = utils.ToInt32(platform)
	if err != nil || uidVal == 0 {
		xhttp.Error(ctx, 1004, "获取用户平台信息失败")
		return
	}

	client := &Client{uid: uidVal, platform: platformVal, hub: h, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
