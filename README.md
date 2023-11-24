# Socket Chat in Go

Well, this is an implementation of a web socket in Go.
It has two packages, the client and the server.

Just leave the server running and connect with as many clients as you want.
For now atleast, all the messages only appear on the server and they are not "streamed" to the client

This is something that I look foward to implementing

## Running

```go
# Running the client
cd client/
go run .

# Running the server
cd server/
go run .
```

## So did I learn something?

Yes! I learnt that tcp is the one that is actually responsible for sending the data and a chat app with a dial and listen is quite cool :)
