# Notes for the report

* We use bidirectional streaming. This is not ideal for a production scenario but it is good enough here
* We have a client-server architecture so that we do not have to deal with the pain of peer-to-peer communication
* We only have a single rpc, EnterChat that sets up bidirectional streaming for all subsequent messages. Messages are just a string to represent the text message and an uint64 to represent the Lamport timestamp
* The Lamport timestamp is calculated in the obvious way: Each client and the server keep their own Lamport time stored locally. It is increased every time something is sent and includied in the sent message. When a process receives a message, it updates its own local Lamport time according to the timestamp of the received message.