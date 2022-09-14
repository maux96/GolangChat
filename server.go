package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/http"
	"time"
	"chat/chat"
)


func main(){
	err:=rpc.Register(new(chat.Chat))
	if err != nil {
		fmt.Println("Error in Chat register:",err.Error())
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error in listening:",err.Error())
	}
	go http.Serve(listener,nil)	

	fmt.Println("Server started!")		
	for {
		time.Sleep(20e9)
		// do somthing like kill no connected usernames :D
	}
}
