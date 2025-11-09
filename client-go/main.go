package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pebbe/zmq4"
)

func sendRequest(req map[string]interface{}) (map[string]interface{}, error) {
	// 1. VERIFICAÇÃO ESSENCIAL: Verifica se a criação do socket falhou
	socket, err := zmq4.NewSocket(zmq4.REQ)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar socket ZeroMQ: %w", err)
	}
	defer socket.Close() // Agora é seguro chamar Close, pois 'socket' não é nil

	// 2. CORREÇÃO DE ATRIBUIÇÃO: socket.Connect retorna apenas um int (resultado C)
	// Se houver um erro, o zmq4.SendBytes ou RecvBytes o reportará.
	socket.Connect("tcp://server:5555")
	// ^ Removemos: _, err = socket.Connect(...)

	b, _ := json.Marshal(req)

	// O primeiro erro real de rede/conexão deve ser capturado aqui:
	_, err = socket.SendBytes(b, 0)
	if err != nil {
		return nil, fmt.Errorf("falha ao enviar dados: %w", err)
	}

	replyBytes, err := socket.RecvBytes(0)
	if err != nil {
		return nil, fmt.Errorf("falha ao receber resposta: %w", err)
	}

	var reply map[string]interface{}
	json.Unmarshal(replyBytes, &reply)
	return reply, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: client-go <username>")
		return
	}
	user := os.Args[1]
	req := map[string]interface{}{
		"service": "login",
		"data": map[string]interface{}{
			"user":      user,
			"timestamp": time.Now().Unix(),
		},
	}
	fmt.Println("Enviando login:", user)
	r, err := sendRequest(req)
	if err != nil {
		// Se houver um erro, o programa imprime a mensagem de erro detalhada
		fmt.Println("erro no login:", err)
		return
	}
	fmt.Println("reply login:", r)

	// pedir lista de usuários
	req2 := map[string]interface{}{
		"service": "users",
		"data":    map[string]interface{}{"timestamp": time.Now().Unix()},
	}
	r2, err := sendRequest(req2)
	if err != nil {
		fmt.Println("erro ao pedir usuários:", err)
		return
	}
	fmt.Println("users reply:", r2)
}
