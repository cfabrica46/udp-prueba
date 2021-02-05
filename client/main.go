package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) == 3 {

		serverAddress := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])

		udpAddress, err := net.ResolveUDPAddr("udp", serverAddress)

		if err != nil {
			log.Fatal(err)
		}

		conn, err := net.DialUDP("udp", nil, udpAddress)

		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		for {

			fmt.Print(">")
			r := bufio.NewReader(os.Stdin)

			message, err := r.ReadString('\n')

			if err != nil {
				log.Fatal(err)
			}

			_, err = conn.Write([]byte(message))

			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
