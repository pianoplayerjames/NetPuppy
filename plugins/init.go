package plugins

import (
    "fmt"
)

type Command interface {
    Execute(comm Communicator, pluginDataChan chan<- string)
    Description() string
}

type Communicator interface {
    Send(message string)
    Receive() (string, error)
    IsServer() bool
}

var Commands = make(map[string]Command)

func Register(name string, cmd Command) {
    if _, exists := Commands[name]; exists {
        fmt.Printf("Warning: Command '%s' is already registered.\n", name)
    }
    Commands[name] = cmd
}
