package user

import (
	"math/rand"
	"time"
	"chat/message"
)

type User struct{
	Name string
	Hash string
	Messeges chan message.Message 
	LastConnection int64
}
type HashAndUserName struct {
	Hash string
	Name string
}


func CreateUser(name string, users map[string]*User) string{
	hash:=generateHash()
	users[name]=&User{
					Name:name,
					Hash:hash,
					Messeges: make(chan message.Message,10),
					LastConnection: time.Now().Unix(),
				} 
	return hash
}

func IsHashValid(userName string, hash string, users map[string]*User) bool{
	user,ok := users[userName] 
	return ok && user.Hash == hash
}

func generateHash( )string{
	sol:=""
	const abc="abcdefghijklmnñopqrstuvwxyz0123456789"
	const hashLength = 20
	for i := 0 ; i < hashLength ; i++ {
		sol+= string(abc[rand.Intn(len(abc))])
	}  
	return sol 
}


