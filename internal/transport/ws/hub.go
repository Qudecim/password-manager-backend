package ws

type Hub struct {
	// Список активных клиентов
	clients map[*Client]bool

	// Полученные сообщения отк клиентов, которые еще не расскиданы на отправку
	broadcast chan []byte

	// Новые клиенты
	register chan *Client

	// Клиенты которые уходят
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {

		// обрабатываем новых пользователей
		case client := <-h.register:
			h.clients[client] = true

		// обрабатываем пользователей которые дисконектнулись
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		// берем сообщения из буфера в хабе и ставим очередь для отправки пользователю
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:

				// В select case не нужно ставить брейк, поэтому если не удалось поставить на отправку, то дисконектим пользователя
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
