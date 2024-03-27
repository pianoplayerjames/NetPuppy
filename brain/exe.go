package brain

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "strings"
    "sync"
    "netsquirrel/utils"
    "netsquirrel/plugins"
)

type Client struct {
    conn net.Conn
    name string
}

type ClientCommunicator struct {
    client *Client
    isServer bool
}

var (
    clients sync.Map
)

func (cc *ClientCommunicator) IsServer() bool {
    return cc.isServer
}

func (cc *ClientCommunicator) Send(message string) {
    fmt.Fprintf(cc.client.conn, "%s\n", message)
}

func (cc *ClientCommunicator) Receive() (string, error) {
    reader := bufio.NewReader(cc.client.conn)
    input, err := reader.ReadString('\n')
    return strings.TrimSpace(input), err
}

func broadcast(message string, sender *Client) {
    clients.Range(func(_, value interface{}) bool {
        client := value.(*Client)
        if client != sender {
            _, err := client.conn.Write([]byte(message + "\n"))
            if err != nil {
                log.Printf("Failed to send message to %s: %v", client.name, err)
                client.conn.Close()
                clients.Delete(client)
            }
        }
        return true
    })
}

func handleClient(conn net.Conn) {
    defer conn.Close()

    server := &ClientCommunicator{client: &Client{conn: conn}}
    server.Send("choose a nickname:")

    name, err := server.Receive()
    if err != nil {
        log.Printf("Failed to read nickname: %v", err)
        return
    }

    client := &Client{conn: conn, name: name}
    clients.Store(client, struct{}{})
    log.Printf("%s connected to the server", client.name)
    server.Send(utils.Banner() + "\n\n")
    server.Send("Hey " + client.name + "! Welcome to the server! Type help for commands")

    for {
        command, err := server.Receive()
        if err != nil {
            log.Printf("Error reading from %s: %v", client.name, err)
            break
        }

        log.Printf("[%s] %s", client.name, command)

        if cmd, exists := plugins.Commands[command]; exists {
            pluginDataChan := make(chan string)
            go func() {
                cmd.Execute(server, pluginDataChan)
                close(pluginDataChan)
            }()

            for pluginInput := range pluginDataChan {
                log.Printf("Plugin output: %s", pluginInput)
            }
        } else {
            server.Send(fmt.Sprintf("You said: %s", command))
        }
    }

    clients.Delete(client)
    log.Printf("Closed connection for %s", client.name)
}

func RunServer(port int) {
    server, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
    defer server.Close()
    log.Printf("\n\n" + utils.Banner() + "\n\n")
    log.Printf("Server started on port %d", port)

    for {
        conn, err := server.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }
        go handleClient(conn)
    }
}
