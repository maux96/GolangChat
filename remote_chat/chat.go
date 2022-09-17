package remote_chat 

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"chat/user"
	"chat/message"
)


var MAX_DEAD_TIME int64=20

var users map[string]*user.User
func init(){
	users=make(map[string]*user.User)
}

type Chat struct {}

func (t *Chat) Send( hashAndMessage message.HashAndMessage, reply *string ) error{

	hash, message := hashAndMessage.Hash, hashAndMessage.Message	
			
	fmt.Println("User",message.From,"send message to",message.To)
	if !user.IsHashValid(message.From,hash,users){
		return errors.New("Invalid User")
	}
	
	if strings.Trim(message.Body," \n\r") == ""{
		return errors.New("The message can't be empty!")
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

func (t *Chat) Recieve( hashAndUser user.HashAndUserName, reply *[10]message.Message) error {
	hash, name := hashAndUser.Hash, hashAndUser.Name
	if !user.IsHashValid(name,hash,users){
		return errors.New("Invalid User")
	}	
	
	user:=users[name]		
	user.LastConnection=time.Now().Unix()
	i := 0
	for ;len(user.Messeges) != 0 && i < len(reply); {
		reply[i] = <- user.Messeges
		i++
	}
	return nil
}

func (t *Chat) Register( name string, reply *string ) error{
	if user,ok := users[name] ; ok && time.Now().Unix()-(*user).LastConnection < MAX_DEAD_TIME {
		fmt.Println("User Creation for '",name,"' denied, user exist!")
		*reply = "User exists!"
		return  errors.New("UserExist!")
	}
	hash:=user.CreateUser(name,users)
	fmt.Println("User",name,"registered!" )
	*reply=hash
	return nil
}

func (t *Chat) GetAllUsers(hashAndUser user.HashAndUserName, reply* string) error{
	hash, name := hashAndUser.Hash, hashAndUser.Name
	if !user.IsHashValid(name,hash,users){
		return errors.New("Invalid User")
	}				

	sol := ""
	for s,user := range users{
		if time.Now().Unix()-(*user).LastConnection > MAX_DEAD_TIME {
			// the user is disconected 
			delete(users,s)
			fmt.Println("User",s,"removed. (max dead time)")
			continue
		}
		if s == name {
			sol+= "*"+s+"* "
		} else {
			sol+=s+" "
		}
		
	}
	*reply = sol	
	return nil
}

