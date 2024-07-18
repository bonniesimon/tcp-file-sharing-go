### Resources
- My main inspiration to start this project  https://github.com/schollz/croc/blob/a82fe6f074de5f42ad57fdd1b97f7284d7af901b/client.go (early commit)
- Distributed file sharing system in go https://www.machinet.net/tutorial-eng/create-distributed-file-sharing-system-go
- wShare https://github.com/thewhitetulip/wshare/blob/master/main.go
- How To Stream Large Files Over TCP In Golang  https://www.youtube.com/watch?v=82oFmY-Qeok
- Simplest TCP server for file transfer I could find https://github.com/pplam/tcp-file-transfer/blob/master/client/main.go
- Other file sharing golang projects from github https://github.com/topics/file-sharing?l=go&o=desc&s=updated
- Another implementation with selecting devices in local network https://github.com/varuuntiwari/share/blob/main/send/send.go
  - Maybe I can also implement this feature of selecting which device to send to.
- Theory behind TCP server that clears a lot of my doubts https://youtu.be/f9gUFy-9uCM

### Logs

#### 17 July 2024
- In case of a web server, the server is the application we write and the client is the browser (or curl). Thus we need two programs. That is why we see client and server in the related go projects that we look at.
- Connect to TCP server using curl
  ```
  ❯ curl -v telnet://127.0.0.1:3005
  ```
- TCP socket primitives:
  SOCKET:    Create a new communication endpoint
  BIND:      Associate a local address with a socket
  LISTEN:    Announce willingness to accept connections; give queue size
  ACCEPT:    Passively establish an incoming connection
  CONNECT:   Actively attempt to establish a connection
  SEND:      Send some data over the connection
  RECEIVE:   Receive some data from the connection
  CLOSE:     Release the connection
- Next step is implementing the case where we can keep sending stuff to the server without the server closing. I should checkout the initial commit of croc to see how they do this