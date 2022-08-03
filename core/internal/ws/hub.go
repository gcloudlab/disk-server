package ws

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	clients map[*Client]bool // Registered clients. 上线clients

	broadcast chan []byte // Inbound messages from the clients. 客户端发送的消息 ->广播给其他的客户端

	register chan *Client // Register requests from the clients. 注册channel，接收注册msg

	unregister chan *Client // Unregister requests from clients. 下线channel
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		// 注册channel：存放到注册表中，数据流也就在这些client中发生
		case client := <-h.register:
			h.clients[client] = true
		// 下线channel：从注册表里面删除
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		// 广播消息：发送给注册表中的client中，send接收到并显示到client上
		case message := <-h.broadcast:
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
