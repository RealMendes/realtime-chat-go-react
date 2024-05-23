package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Configurando o upgrader do WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // Tamanho do buffer de leitura
	WriteBufferSize: 1024, // Tamanho do buffer de escrita

	// Verificando a origem da requisição
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Função para ler mensagens do WebSocket
func reader(conn *websocket.Conn) {
	for {
		// Lendo a mensagem do WebSocket
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Imprimindo a mensagem
		fmt.Println(string(p))

		// Escrevendo a mensagem de volta para o WebSocket
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}
	}
}

// Função para lidar com conexões WebSocket
func serveWs(w http.ResponseWriter, r *http.Request) {
	// Atualizando a conexão HTTP para uma conexão WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Imprimindo uma mensagem quando um cliente se conecta
	fmt.Println("Client Connected")

	// Chamando a função reader para lidar com mensagens do WebSocket
	reader(ws)
}

// Função para configurar as rotas
func setupRoutes() {
	// Rota raiz que retorna "Simple Server"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	// Rota WebSocket que chama a função serveWs
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
