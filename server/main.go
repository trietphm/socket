package main

import (
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

	// Accept connection
	_, _, err = syscall.Accept(socketFd)
	if err != nil {
		panic(err)
	}
}
