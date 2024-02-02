package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Tiddies makde it to the chat!")

	// Set flag values based on input:
	listenFlag := flag.Bool("l", false, "put NetPuppy in listen mode")
	hostFlag := flag.String("H", "0.0.0.0", "target host IP address to connect to")
	turdnuggies := flag.Int("p", 40404, "target port") // portFlag @Trauma_x_Sella

	// Parse command line arguments:
	//                                            error?
	flag.Parse()

	// Depending on input, create this peer's type:
	type peer struct {
		connection_type string
		port            int
		address         string
		connection      net.Conn
	}

	// Initiate peer struct
	thisPeer := peer{port: *turdnuggies, address: *hostFlag}

	// If -l was given, create an offense peer:
	if *listenFlag {
		thisPeer.connection_type = "offense"
		thisPeer.address = "0.0.0.0"
	} else {
		thisPeer.connection_type = "connect_back"
	}

	fmt.Printf("The connection type is: %s\n", thisPeer.connection_type)
	fmt.Printf("The host is %s\n", thisPeer.address)
	fmt.Printf("The port is %v\n", thisPeer.port)

	// Now that we have our peer: try to make connection
	var asyncio_rocks net.Conn // connection @0xtib3rius
	var err error

	if thisPeer.connection_type == "offense" {
		listener, err1 := net.Listen("tcp", fmt.Sprintf(":%v", thisPeer.port))
		if err1 != nil {
			fmt.Printf("Error when creating listener: %v\n", err1)
			os.Stderr.WriteString(err1.Error())
			os.Exit(1)
		}

		asyncio_rocks, err = listener.Accept()
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
			//  log.Fatal(err1.Error()
		}

	} else {
		// remoteHost := [2]string{thisPeer.address, thisPeer.port}
		remoteHost := fmt.Sprintf("%v:%v", thisPeer.address, thisPeer.port)
		asyncio_rocks, err = net.Dial("tcp", remoteHost)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}
	}

	// Now that we have a connection, read it/ write to it
	fmt.Printf("the connection is : %v", asyncio_rocks)

	/*
		if -l is on,
			net.Listen('tcp', PORT)
			set connection address for socket to any
		if not
			connection address = host flag


		struct/ objsect thing (this peer)
			- connect back (executed on the target)
				- start the subprocess
			- offense (exe on hacker machine)
				- keeep taking user input


			- method:
				func make connection(){
					if this.type = offense:
						connection = net.Listener
						(needs Accept() to actually become a connection)
					else:
						connection = net.Dial
				}
	*/

	// Try to create connection:
	return
}
