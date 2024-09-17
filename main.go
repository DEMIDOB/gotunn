package main

import (
	"fmt"
	"flag"
	"log"
	"net"
	"tunnel/client"
	"tunnel/server"
	"tunnel/types"
)

func printCoolAscii() {
	// like in SpringBoot and staff
	// Even though it doesnt look fine now it will drink some coffee and get better at runtime
	// DONT TRY TO FIX WHAT IS ALREADY WORKING
	fmt.Printf("________________________________________________________________________________________________________________\n")
	fmt.Printf("||-------------------------------------------------------------------------------------------------------------|\n")
	fmt.Printf("||             _______        ________     ____________  __         __   __  ______      __  ______           ||\n")
	fmt.Printf("||           //------\\\\      //------\\\\    ------------  ||         ||   || //----\\\\     || //----\\\\          ||\n")
	fmt.Printf("||          //              //        \\\\        ||       ||         ||   ||//      \\\\    ||//      \\\\         ||\n")
	fmt.Printf("||         //              //          \\\\       ||       ||         ||   ||/        \\\\   ||/        \\\\        ||\n")
	fmt.Printf("||         ||  _________   ||          ||       ||       ||         ||   ||         ||   ||         ||        ||\n")
	fmt.Printf("||         \\\\  -------//   \\\\          //       ||       \\\\         //   ||         ||   ||         ||        ||\n")
	fmt.Printf("||          \\\\       //     \\\\        //        ||        \\\\       //    ||         ||   ||         ||        ||\n")
	fmt.Printf("||           \\\\_____//       \\\\______//         ||         \\\\_____//     ||         ||   ||         ||        ||\n")
	fmt.Printf("||            -------         --------          ||           -----       ||         ||   ||         ||        ||\n")
	fmt.Printf("||                                                                                                            ||\n")
	fmt.Printf("|--------------------------------------------------------------------------------------------------------------|\n")
	fmt.Printf("----------------------------------------------------------------------------------------------------------------\n\n\n")
}

func main() {
	isServerPtr := flag.Bool("s", false, "Whether to run in server mode")
	port := flag.Int("p", 0, "Port to accept targets in server mode")
	flag.Parse()

	isServer := *isServerPtr

	if isServer {
		printCoolAscii()
		publicRequestCh := make(chan types.PublicRequest)
		publicResponseCh := make(chan types.PublicResponse)

		go func() {
			err := server.ListenToTargets(*port, publicRequestCh, publicResponseCh)
			if err != nil {
				log.Panic(err)
			}
		}()

		err := server.ListenToPublic(*port+1, publicRequestCh, publicResponseCh)
		if err != nil {
			log.Panic(err)
		}
	} else {
		// temporarily assume there is only one server with only one port to listen
		//ip := net.IPv4(109, 201, 65, 61)
		//ip := net.IPv4(10, 8, 0, 4)
		ip := net.IPv4(127, 0, 0, 1)

		serverPort := 2004
		targetPort := *port

		err := client.ClientPolling(ip, serverPort, targetPort)
		if err != nil {
			log.Panic(err)
		}
	}
}
