package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/http"
	"time"
	"chat/remote_chat"
	"os"
	"strconv"
)

var (
	server="localhost"
	port="1234"
)

func processServerArgs(){
	for i := range os.Args{
		switch os.Args[i]{
			case "--server":
				server=os.Args[i+1]	
			case "--port":
				port=os.Args[i+1]	
			case "--max-dead-time":
				deadTime,err:= strconv.Atoi(os.Args[i+1])
				if err!=nil{
					fmt.Println("Wrong value in --max-dead-time :",err.Error())
				}	
				remote_chat.MAX_DEAD_TIME = int64(deadTime)
			case "--help":
				fmt.Println("Help:")
				fmt.Println("--server    Set server address. (default: localhost)")
				fmt.Println("--port      Set port. (default: 1234)\n")

				fmt.Println("\nhttps://github.com/maux96\n")
				os.Exit(1)
				
		}
	}
}


func main(){
	processServerArgs()

	err:=rpc.Register(new(remote_chat.Chat))
	if err != nil {
		fmt.Println("Error in Chat register:",err.Error())
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", server+":"+port)
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
