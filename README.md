## Bidirectional streaming with GRPC

The main advantage of GRPC over traditional HTTP/1.1 is that it uses a single TCP
connection for sending and receiving multiple messages between the server and the client.

Another such use case is when a server needs to notify the client whenever some processing
is performed. This is called a server push model. The server can send a stream of results
back when a client asked for them only once. This is different to polling, where the client
requests something each and every time. This can be useful when there are a series of time-
taking steps that need to be done. The GRPC client can escalate that job to the GRPC server.
Then, the server takes its time and relays the message back to the client, which reads them
and does something useful. Let us implement this.

PS: Knowledge of Protobuff is required to understand and work with the repository.
