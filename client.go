package main

import (
	"fmt"
	"net/rpc"
	"os"
	"math/rand"
	"time"
	"chat/chat"
	"bufio"
	"strings"
	"strconv"
)

const SERVER = "localhost"
const refreshTime = 5

var session_hash string
var client *rpc.Client
var userName string

func getMessages(){

	hashAndUser := chat.HashAndUserName{Hash:session_hash, Name:userName} 
	for {	
		var reply [10]chat.Message
		err := client.Call("Chat.Recieve",hashAndUser, &reply)
		if err != nil {
			fmt.Println("Error in get messages",err.Error())
		}

		for _,message := range reply {
			if message.From != ""{
				fmt.Println("\r"+message.From,":",message.Body)
				fmt.Print(">")
			}
		}
		time.Sleep(refreshTime * 1e9)
	}
}

func resolveCommand(toSend string) bool{
	hashAndUser:= chat.HashAndUserName{Name:userName,Hash:session_hash}

	vec := strings.Split(toSend," ")
	switch vec[0] {
		case "/users":
			// get users
			fmt.Println("\rUsers:")
			var reply string 
			err :=client.Call("Chat.GetAllUsers",hashAndUser,&reply)
			if err != nil {
				fmt.Println("Error in /users:",err.Error())
				return true		
			}
			fmt.Println(reply)
		case "/broadcast":
			fmt.Println("Sending to all users...")	
			var reply string 
			err :=client.Call("Chat.Send",chat.HashAndMessage{Hash:session_hash,Message: chat.Message{From:userName,To:"$all",Body:strings.Join(vec[1:]," ")}},&reply)
			if err != nil {
				fmt.Println("Error in /broadcast:",err.Error())
				return true		
			}	
			
		case "/help":
			fmt.Println("Basic help:")		
			fmt.Println("To send a message to a registered user use:\n\t paco < this is my message to user paco\n")

			fmt.Println(" /help       display this help")
			fmt.Println(" /users      display all registered users")
			fmt.Println(" /broadcast  send a message to all the registered users")

			fmt.Println("\nhttps://github.com/maux96")

		default:
			return false
	}
	return true
}

func startSendingMessages(){
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("For help write /help.")		

	for {
		fmt.Print("> ")
		toSend,_ := inputReader.ReadString('\n')
		toSend=strings.Trim(toSend, " \n\r")
		
		if resolveCommand(toSend) {
			continue	
		}

		nameAndMessage := strings.Split(toSend,"<")
		if len(nameAndMessage) != 2{
			fmt.Println("Error: InvalidMessage, char '<' must be writed once.")
			continue
		}
		to := strings.Trim(nameAndMessage[0]," \r\n")
		body := strings.Trim(nameAndMessage[1]," \r\n")
		message := chat.Message{From: userName, To:to, Body: body}
		var reply string	
		err :=client.Call("Chat.Send",chat.HashAndMessage{Hash:session_hash,Message:message}, &reply)
		if err != nil {
			fmt.Println("Error sending the message",err.Error())
			continue
		}
		//fmt.Println("Message response:",reply)
	}

}

func register() error{

	var reply string
	err := client.Call("Chat.Register",userName, &reply)
	if err != nil {
		return err 
	}
	session_hash = reply		
	return nil	
}

func main(){

	rand.Seed(time.Now().Local().Unix())
	if len(os.Args)>1{
		userName=os.Args[1]
		userName=strings.Trim(userName," \r\n") 
		if userName == "" || userName[0] == '$' || userName[0] =='/' {
			panic("No permited user name!")
		}	
	} else {
		rnum := strconv.FormatInt(int64(rand.Intn(50000)),10)
		userName = "unknowUser_" + rnum
	}


	fmt.Println("Initializing client...")		

	var err error
	client, err = rpc.DialHTTP("tcp",SERVER+":1234")
	if err != nil {
		fmt.Println("Error Dialing: ", err.Error())
		return 
	}

	err = register()
	if err != nil {
		fmt.Println("Error in register:",err.Error())
		return
	}	
		
	go getMessages()
	startSendingMessages()
}
