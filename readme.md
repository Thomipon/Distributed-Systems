# Notes for the report

* We use bidirectional streaming. This is not ideal for a production scenario but it is good enough here
* We have a client-server architecture so that we do not have to deal with the pain of peer-to-peer communication
* We only have a single rpc, EnterChat that sets up bidirectional streaming for all subsequent messages. Messages are just a string to represent the text message and an uint64 to represent the Lamport timestamp
* The Lamport timestamp is calculated in the obvious way: Each client and the server keep their own Lamport time stored locally. It is increased every time something is sent and includied in the sent message. When a process receives a message, it updates its own local Lamport time according to the timestamp of the received message.

# Usage

When the server application (ChittyChat_Server.go) is running, start as many clients (ChittyChat_Client.go) as you want. Clients are prompted to enter a (unique) name (if the client requests to join with a name that is already taken by another user, the server refuses the request and the client application will crash). After entering a name, the client joins. Every client can send messages by just typing them in the command line and pressing enter. Linebreaks are not supported. <br>
Messages starting with a '\' are treated as cmmand by the client application. Commands will not be sent to the server and instead be processed locally. Currently the only command is "\exit" which will cause the client to leave the chatroom. At the moment, exiting the chatroom and terminating the program causes an exit code 1 which is suboptimal.