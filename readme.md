# Usage
- First start up the server application `go run ChittyChat_Server.go`
- Once the server is running, start as many clients as you want `go run ChittyChat_Client.go`
- Clients are prompted to enter a (unique) name. After entering a name, the client joins. If the client requests to join with a name that is already taken by another user, the server refuses the request and the client application will crash.
- Every client can send messages by just typing them in the command line and pressing enter. Linebreaks are not supported.
- Messages starting with a `\` are treated as commands by the client application. Commands will not be sent to the server and instead be processed locally. Currently the only command is `\exit` which will cause the client to leave the chatroom. At the moment, exiting the chatroom and terminating the program causes an exit code 1 which is suboptimal.
