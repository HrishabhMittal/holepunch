package main

import (
	"fmt"
	"net"
	"time"
)
func punchHole(con *net.UDPConn,addr *net.UDPAddr) (bool) {
	buf := make([]byte, 128)
	ch := make(chan bool,1)
	go func(chan bool) {
		for {
			select {
				case <-ch:
					return
				default:
					con.WriteToUDP([]byte("hello"),addr)
			}
		}
	}(ch)
	fmt.Println("send hello to addr")
	con.SetReadDeadline(time.Now().Add(time.Second*10))
	n, _, err := con.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		return false
	}
	fmt.Printf("read %d bytes\n",n)
	return true
}

