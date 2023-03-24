package ws

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":3001", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Init() {

	// hub основная структура состояния, и может общаться между горутинами
	hub := newHub()
	go hub.run()

	fs := http.FileServer(http.Dir("/root/apps/chess/build/public"))
	http.Handle("/", fs)

	// отлавливаем через http, что мы хотим общаться сокетами
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler(hub, w, r)
	})

	// Запускаем http сервер
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(hub *Hub, w http.ResponseWriter, r *http.Request) {

	// разрешаем подключение со сторонних доменов
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// апгрейдим http соединение на TCP
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// поднимаем нового клиента
	// записываем туда соединение,  хаб, и буфер сообщений для клиента
	client := &Client{conn: conn, hub: hub, send: make(chan []byte, 256)}

	// в хаб ложим инфу о новом пользователе
	client.hub.register <- client

	// проход буфера сообщений, и отправка
	go client.write()

	// чтение сообщений от юзера
	go client.read()

}
