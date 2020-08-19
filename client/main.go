package main

import (
	"syscall"
)

func main() {
	// Open socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	sockAddress := syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1},
	}
	if err := syscall.Connect(fd, &sockAddress); err != nil {
		panic(err)
	}
}
