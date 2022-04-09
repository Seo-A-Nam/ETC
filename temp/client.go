package main

import (
	"fmt"
	"net"
)

func main() {
	// create a connection
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(141, 223, 65, 111),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("connection failed!", err)
		return
	}
	defer socket.Close()
	// send data
	senddata := []byte("ping!")
	_, err = socket.Write(senddata)
	if err != nil {
		fmt.Println("send data failed!", err)
		return
	}
	// receive data
	data := make([]byte, 4096)
	read, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read data failed!", err)
		return
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", data[:read])
}
