package chat

import (
	"errors"
	"fmt"
)


var users map[string]User
func init(){
	users=make(map[string]User)
}

type Chat struct {}

func (t *Chat) Send( hashAndMessage HashAndMessage, reply *string ) error{

	hash, message := hashAndMessage.Hash, hashAndMessage.Message	
			
	fmt.Println("User",message.From,"send message to",message.To)
	if !IsHashValid(message.From,hash,users){
		return errors.New("Invalid User")
	}

	if user,ok:=users[message.To];ok{
		user.Messeges <- message
		*reply="Done!"
	}else if message.To == "$all"{
		// send message to all 				
		message.From +=" (broadcast)"
		for user := range users{
			users[user].Messeges <- message
		}

	}else {
		*reply="User not founded!"
		return errors.New("User not found!")
	}
	return nil	
}

func (t *Chat) Recieve( hashAndUser HashAndUserName, reply *[10]Message) error {
	hash, name := hashAndUser.Hash, hashAndUser.Name
	if !IsHashValid(name,hash,users){
		return errors.New("Invalid User")
	}	
	
	user:=users[name]		
	i := 0
	for ;len(user.Messeges) != 0 && i < len(reply); {
		reply[i] = <- user.Messeges
		i++
	}
	return nil
}

func (t *Chat) Register( name string, reply *string ) error{
	if _,ok := users[name] ; ok {
		*reply = "User exists!"
		return  errors.New("UserExist!")
	}
	hash:=CreateUser(name,users)
	fmt.Println("User",name,"registered!" )
	*reply=hash
	return nil
}

func (t *Chat) GetAllUsers(hashAndUser HashAndUserName, reply* string) error{
	hash, name := hashAndUser.Hash, hashAndUser.Name
	if !IsHashValid(name,hash,users){
		return errors.New("Invalid User")
	}				

	sol := ""
	for s := range users{
		if s == name {
			sol+= "*"+s+"* "
		} else {
			sol+=s+" "
		}
		
	}
	*reply = sol	
	return nil
}

