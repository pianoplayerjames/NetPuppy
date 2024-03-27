package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "netsquirrel/brain"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter port to start the server:")
    portInput, _ := reader.ReadString('\n')
    port, err := strconv.Atoi(strings.TrimSpace(portInput))
    if err != nil {
        fmt.Println("Invalid port number. Exiting.")
        return
    }
    brain.RunServer(port)
}
