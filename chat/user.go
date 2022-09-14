package chat

import (
	"math/rand"
)

type User struct{
	Name string
	Hash string
	Messeges chan Message 
}
type HashAndUserName struct {
	Hash string
	Name string
}


func CreateUser(name string, users map[string]User) string{
	hash:=generateHash()
	users[name]=User{
					Name:name,
					Hash:hash,
					Messeges: make(chan Message,10),
				} 
	return hash
}

func IsHashValid(userName string, hash string, users map[string]User) bool{
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

