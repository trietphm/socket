package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Open socket
	socketFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	// Bind socket to an address
	sockAddress := syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1},
	}
	err = syscall.Bind(socketFd, &sockAddress)

	if err != nil {
		panic(err)
	}

	// Listen
	err = syscall.Listen(socketFd, syscall.SOMAXCONN)
	if err != nil {
		panic(err)
	}

	fmt.Println("accept")
	// Accept connection
	nfd, _, err := syscall.Accept(socketFd)
	if err != nil {
		panic(err)
	}

	// read data from client
	p := make([]byte, 100)
	_, err = syscall.Read(nfd, p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(p))

	// Close socket
	if err := syscall.Close(socketFd); err != nil {
		panic(err)
	}
}
