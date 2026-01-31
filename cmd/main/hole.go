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
					time.Sleep(100*time.Millisecond)
			}
		}
	}(ch)
	con.SetReadDeadline(time.Now().Add(time.Second*10))
	_, _, err := con.ReadFromUDP(buf)
	ch<-true
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		return false
	}
	return true
}
