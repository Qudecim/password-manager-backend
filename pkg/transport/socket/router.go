package socket

import "log"

type MessageHandler func(message []byte)

type Router struct {
	handlers map[string]MessageHandler
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]MessageHandler),
	}
}

func (r *Router) AddHandler(messageType string, handler MessageHandler) {
	r.handlers[messageType] = handler
}

func (r *Router) HandleMessage(c *Client, message []byte) {
	// Извлекаем тип сообщения
	messageType := string(message)

	// Получаем соответствующий обработчик
	handler, ok := r.handlers[messageType]
	if !ok {
		// Если нет обработчика для данного типа сообщения, можно выполнить обработку по умолчанию
		handler = r.defaultHandler
	}

	// Вызываем обработчик
	handler(message)
}

func (r *Router) defaultHandler(message []byte) {
	// Обработка по умолчанию, если нет обработчика для данного типа сообщения
	log.Printf("Received message: %s", message)
}
