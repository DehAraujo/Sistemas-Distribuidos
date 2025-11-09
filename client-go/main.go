package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pebbe/zmq4"
)

func sendRequest(req map[string]interface{}) (map[string]interface{}, error) {
	socket, _ := zmq4.NewSocket(zmq4.REQ)
	defer socket.Close()
	socket.Connect("tcp://server:5555") // in docker-compose we'll name service 'server'
	b, _ := json.Marshal(req)
	_, err := socket.SendBytes(b, 0)
	if err != nil {
		return nil, err
	}
	replyBytes, err := socket.RecvBytes(0)
	if err != nil {
		return nil, err
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
		fmt.Println("erro:", err)
		return
	}
	fmt.Println("reply:", r)

	// pedir lista de usuários
	req2 := map[string]interface{}{
		"service": "users",
		"data":    map[string]interface{}{"timestamp": time.Now().Unix()},
	}
	r2, err := sendRequest(req2)
	if err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Println("users reply:", r2)
}
