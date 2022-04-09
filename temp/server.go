package main

import (
	"fmt"
	"net"
)

func main() {
	// create a listener
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("listening failed!", err)
		return
	}
	defer socket.Close()
	for {
		// read data
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("read data failed!", err)
			continue
		}
		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n\n", data[:read])
		// send data
		senddata := []byte("pong!")
		_, err = socket.WriteToUDP(senddata, remoteAddr)
		if err != nil {
			fmt.Println("send data failed!", err)
			return
		}
	}
}
