# GolangChat

I'm learning golang, so it's just for testing :D

Chat made with Golang using net/rpc module.

## Compilation
```bash 
  //to compile the server
  go build server.go
  
  //to compile the client
  go build client.go
```


## Client
```bash
  ./client --server [direction] --port [port] --name [name]
```
Where `direction` is like 192.168.1.34, `port` like 1234 and `name` like Bob.

## Server
```bash
  ./server --server [direction] --port [port]
```
Where `direction` and `port` are like in the client.

## Basic Client usage
Send a message to other user
```
other_user < This is a message for other_user, hola, hello, salut!
```

You can broadcast a message
```
/broadcast helloooo, this is a message for every user in this server 
```
or 
```
$all < helloooo, this is a message for every user in this server 
```

## All Commands

```
/help       display help
/users      display all connected users
/broadcast  send a message to all the registered users
/whoami     print the user name
```
